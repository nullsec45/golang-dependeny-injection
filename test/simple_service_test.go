package test

import (
	"testing"
	// "fmt"
	"github.com/stretchr/testify/assert"
	"nullsec45/golang-dependency-injection/simple"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitalizedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)

	// fmt.Println(err)

	// if error false
	// fmt.Println(simpleService.SimpleRepository)

	// fmt.Println(simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitalizedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}