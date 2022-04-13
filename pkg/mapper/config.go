package mapper

import (
	"github.com/mihai-valentin/github-packages-manager/pkg/contract"
	"github.com/spf13/viper"
)

type OrganizationList []*Organization

func (l OrganizationList) Each(handler contract.HandlerOrganizationFunction) {
	for _, organization := range l {
		if err := handler(organization); err != nil {

		}
	}
}

type Config struct {
	path string
	name string
}

func NewConfig(path string, name string) (*Config, error) {
	config := &Config{
		path: path,
		name: name,
	}

	viper.AddConfigPath(config.path)
	viper.SetConfigName(config.name)

	return config, viper.ReadInConfig()
}

func (c *Config) GetOrganisations() OrganizationList {
	var organizations OrganizationList

	organizationsConfig := viper.GetStringMap("organizations")
	for name, _ := range organizationsConfig {
		packageConfig := NewOrganization(
			name,
			viper.GetString("organizations."+name+".package_type"),
			viper.GetInt("organizations."+name+".keep_versions"),
			viper.GetStringSlice("organizations."+name+".skip_packages"),
		)

		organizations = append(organizations, packageConfig)
	}

	return organizations
}
