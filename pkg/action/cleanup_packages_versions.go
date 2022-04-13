package action

import "github.com/mihai-valentin/github-packages-manager/pkg/contract"

type CleanupPackagesVersionsAction struct {
}

func (a *CleanupPackagesVersionsAction) Handle(
	organizations contract.OrganizationsListInterface,
	organizationContext contract.OrganizationContextInterface,
) {
	organizations.Each(func(organization contract.OrganizationInterface) error {
		organizationContext.SetOrganization(organization)
		if err := a.cleanupVersions(organizationContext); err != nil {
			return err
		}

		return nil
	})
}

func (a *CleanupPackagesVersionsAction) cleanupVersions(oc contract.OrganizationContextInterface) error {
	organizationPackages, err := oc.GetOrganizationPackages()
	if err != nil {
		return err
	}

	for _, organizationPackage := range organizationPackages {
		packageName := organizationPackage.GetName()
		if oc.ShouldSkipPackage(packageName) {
			continue
		}

		packageVersions, err := oc.GetPackageVersions(packageName)
		if err != nil {
			return err
		}

		if err := oc.CleanUpPackageVersions(packageName, packageVersions); err != nil {
			return err
		}
	}

	return nil
}
