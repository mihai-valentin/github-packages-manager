package github

import (
	"context"
	"github.com/google/go-github/v43/github"
	"net/url"
	"sort"
	"strings"
)

type Context struct {
	client  *github.Client
	context context.Context
}

func NewAuthorizedContext(login string, token string) *Context {
	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(login),
		Password: strings.TrimSpace(token),
	}

	return &Context{
		client:  github.NewClient(tp.Client()),
		context: context.Background(),
	}
}

type OrganizationInterface interface {
	GetName() string
	GetPackageType() string
	GetKeepVersions() int
}

func (c *Context) GetOrganizationPackages(o OrganizationInterface) ([]*github.Package, error) {
	packageType := o.GetPackageType()
	options := &github.PackageListOptions{PackageType: &packageType}

	organizationPackages, _, err := c.client.Organizations.ListPackages(c.context, o.GetName(), options)
	if err != nil {
		return nil, err
	}

	return organizationPackages, nil
}

func (c *Context) GetPackageVersions(o OrganizationInterface, packageName string) ([]*github.PackageVersion, error) {
	escapedPackageName := url.QueryEscape(packageName)
	packageVersions, _, err := c.client.Organizations.PackageGetAllVersions(
		c.context,
		o.GetName(),
		o.GetPackageType(),
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

func (c *Context) CleanUpPackageVersions(
	o OrganizationInterface,
	packageName string,
	packageVersions []*github.PackageVersion,
) error {
	escapedPackageName := url.QueryEscape(packageName)
	packageVersionsCount := len(packageVersions)
	for i := o.GetKeepVersions(); i < packageVersionsCount; i++ {
		_, err := c.client.Organizations.PackageDeleteVersion(
			c.context,
			o.GetName(),
			o.GetPackageType(),
			escapedPackageName,
			packageVersions[i].GetID(),
		)

		if err != nil {
			return err
		}
	}

	return nil
}
