package filter

import (
	"database/sql"
	"net/http"
	"net/url"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id       int64
	Username string `filter:"param:login;searchable;filterable"`
	FullName string `filter:"param:name;searchable"`
	Email    string `filter:"filterable"`
	// This field is not filtered.
	Password string
}

type TestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func (t *TestSuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)

	db, t.mock, err = sqlmock.New()
	t.NoError(err)
	t.NotNil(db)
	t.NotNil(t.mock)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	t.db, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t.T(), err)
	assert.NotNil(t.T(), t.db)
}

func (t *TestSuite) TearDownTest() {
	db, err := t.db.DB()
	assert.NoError(t.T(), err)
	db.Close()
}

// TestFilterBasic is a test suite for basic filters functionality.
func (t *TestSuite) TestFilterBasic() {
	var users []User
	ctx := gin.Context{}
	ctx.Request = &http.Request{
		URL: &url.URL{
			RawQuery: "filter=login:sampleUser",
		},
	}

	t.mock.ExpectQuery(`^SELECT \* FROM "users" WHERE "Username" = \$1`).
		WithArgs("sampleUser").
		WillReturnRows(sqlmock.NewRows([]string{"id", "Username", "FullName", "Email", "Password"}))

	err := t.db.Model(&User{}).Scopes(FilterByQuery(&ctx, ALL)).Find(&users).Error
	t.NoError(err)
}

// Filtering for a field that is not filterable should not be performed
func (t *TestSuite) TestFilterNotFilterable() {
	var users []User
	ctx := gin.Context{}
	ctx.Request = &http.Request{
		URL: &url.URL{
			RawQuery: "filter=password:samplePassword",
		},
	}
	t.mock.ExpectQuery(`^SELECT \* FROM "users" ORDER`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "Username", "FullName", "Email", "Password"}))
	err := t.db.Model(&User{}).Scopes(FilterByQuery(&ctx, ALL)).Find(&users).Error
	t.NoError(err)
}

// Filtering would not be applied if no config is provided.
func (s *TestSuite) TestFiltersNoFilterConfig() {
	var users []User
	ctx := gin.Context{}
	ctx.Request = &http.Request{
		URL: &url.URL{
			RawQuery: "filter=login:sampleUser",
		},
	}

	s.mock.ExpectQuery(`^SELECT \* FROM "users"$`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "Username", "FullName", "Email", "Password"}))
	err := s.db.Model(&User{}).Scopes(FilterByQuery(&ctx, SEARCH)).Find(&users).Error
	s.NoError(err)
}

// TestFiltersSearchable is a test suite for searchable filters functionality.
func (s *TestSuite) TestFiltersSearchable() {
	var users []User
	ctx := gin.Context{}
	ctx.Request = &http.Request{
		URL: &url.URL{
			RawQuery: "search=John",
		},
	}

	s.mock.ExpectQuery(`^SELECT \* FROM "users" WHERE \("Username" LIKE \$1 OR "FullName" LIKE \$2\)`).
		WithArgs("%John%", "%John%").
		WillReturnRows(sqlmock.NewRows([]string{"id", "Username", "FullName", "Email", "Password"}))
	err := s.db.Model(&User{}).Scopes(FilterByQuery(&ctx, ALL)).Find(&users).Error
	s.NoError(err)
}

// TestFiltersPaginateOnly is a test suite for pagination functionality.
func (s *TestSuite) TestFiltersPaginateOnly() {
	var users []User
	ctx := gin.Context{}
	ctx.Request = &http.Request{
		URL: &url.URL{
			RawQuery: "page=2&per_page=10",
		},
	}

	s.mock.ExpectQuery(`^SELECT \* FROM "users" ORDER BY "id" DESC LIMIT 10 OFFSET 10$`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "Username", "FullName", "Email", "Password"}))
	err := s.db.Model(&User{}).Scopes(FilterByQuery(&ctx, ALL)).Find(&users).Error
	s.NoError(err)
}

// TestFiltersOrderBy is a test suite for order by functionality.
func (t *TestSuite) TestFiltersOrderBy() {
	var users []User
	ctx := gin.Context{}
	ctx.Request = &http.Request{
		URL: &url.URL{
			RawQuery: "order_by=Email&order_direction=asc",
		},
	}

	t.mock.ExpectQuery(`^SELECT \* FROM "users" ORDER BY "Email"$`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "Username", "FullName", "Email", "Password"}))
	err := t.db.Model(&User{}).Scopes(FilterByQuery(&ctx, ORDER_BY)).Find(&users).Error
	t.NoError(err)
}

func TestRunSu(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
