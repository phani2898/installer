package gravity

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/openshift/installer/pkg/types/gravity"
)

// Gathers all the inputs required for LPAR installation on gravity platform
func Platform() (*gravity.Platform, error) {

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

	var hosts []*gravity.Host

	gravityInfo := &gravity.Platform{
		ExternalDNSIP:  "",
		LoadBalancerIP: "",
		NetworkType:    networkType,
		DiskType:       diskType,
		DHCP:           enableDHCP,
		Hosts:          hosts,
	}

	return gravityInfo, nil
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
