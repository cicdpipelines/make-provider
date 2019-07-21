package make

import (
	cicd "github.com/NoUseFreak/cicd/pkg/helper"
)

// Provider returns cicd configuration for make
func Provider() *cicd.Provider {
	return &cicd.Provider{
		// TODO
		ProviderSchema: map[string]*cicd.ProviderSchema{
			"schema_attr_name": {
				Type:        cicd.TypeString,
				Description: "describing schema_attr_name",
			},
		},
		StepMap: map[string]*cicd.Step{
			"make": stepMake(),
		},
	}
}
