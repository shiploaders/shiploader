package generators

import (
	"fmt"
	"path/filepath"
	"shiploader/apis/apps"
	"sigs.k8s.io/yaml"
)

const (
	DeploymentFileName = "deployment.yaml"
)

func GenerateDeployment(app apps.App, dest string, w WriterInterface) error {
	fileName := filepath.Join(dest, fmt.Sprintf("%s-%s-%s", app.Namespace, app.Name, DeploymentFileName))
	rawDeploy, errMarshallingToDeploy := yaml.Marshal(app.ToDeployment())
	if errMarshallingToDeploy != nil {
		return errMarshallingToDeploy
	}
	return w.WriteFile(fileName, rawDeploy, 0666)
}
