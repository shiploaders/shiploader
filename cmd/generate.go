package command

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/urfave/cli"
	"log"
	"os"
	"shiploader/apis/apps"
	"shiploader/pkg"
	"shiploader/pkg/generators"
	"sync"
)



const (
	DeploymentFileName = "deployment.yaml"
	ServiceFileName = "service.yaml"
)
var (
	validate = validator.New()
)

// Will return the command line ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Shiploader"
	app.Usage = "TODO: Add usage here"
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "A config file that defines your desired state of your apps",
			Required: true,
			Value: "config.yaml",
		},
		cli.StringFlag{
			Name:  "dest",
			Usage: fmt.Sprintf("The destination directory where generated files will be saved to."),
			Value: currentDir,
		},
	}

	// we create our commands
	app.Commands = []cli.Command{
		// TODO: perhaps we create a command that does "generate + check-in-git + deploy to a cluster(take inputs from the config + kubeconfig)"
		{
			Name:      "generate",
			ShortName: "g",
			Usage:     "Generate files yaml",
			Flags:     flags,
			Action:    GenerateYamlForApps,
		},
	}
	return app
}

func GenerateYamlForApps(c *cli.Context) error {
	allApps, err := pkg.ConfigToApps(c.String("config"))
	if err != nil {
		log.Fatalf("Failed to parse config file [%v]: %v", c.String("config"),err)
	}
	destDir := c.String("dest")
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		if errMakingDir := os.MkdirAll(destDir, 0777); errMakingDir != nil {
			return errMakingDir
		}
	}

	var wg sync.WaitGroup
	wr := generators.Writer{}
	wg.Add(len(allApps.Apps))
	for _, app := range allApps.Apps {

		if errValidatingAppFields := validate.Struct(app); errValidatingAppFields != nil {
			log.Fatal(errValidatingAppFields)
		}

		go func(a apps.App) {
			defer wg.Done()
			// run generators
			if err := generators.GenerateDeployment(a, destDir, &wr); err != nil {
				log.Fatal(err)
			}
			if err := generators.GenerateSvc(a, destDir, &wr); err != nil {
				log.Fatal(err)
			}
		}(app)
	}
	wg.Wait()
	return nil
}