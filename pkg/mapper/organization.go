package mapper

type SkipList struct {
	items []string
}

func newSkipList(items []string) *SkipList {
	return &SkipList{items}
}

func (l *SkipList) Contains(item string) bool {
	for _, i := range l.items {
		if item == i {
			return true
		}
	}

	return false
}

type Organization struct {
	Name         string
	PackageType  string
	KeepVersions int
	SkipList     *SkipList
}

func NewOrganization(
	name string,
	packageType string,
	keepVersions int,
	skipList []string,
) *Organization {
	return &Organization{
		Name:         name,
		PackageType:  packageType,
		KeepVersions: keepVersions,
		SkipList:     newSkipList(skipList),
	}
}

func (o *Organization) GetName() string {
	return o.Name
}

func (o *Organization) GetPackageType() string {
	return o.PackageType
}

func (o *Organization) GetKeepVersions() int {
	return o.KeepVersions
}
