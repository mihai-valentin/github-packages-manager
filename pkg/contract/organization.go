package contract

type OrganizationInterface interface {
	GetName() string
	GetPackageType() string
	GetKeepVersions() int
	IsInSkipList(packageName string) bool
}
