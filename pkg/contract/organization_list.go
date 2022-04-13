package contract

type HandlerOrganizationFunction func(OrganizationInterface) error

type OrganizationsListInterface interface {
	Each(handler HandlerOrganizationFunction)
}
