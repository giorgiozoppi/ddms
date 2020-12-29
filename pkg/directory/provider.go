package directory

// Provider is the provider for a directory service
type Provider interface {
	GetName() string
	CreateSingleValueDirectory(configuration Configuration) (SingleValueDirectory, error)
	CreateMultipleValueDirectory(configuration Configuration) (MultiValueDirectory, error)
	RemoveDirectory(path string) error
}
