package apis

type Apps struct {
	Apps []App `yaml:"apps" validate:"required"`
}

type App struct {
	Name      string `yaml:"name" validate:"required"`
	Namespace string `yaml:"namespace" validate:"required"`
	Image     string `yaml:"image" validate:"required"`
	Replicas  int32  `yaml:"replicas"`
	Port      int32  `yaml:"port" validate:"required"`
}

