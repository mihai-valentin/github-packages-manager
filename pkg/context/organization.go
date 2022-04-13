package context

import (
	"github.com/google/go-github/v43/github"
	"github.com/mihai-valentin/github-packages-manager/pkg/contract"
	"net/url"
	"sort"
)

type OrganizationContext struct {
	*AuthorizedContext
	organization contract.OrganizationInterface
}

func NewOrganizationContext(login string, token string) *OrganizationContext {
	return &OrganizationContext{
		AuthorizedContext: NewAuthorizedContext(login, token),
	}
}

func (c *OrganizationContext) SetOrganization(o contract.OrganizationInterface) {
	c.organization = o
}

func (c *OrganizationContext) ShouldSkipPackage(packageName string) bool {
	return c.organization.IsInSkipList(packageName)
}

func (c *OrganizationContext) GetOrganizationPackages() ([]*github.Package, error) {
	packageType := c.organization.GetPackageType()
	options := &github.PackageListOptions{PackageType: &packageType}

	organizationPackages, _, err := c.client.Organizations.ListPackages(c.context, c.organization.GetName(), options)
	if err != nil {
		return nil, err
	}

	return organizationPackages, nil
}

func (c *OrganizationContext) GetPackageVersions(packageName string) ([]*github.PackageVersion, error) {
	escapedPackageName := url.QueryEscape(packageName)
	packageVersions, _, err := c.client.Organizations.PackageGetAllVersions(
		c.context,
		c.organization.GetName(),
		c.organization.GetPackageType(),
		escapedPackageName,
		nil,
	)

	if err != nil {
		return nil, err
	}

	sort.Slice(packageVersions, func(i, j int) bool {
		return packageVersions[i].CreatedAt.After(packageVersions[j].CreatedAt.Time)
	})

	return packageVersions, err
}

func (c *OrganizationContext) CleanUpPackageVersions(
	packageName string,
	packageVersions []*github.PackageVersion,
) error {
	keepVersions := c.organization.GetKeepVersions()
	escapedPackageName := url.QueryEscape(packageName)
	for _, packageVersion := range packageVersions[keepVersions:] {
		_, err := c.client.Organizations.PackageDeleteVersion(
			c.context,
			c.organization.GetName(),
			c.organization.GetPackageType(),
			escapedPackageName,
			packageVersion.GetID(),
		)

		if err != nil {
			return err
		}
	}

	return nil
}
