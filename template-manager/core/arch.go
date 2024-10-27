package core

import (
	"fmt"
	"log"

	"github.com/perfectogo/gtm/template-manager/initialize"
	"github.com/perfectogo/gtm/template-manager/templates"
)

func CreateApp(ProjectName, modPath string, crudName *string) {

	// create folders
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

	if err := initialize.CreateFolder(folders); err != nil {
		log.Fatalln(err)
	}

	// create files
	var files map[string]string = map[string]string{

		// cmd
		fmt.Sprintf("%s/cmd/main.go", ProjectName): templates.GetCmdTemplate(),

		// api
		fmt.Sprintf("%s/api/api.go", ProjectName):                   templates.GetApiTemplate(),
		fmt.Sprintf("%s/api/handlers/handlers.go", ProjectName):     templates.GetHandlerTemplate(),
		fmt.Sprintf("%s/api/middleware/middleware.go", ProjectName): templates.GetMiddlewareTemplate(),

		// internal

		// other files
		fmt.Sprintf("%s/.gitignore", ProjectName): templates.GetGitignoreTemplate(),
		fmt.Sprintf("%s/.env", ProjectName):       templates.GetEnvTemplate(),
		fmt.Sprintf("%s/README.md", ProjectName):  templates.GetReadmeTemplate(),
	}

	if err := initialize.CreateFile(ProjectName, files); err != nil {
		log.Fatalln(err)
	}

	// init project
	if err := initialize.InitProject(modPath, ProjectName); err != nil {
		log.Fatalln(err)
	}
}
