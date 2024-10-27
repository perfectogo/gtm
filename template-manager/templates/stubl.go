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

//go:embed handler.go.templ
var handlerStub string

func GetHandlerTemplate() string {
	if handlerStub == "" {
		log.Fatalln("handler stub is empty")
	}
	return handlerStub
}

//go:embed middleware.go.templ
var middlewareStub string

func GetMiddlewareTemplate() string {
	if middlewareStub == "" {
		log.Fatalln("middleware stub is empty")
	}
	return middlewareStub
}

//go:embed .gitignore.go.templ
var gitignoreStub string

func GetGitignoreTemplate() string {
	if gitignoreStub == "" {
		log.Fatalln("gitignore stub is empty")
	}
	return gitignoreStub
}

//go:embed .env.go.templ
var envStub string

func GetEnvTemplate() string {
	if envStub == "" {
		log.Fatalln("env stub is empty")
	}
	return envStub
}

//go:embed README.go.templ
var readmeStub string

func GetReadmeTemplate() string {
	if readmeStub == "" {
		log.Fatalln("readme stub is empty")
	}
	return readmeStub
}
