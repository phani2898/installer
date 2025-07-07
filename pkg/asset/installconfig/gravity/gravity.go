package gravity

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/openshift/installer/pkg/types/gravity"
)

// Gathers all the inputs required for LPAR installation on gravity platform
func Platform() (*gravity.Platform, error) {

	lbVIP, err := getVIP("Load Balancer VIP", "Virtual IP address for the OpenShift Loadbalancer")
	if err != nil {
		return nil, err
	}

	dnsVIP, err := getVIP("DNS VIP", "IP address of the external DNS Server")
	if err != nil {
		return nil, err
	}

	networkType, err := selectNetworkType()
	if err != nil {
		return nil, fmt.Errorf("failed to survey desired network type: %w", err)
	}

	diskType, err := selectDiskType()
	if err != nil {
		return nil, fmt.Errorf("failed to survey desired disk type: %w", err)
	}

	enableDHCP, err := selectDHCP()
	if err != nil {
		return nil, fmt.Errorf("failed to survey dhcp enablement: %w", err)
	}

	controlNodesProfile, err := selectNodesProfile("Control")
	if err != nil {
		return nil, fmt.Errorf("failed to survey control nodes profile: %w", err)
	}

	computeNodesProfile, err := selectNodesProfile("Compute")
	if err != nil {
		return nil, fmt.Errorf("failed to survey compute nodes profile: %w", err)
	}

	var hosts []*gravity.Host

	gravityInfo := &gravity.Platform{
		DNSVIP:              dnsVIP,
		LoadBalancerVIP:     lbVIP,
		NetworkType:         networkType,
		DiskType:            diskType,
		ControlNodesProfile: controlNodesProfile,
		ComputeNodesProfile: computeNodesProfile,
		DHCP:                enableDHCP,
		Hosts:               hosts,
	}

	return gravityInfo, nil
}

func getVIP(msg string, help string) (ipAddress string, err error) {

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: msg,
				Help:    help,
			},
			Validate: survey.Required,
		},
	}, &ipAddress)

	if err != nil {
		return "", fmt.Errorf("failed to get %s: %w", msg, err)
	}

	return ipAddress, nil
}

func selectNetworkType() (networkType string, err error) {
	var networkTypes = []string{"RoCE", "Hipersockets"}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Select{
				Message: "Network Type",
				Help:    "Network type to be used for LPAR installation.",
				Default: "RoCE",
				Options: networkTypes,
			},
			Validate: survey.Required,
		},
	}, &networkType)

	if err != nil {
		return "", err
	}
	return networkType, err
}

func selectDiskType() (diskType string, err error) {
	var diskTypes = []string{"FCP", "DASD", "NVMe"}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Select{
				Message: "Disk Type",
				Help:    "Disk type to be used for LPAR installation.",
				Default: "DASD",
				Options: diskTypes,
			},
			Validate: survey.Required,
		},
	}, &diskType)

	if err != nil {
		return "", err
	}
	return diskType, err
}

func selectDHCP() (enableDHCP bool, err error) {
	var dhcpOptions = []string{"true", "false"}
	var selected string

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Select{
				Message: "Enable DHCP?",
				Help:    "Choose whether DHCP should be enabled for LPAR installation.",
				Default: "true",
				Options: dhcpOptions,
			},
			Validate: survey.Required,
		},
	}, &selected)

	if err != nil {
		return true, err
	}

	if selected == "true" {
		enableDHCP = true
	} else {
		enableDHCP = false
	}
	return enableDHCP, nil
}

func selectNodesProfile(nodeType string) (nodeProfile string, err error) {
	var profileOptions = []string{"4x16", "8x16", "8x32", "16x32", "16x64"}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Select{
				Message: fmt.Sprintf("%s Nodes Profile", nodeType),
				Help:    fmt.Sprintf("Chose the desired vCPU and memory profile for %s nodes", nodeType),
				Default: "4x16",
				Options: profileOptions,
			},
			Validate: survey.Required,
		},
	}, &nodeProfile)

	if err != nil {
		return "", err
	}

	return nodeProfile, nil
}
