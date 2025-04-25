package test

import (
	"fmt"
	"testing"

	"github.com/fahrilhadi/golang-restful-api/simple"
)

func TestSimpleService(t *testing.T)  {
	simpleService := simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}