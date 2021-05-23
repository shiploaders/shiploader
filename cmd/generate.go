package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/urfave/cli"
)

type Params struct {
	AppName string
	Image   string
	Port    string
	Replica string
}

// Will return the command line ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Shiploader"
	app.Usage = "Let's you query Image, Replica, Port!"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "app",
		},
		cli.StringFlag{
			Name:  "replica",
			Value: "1",
		},
		cli.StringFlag{
			Name:  "image",
			Value: "gcr.io/webera/base",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "8080",
		},
	}

	// we create our commands
	app.Commands = []cli.Command{
		{
			Name:      "generate",
			ShortName: "g",
			Usage:     "Generate files yaml",
			Flags:     flags,
			Action:    GenerateYamlByTemplate,
		},
	}

	return app

}

func GenerateYamlByTemplate(c *cli.Context) (err error) {

	filePathTemplates := "./templates"

	params := Params{
		AppName: c.String("name"),
		Image:   c.String("image"),
		Port:    c.String("port"),
		Replica: c.String("replica"),
	}

	var templates *template.Template
	var allFiles []string
	files, err := ioutil.ReadDir(filePathTemplates)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		filename := file.Name()
		fullPath := filePathTemplates + "/" + filename
		if strings.HasSuffix(filename, ".tmpl") {
			allFiles = append(allFiles, fullPath)
		}
	}

	templates, err = template.ParseFiles(allFiles...)
	if err != nil {
		fmt.Println(err)
	}

	if err := CreateYamlFile(&params, templates, "deployment"); err != nil {
		log.Fatal(err)
	}

	if err := CreateYamlFile(&params, templates, "service"); err != nil {
		log.Fatal(err)
	}

	return nil
}

func CreateYamlFile(params *Params, templates *template.Template, typeTmpl string) (err error) {

	s1 := templates.Lookup(typeTmpl + ".tmpl")
	s1.ExecuteTemplate(os.Stdout, typeTmpl+".yml", params)

	outputPath := filepath.Join("./k8s/", typeTmpl+".yml")

	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = s1.Execute(f, params)
	if err != nil {
		panic(err)
	}

	return nil
}
