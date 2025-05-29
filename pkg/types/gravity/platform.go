package gravity

const Name = "gravity"

// NetworkConfig stores all the network details of an LPAR.
type NetworkConfig struct {
	IP      string `json:"ip"`
	Gateway string `json:"gateway"`
	Netmask string `json:"netmask"`
}

// Host stores all the configuration data for an LPAR.
type Host struct {
	Name          string         `json:"name,omitempty"`
	Role          string         `json:"role,omitempty"`
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`
	CPU           string         `json:"vcpu,omitempty"`
	Memory        string         `json:"memory,omitempty"`
}

// Platform stores all the required details of booting an LPAR as a node.
type Platform struct {
	ExternalDNSIP       string  `json:"externalDNSIP"`
	LoadBalancerVIP     string  `json:"loadBalancerVIP"`
	NetworkType         string  `json:"networkType"`
	DiskType            string  `json:"diskType"`
	ControlNodesProfile string  `json:"controlNodesProfile"`
	ComputeNodesProfile string  `json:"computeNodesProfile"`
	DHCP                bool    `json:"dhcp"`
	Hosts               []*Host `json:"hosts,omitempty"`
}
