package initialize

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"os/exec"
)

func UpdateImports(file *ast.File, newServiceName, modPath string) error {

	var importDecl *ast.GenDecl
	for _, decl := range file.Decls {
		if decl, ok := decl.(*ast.GenDecl); ok && decl.Tok == token.IMPORT {
			importDecl = decl
			break
		}
	}

	if importDecl != nil {
		import1 := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("%sHandler \"%s/%s/handler\"", ToLowerCamel(newServiceName), modPath, ToSnake(newServiceName)),
			},
		}
		import2 := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("%sService \"%s/%s/service\"", ToLowerCamel(newServiceName), modPath, ToSnake(newServiceName)),
			},
		}

		importDecl.Specs = append(importDecl.Specs, import1, import2)
	} else {
		return errors.New("import block not found in the file")
	}

	return nil
}

func GoModInit(modPath, dir string) error {
	fmt.Println(modPath, dir)
	if _, err := exec.LookPath("go"); err != nil {
		return errors.New("go path not found, please install go")
	} else {
		cmd := exec.Command("go", "mod", "init", modPath)
		cmd.Dir = dir
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to execute go mod init: %v", err)
		}
	}

	return nil
}

func InitProject(modPath, projectName string) error {

	if err := GoModInit(modPath+"/"+ToSnake(projectName), "./"+ToSnake(projectName)); err != nil {
		return err
	}

	// if err := GoModTidy("./" + projectName); err != nil {
	// 	return err
	// }

	// if err := RunGoImports("golang.org/x/tools/cmd/goimports@latest", "./"+projectName); err != nil {
	// 	return err
	// }

	return nil
}

func GoModTidy(dir string) error {
	if _, err := exec.LookPath("go"); err != nil {
		return errors.New("go path not found, please install go")
	} else {
		cmd := exec.Command("go", "mod", "tidy")
		cmd.Dir = dir
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to execute go mod tidy: %v", err)
		}
	}

	return nil
}

func RunGoImports(packagePath, dir string) error {
	if _, err := exec.LookPath("go"); err != nil {
		return errors.New("go path not found, please install go")
	} else {
		cmd := exec.Command("go", "install", packagePath)
		cmd.Dir = dir
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install goimports %v", err)
		}
	}

	if _, err := exec.LookPath("go"); err != nil {
		return errors.New("go path not found, please install go")
	} else {
		cmd := exec.Command("goimports", "-w", ".")
		cmd.Dir = dir
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to execute goimports: %v", err)
		}
	}

	return nil
}
