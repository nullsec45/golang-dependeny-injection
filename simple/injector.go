//go:build wireinject
// +build wireinject
package simple

import (
	"github.com/google/wire"
)		

func InitalizedService(isError bool) (*SimpleService, error){
	wire.Build(
		NewSimpleRepository, 
		NewSimpleService,
	)
	return nil, nil
}

func InitalizedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)

	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}