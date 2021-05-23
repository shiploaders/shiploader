package command

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"shiploader/apis"
	"sigs.k8s.io/yaml"
	"sync"
)

var (
	validate = validator.New()
)

const (
	DeploymentFileName = "deployment.yaml"
	ServiceFileName = "service.yaml"
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
	allApps, err := configToApps(c.String("config"))
	if err != nil {
		log.Fatalf("Failed to parse config file [%v]: %v", c.String("config"),err)
	}

	var wg sync.WaitGroup
	wg.Add(len(allApps.Apps))
	for _, app := range allApps.Apps {
		go func(a apis.App) {
			defer wg.Done()
			if errValidatingAppFields := validate.Struct(a); errValidatingAppFields != nil {
				log.Fatal(errValidatingAppFields)
			}
			// run generators
			if err := generateDeployment(a, c.String("dest")); err != nil {
				log.Fatal(err)
			}
			if err := generateSvc(a, c.String("dest")); err != nil {
				log.Fatal(err)
			}
		}(app)
	}
	wg.Wait()
	return nil
}

// TODO: -- everything under this TODO should go in a separate pkg and should be importable by the above func
func generateDeployment(app apis.App, dest string) error {
	fileName := filepath.Join(dest, fmt.Sprintf("%s-%s-%s", app.Namespace, app.Name, DeploymentFileName))
	rawDeploy, errMarshallingToDeploy := yaml.Marshal(app.ToDeployment())
	if errMarshallingToDeploy != nil {
		return errMarshallingToDeploy
	}
	return ioutil.WriteFile(fileName, rawDeploy, os.ModePerm)
}

func generateSvc(app apis.App, dest string) error {
	fileName := filepath.Join(dest, fmt.Sprintf("%s-%s-%s", app.Namespace, app.Name, ServiceFileName))
	rawSvc, errMarshallingToSvc := yaml.Marshal(app.ToService())
	if errMarshallingToSvc != nil {
		return errMarshallingToSvc
	}
	return ioutil.WriteFile(fileName, rawSvc, os.ModePerm)
}

func configToApps(configFile string) (*apis.Apps, error) {
	apps := &apis.Apps{}
	rawConfigFile, errReading := os.ReadFile(configFile)
	if errReading != nil {
		return nil, errReading
	}
	if err := yaml.Unmarshal(rawConfigFile, apps); err != nil {
		return nil, err
	}
	if errValidating := validate.Struct(apps); errValidating != nil {
		return nil, errValidating
	}
	return apps, nil
}