package faker

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewUserFaker),
	fx.Provide(NewFaker),
)

type Faker interface {
	Setup()
}

type Fakers []Faker

func (f Fakers) Setup() {
	for _, faker := range f {
		faker.Setup()
	}
}

func NewFaker(userFaker UserFaker) Fakers {
	return Fakers{
		userFaker,
	}
}
