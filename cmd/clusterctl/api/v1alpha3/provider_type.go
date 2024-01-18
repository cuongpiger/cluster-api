package v1alpha3

// ****************************************************** CONSTS ******************************************************

const (
	// CoreProviderType is a type reserved for Cluster API core repository.
	CoreProviderType = ProviderType("CoreProvider")

	// BootstrapProviderType is the type associated with codebases that provide
	// bootstrapping capabilities.
	BootstrapProviderType = ProviderType("BootstrapProvider")

	// InfrastructureProviderType is the type associated with codebases that provide
	// infrastructure capabilities.
	InfrastructureProviderType = ProviderType("InfrastructureProvider")

	// ControlPlaneProviderType is the type associated with codebases that provide
	// control-plane capabilities.
	ControlPlaneProviderType = ProviderType("ControlPlaneProvider")

	// IPAMProviderType is the type associated with codebases that provide
	// IPAM capabilities.
	IPAMProviderType = ProviderType("IPAMProvider")

	// RuntimeExtensionProviderType is the type associated with codebases that provide
	// runtime extensions.
	RuntimeExtensionProviderType = ProviderType("RuntimeExtensionProvider")

	// AddonProviderType is the type associated with codebases that provide
	// add-on capabilities.
	AddonProviderType = ProviderType("AddonProvider")

	// ProviderTypeUnknown is used when the type is unknown.
	ProviderTypeUnknown = ProviderType("")
)

// ****************************************************** OBJECTS ******************************************************

// ________________________________________________________________________________________________________ ProviderType

// ProviderType is a string representation of a Provider type.
type ProviderType string

// Order return an integer that can be used to sort ProviderType values.
func (p ProviderType) Order() int {
	switch p {
	case CoreProviderType:
		return 0
	case BootstrapProviderType:
		return 1
	case ControlPlaneProviderType:
		return 2
	case InfrastructureProviderType:
		return 3
	case IPAMProviderType:
		return 4
	case RuntimeExtensionProviderType:
		return 5
	case AddonProviderType:
		return 6
	default:
		return 99
	}
}
