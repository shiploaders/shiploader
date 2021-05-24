package generators

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"shiploader/apis/apps"
	"sigs.k8s.io/yaml"
)

func GenerateSvc(app apps.App, dest, svcFileName string) error {
	fileName := filepath.Join(dest, fmt.Sprintf("%s-%s-%s", app.Namespace, app.Name, svcFileName))
	rawSvc, errMarshallingToSvc := yaml.Marshal(app.ToService())
	if errMarshallingToSvc != nil {
		return errMarshallingToSvc
	}
	return ioutil.WriteFile(fileName, rawSvc, os.ModePerm)
}
