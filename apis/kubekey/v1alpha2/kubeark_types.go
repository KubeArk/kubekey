package v1alpha2

// Kubeark contains application configuration
type Kubeark struct {
	IngressHost string `yaml:"ingressHost" json:"ingressHost,omitempty"`
}
