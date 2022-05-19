package generators

import (
	"fmt"
	"path/filepath"
	"shiploader/apis/apps"
	"sigs.k8s.io/yaml"
)

const (
	ServiceFileName = "service.yaml"
)

func GenerateSvc(app apps.App, dest string, w WriterInterface) error {
	fileName := filepath.Join(dest, fmt.Sprintf("%s-%s-%s", app.Namespace, app.Name, ServiceFileName))
	rawSvc, errMarshallingToSvc := yaml.Marshal(app.ToService())
	if errMarshallingToSvc != nil {
		return errMarshallingToSvc
	}
	return w.WriteFile(fileName, rawSvc, 0666)
}
