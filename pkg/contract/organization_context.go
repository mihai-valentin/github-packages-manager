package contract

import "github.com/google/go-github/v43/github"

type OrganizationContextInterface interface {
	SetOrganization(OrganizationInterface)
	GetOrganizationPackages() ([]*github.Package, error)
	ShouldSkipPackage(packageName string) bool
	GetPackageVersions(packageName string) ([]*github.PackageVersion, error)
	CleanUpPackageVersions(packageName string, versions []*github.PackageVersion) error
}
