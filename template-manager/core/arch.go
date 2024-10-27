package core

import (
	"fmt"
	"log"

	"github.com/perfectogo/gtm/template-manager/initialize"
	"github.com/perfectogo/gtm/template-manager/templates"
)

func CreateApp(ProjectName, modPath string, crudName *string) {

	var folders []string = []string{
		ProjectName,
		fmt.Sprintf("%s/cmd", ProjectName),

		// api
		fmt.Sprintf("%s/api", ProjectName),
		fmt.Sprintf("%s/api/handlers", ProjectName),
		fmt.Sprintf("%s/api/middleware", ProjectName),

		// internal
		fmt.Sprintf("%s/internal", ProjectName),
		fmt.Sprintf("%s/internal/config", ProjectName),
		fmt.Sprintf("%s/internal/database", ProjectName),
		fmt.Sprintf("%s/internal/models", ProjectName),
		fmt.Sprintf("%s/internal/repository", ProjectName),
		fmt.Sprintf("%s/internal/services", ProjectName),

		// pkg
		fmt.Sprintf("%s/pkg", ProjectName),
	}

	var files map[string]string = map[string]string{

		// cmd
		fmt.Sprintf("%s/cmd/main.go", ProjectName): templates.GetCmdTemplate(),

		// api
		fmt.Sprintf("%s/api/api.go", ProjectName): templates.GetApiTemplate(),
		fmt.Sprintf("%s/api/middleware/middleware.go", ProjectName): templates.GetMiddlewareTemplate(),

		// internal
	}

	if err := initialize.CreateFolder(folders); err != nil {
		log.Fatalln(err)
	}

	if err := initialize.CreateFile(ProjectName, files); err != nil {
		log.Fatalln(err)
	}

}
