package v1

import "github.com/miaogu-go/Gof/api/v1/test"

type ApiV1Group struct {
	TestGroup test.Test
}

var ApiV1App = new(ApiV1Group)
