package main

import (
	"github.com/joho/godotenv"
	"github.com/mihai-valentin/github-packages-manager/pkg/action"
	"github.com/mihai-valentin/github-packages-manager/pkg/context"
	"github.com/mihai-valentin/github-packages-manager/pkg/mapper"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot load .env file")
	}

	config, err := mapper.NewConfig(os.Getenv("CONFIG_PATH"), os.Getenv("CONFIG_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	organizationContext := context.NewOrganizationContext(
		os.Getenv("GITHUB_LOGIN"),
		os.Getenv("GITHUB_TOKEN"),
	)

	new(action.CleanupPackagesVersionsAction).Handle(config.GetOrganisations(), organizationContext)
}
