package templates

import (
	_ "embed"
	"log"
)

//go:embed cmd.go.templ
var cmdStub string

func GetCmdTemplate() string {
	if cmdStub == "" {
		log.Fatalln("cmd stub is empty")
	}
	return cmdStub
}

//go:embed api.go.templ
var apiStub string

func GetApiTemplate() string {
	if apiStub == "" {
		log.Fatalln("api stub is empty")
	}
	return apiStub
}

//go:embed middleware.go.templ
var middlewareStub string

func GetMiddlewareTemplate() string {
	if middlewareStub == "" {
		log.Fatalln("middleware stub is empty")
	}
	return middlewareStub
}
