package v1alpha2

// Kubeark contains application configuration
type Kubeark struct {
	IngressHost string   `yaml:"ingressHost" json:"ingressHost,omitempty"`
	AcmeEmail   string   `yaml:"acmeEmail" json:"acmeEmail,omitempty"`
	Storage     string   `yaml:"storage" json:"storage,omitempty"`
	Rook        Rook     `yaml:"rook" json:"sources"`
	Postgres    Postgres `yaml:"postgres" json:"postgres"`
}

type Rook struct {
	MonCount         int `yaml:"monCount" json:"monCount,omitempty"`
	MgrCount         int `yaml:"mgrCount" json:"mgrCount,omitempty"`
	MetadataPoolSize int `yaml:"metadataPoolSize" json:"metadataPoolSize,omitempty"`
	DataPoolSize     int `yaml:"dataPoolSize" json:"dataPoolSize,omitempty"`
}

type Postgres struct {
	InstanceStorage string `yaml:"instanceStorage" json:"instanceStorage,omitempty"`
	BackupStorage   string `yaml:"backupStorage" json:"backupStorage,omitempty"`
}
