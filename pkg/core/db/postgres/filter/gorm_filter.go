package filter

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/fazanurfaizi/go-rest-template/pkg/converter"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type queryParams struct {
	Search         string   `form:"search"`
	Filter         string   `form:"filter"`
	Page           int      `form:"page,default=1"`
	PageSize       int      `form:"page_size,default=10"`
	All            bool     `form:"all,default=false"`
	OrderBy        string   `form:"order_by,default=created_at"`
	OrderDirection string   `form:"order_dir,default=desc,oneof=desc asc"`
	Includes       []string `form:"includes"`
}

const (
	SEARCH   = 1 // Search response with LIKE query "search={search_phrase}"
	FILTER   = 2 // Filter response by column name values "filter={column_name}:{value}"
	PAGINATE = 4 // Paginate response with page and page_size
	ORDER_BY = 8 // Order response by column_name
	INCLUDES = 10
	ALL      = 15 // Equivalent to SEARCH|FILTER|PAGINATE|ORDER_BY
	tagKey   = "filter"
)

func orderBy(db *gorm.DB, params queryParams) *gorm.DB {
	return db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: params.OrderBy},
		Desc:   params.OrderDirection == "desc",
	})
}

func paginate(db *gorm.DB, params queryParams) *gorm.DB {
	if params.All {
		return db
	}

	if params.Page == 0 {
		params.Page = 1
	}

	switch {
	case params.PageSize > 100:
		params.PageSize = 100
	case params.PageSize <= 0:
		params.PageSize = 10
	}

	offset := (params.Page - 1) * params.PageSize
	return db.Offset(offset).Limit(params.PageSize)
}

func includes(db *gorm.DB, params string) *gorm.DB {
	return db.Preload(params)
}

func searchField(field reflect.StructField, phrase string) clause.Expression {
	filterTag := field.Tag.Get(tagKey)
	columnName := converter.GetColumnNameForField(field)
	if strings.Contains(filterTag, "searchable") {
		return clause.Like{Column: columnName, Value: "%" + phrase + "%"}
	}
	return nil
}

func filterField(field reflect.StructField, phrase string) clause.Expression {
	var paramName string
	if !strings.Contains(field.Tag.Get(tagKey), "filterable") {
		return nil
	}

	columnName := converter.GetColumnNameForField(field)
	paramMatch := converter.ParamNameRegexp.FindStringSubmatch(field.Tag.Get(tagKey))
	if len(paramMatch) == 2 {
		paramName = paramMatch[1]
	} else {
		paramName = columnName
	}

	re, err := regexp.Compile(fmt.Sprintf(`(?m)%v:(\w{1,}).*`, paramName))
	if err != nil {
		return nil
	}
	filterSubPhraseMatch := re.FindStringSubmatch(phrase)
	if len(filterSubPhraseMatch) == 2 {
		return clause.Eq{Column: columnName, Value: filterSubPhraseMatch[1]}
	}
	return nil
}

func expressionByField(
	db *gorm.DB,
	phrase string,
	modelType reflect.Type,
	operator OperatorFunc,
	predicate PredicateFunc,
) *gorm.DB {
	numFields := modelType.NumField()
	expressions := make([]clause.Expression, 0, numFields)
	for i := 0; i < numFields; i++ {
		field := modelType.Field(i)
		expression := operator(field, phrase)
		if expression != nil {
			expressions = append(expressions, expression)
		}
	}
	if len(expressions) == 1 {
		db = db.Where(expressions[0])
	} else if len(expressions) > 1 {
		db = db.Where(predicate(expressions...))
	}

	return db
}

func FilterByQuery(c *gin.Context, config int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var params queryParams
		err := c.BindQuery(&params)
		if err != nil {
			return db
		}

		model := db.Statement.Model
		modelType := reflect.TypeOf(model)
		if model != nil && modelType.Kind() == reflect.Ptr && modelType.Elem().Kind() == reflect.Struct {
			if config&SEARCH > 0 && params.Search != "" {
				db = expressionByField(db, params.Search, modelType.Elem(), searchField, clause.Or)
			}
			if config&FILTER > 0 && params.Filter != "" {
				db = expressionByField(db, params.Filter, modelType.Elem(), filterField, clause.And)
			}
		}

		if config&INCLUDES > 0 && len(params.Includes) > 0 {
			for _, v := range params.Includes {
				db = includes(db, v)
			}
		}

		if config&ORDER_BY > 0 {
			db = orderBy(db, params)
		}

		if config&PAGINATE > 0 {
			db = paginate(db, params)
		}

		return db
	}
}

type OperatorFunc func(reflect.StructField, string) clause.Expression
type PredicateFunc func(...clause.Expression) clause.Expression
