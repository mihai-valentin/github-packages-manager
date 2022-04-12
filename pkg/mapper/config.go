package mapper

import "github.com/spf13/viper"

const (
	configPath = "./"
	configName = "mapper"
)

type Config struct {
	path string
	name string
}

func NewConfig() *Config {
	return &Config{
		path: configPath,
		name: configName,
	}
}

func (c *Config) Init() error {
	viper.AddConfigPath(c.path)
	viper.SetConfigName(c.name)

	return viper.ReadInConfig()
}

func (c *Config) GetOrganisations() []*Organization {
	var organizations []*Organization

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
