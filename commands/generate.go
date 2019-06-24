package commands

import (
	"database/sql"
	"errors"
	"fmt"
	"go/build"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/afanxia/gin-scaffold/symbol"
	"github.com/codegangsta/cli"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func srcTemplateDir(ctx *cli.Context) (string, error) {
	templateFolderDir := ctx.GlobalString("template-folder")
	if templateFolderDir != "" {
		return templateFolderDir, nil
	}

	template := ctx.GlobalString("template")
	gopath := build.Default.GOPATH
	if gopath == "" {
		return "", errors.New("Abort: GOPATH environment variable is not set. ")
	}
	return path.Join(filepath.SplitList(gopath)[0], "src", "github.com/afanxia/gin-scaffold", "templates", template), nil
}

func destProjectDir(project string) (string, error) {
	gopath := build.Default.GOPATH
	if gopath == "" {
		return "", errors.New("Abort: GOPATH environment variable is not set. ")
	}

	return path.Join(filepath.SplitList(gopath)[0], "src", project), nil
}

func projectName(project string) string {
	idx := strings.Split(project, "/")
	return idx[len(idx)-1]
}

func hasSuffix(fpath string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(fpath, suffix) {
			return true
		}
	}
	return false
}

// Generate is the main entry
func Generate(ctx *cli.Context) {
	templateDir, err := srcTemplateDir(ctx)
	if err != nil {
		fmt.Println("scaffold generate failed:", err)
		return
	}
	if templateDir == "" {
		fmt.Println("scaffold generate failed: none templates")
		return
	}

	args := ctx.Args()
	if len(args) != 1 {
		fmt.Println("Usage: scaffold generate <project/path>")
		return
	}

	project := args[0]
	projectDir, err := destProjectDir(project)
	projectName := projectName(project)
	if err != nil {
		fmt.Println("scaffold generate failed:", err)
		return
	}

	suffixes := ctx.GlobalStringSlice("include-template-suffix")
	ignores := ctx.GlobalStringSlice("exclude-template-suffix")
	force := ctx.GlobalBool("force")
	driver := ctx.String("driver")
	database := ctx.String("database")
	host := ctx.String("host")
	port := ctx.Int("port")
	user := ctx.String("username")
	pass := ctx.String("password")

	dsn := symbol.DSNFormat(host, port, user, pass, database)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		fmt.Println("scaffold generate failed:", err)
		return
	}
	defer db.Close()

	tables, err := symbol.GetAllTables(db)
	if err != nil {
		fmt.Println("scaffold generate failed:", err)
		return
	}

	if err := filepath.Walk(templateDir, func(srcPath string, info os.FileInfo, err error) error {
		data := map[string]interface{}{
			"project": project,
			"tables":  tables,
			"projectName": projectName,
		}
		for i, table := range tables {
			data["table"] = table
			data["index"] = i

			pathName, err := symbol.RenderString(srcPath, data)
			if err != nil {
				return err
			}

			destPath := path.Join(projectDir, strings.TrimPrefix(pathName, templateDir))

			if info.IsDir() {
				if info.Name() == "vendor" { // gonna skip "vendor" directory
					if err := symbol.CopyDir(srcPath, destPath, force); err != nil {
						return err
					}
				} else {
					if err := os.MkdirAll(destPath, info.Mode()); err != nil {
						return err
					}
				}
			} else {
				if hasSuffix(pathName, ignores) {
					break
				} else if hasSuffix(pathName, suffixes) {
					if err := symbol.RenderTemplate(srcPath, destPath, data, force); err != nil {
						return err
					}
				} else {
					if err := symbol.CopyFile(srcPath, destPath, force); err != nil {
						return err
					}
				}
			}

			if srcPath == pathName {
				break
			}
		}
		return nil
	}); err != nil {
		fmt.Println("scaffold generate failed:", err)
		return
	}
	if len(suffixes) == 0 {
		fmt.Println("WARNING: scaffold template suffix is unset, just copy files.")
	}
}
