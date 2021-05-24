package generators

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"shiploader/apis/apps"
	"sigs.k8s.io/yaml"
)

func GenerateDeployment(app apps.App, dest, deployFileName string) error {
	fileName := filepath.Join(dest, fmt.Sprintf("%s-%s-%s", app.Namespace, app.Name, deployFileName))
	rawDeploy, errMarshallingToDeploy := yaml.Marshal(app.ToDeployment())
	if errMarshallingToDeploy != nil {
		return errMarshallingToDeploy
	}
	return ioutil.WriteFile(fileName, rawDeploy, os.ModePerm)
}
