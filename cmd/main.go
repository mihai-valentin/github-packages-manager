package main

import (
	"github.com/joho/godotenv"
	"github.com/mihai-valentin/github-packages-manager/pkg/github"
	"github.com/mihai-valentin/github-packages-manager/pkg/mapper"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot load .env file")
	}

	config := mapper.NewConfig()
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}
	organizations := config.GetOrganisations()

	authorizedContext := github.NewAuthorizedContext(os.Getenv("GITHUB_LOGIN"), os.Getenv("GITHUB_TOKEN"))

	for _, organization := range organizations {
		organizationPackages, err := authorizedContext.GetOrganizationPackages(organization)
		if err != nil {
			log.Fatal(err)
		}

		for _, organizationPackage := range organizationPackages {
			packageName := organizationPackage.GetName()
			if organization.SkipList.Contains(packageName) {
				continue
			}

			packageVersions, err := authorizedContext.GetPackageVersions(organization, packageName)
			if err != nil {
				log.Fatal(err)
			}

			if err := authorizedContext.CleanUpPackageVersions(organization, packageName, packageVersions); err != nil {
				log.Fatal(err)
			}
		}
	}
}
