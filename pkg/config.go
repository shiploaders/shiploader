package pkg

import (
	"github.com/go-playground/validator"
	"os"
	"shiploader/apis/apps"
	"sigs.k8s.io/yaml"
)
var (
	validate = validator.New()
)
func ConfigToApps(configFile string) (*apps.Apps, error) {
	a := &apps.Apps{}
	rawConfigFile, errReading := os.ReadFile(configFile)
	if errReading != nil {
		return nil, errReading
	}
	if err := yaml.Unmarshal(rawConfigFile, a); err != nil {
		return nil, err
	}
	if errValidating := validate.Struct(a); errValidating != nil {
		return nil, errValidating
	}
	return a, nil
}