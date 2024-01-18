package v1alpha3

import "fmt"

// ManifestLabel returns the cluster.x-k8s.io/provider label value for a provider/type.
//
// Note: the label uniquely describes the provider type and its kind (e.g. bootstrap-kubeadm);
// it's not meant to be used to describe each instance of a particular provider.
func ManifestLabel(name string, providerType ProviderType) string {
	switch providerType {
	case BootstrapProviderType:
		return fmt.Sprintf("bootstrap-%s", name)
	case ControlPlaneProviderType:
		return fmt.Sprintf("control-plane-%s", name)
	case InfrastructureProviderType:
		return fmt.Sprintf("infrastructure-%s", name)
	case IPAMProviderType:
		return fmt.Sprintf("ipam-%s", name)
	case RuntimeExtensionProviderType:
		return fmt.Sprintf("runtime-extension-%s", name)
	case AddonProviderType:
		return fmt.Sprintf("addon-%s", name)
	default:
		return name
	}
}
