package dlm


// Experimental.
type AwsLifecyclePolicy_PolicyDetailsParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#exclude_boot_volume AwsLifecyclePolicy#exclude_boot_volume}.
	// Experimental.
	ExcludeBootVolume interface{} `field:"optional" json:"excludeBootVolume" yaml:"excludeBootVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#exclude_data_volume_tags AwsLifecyclePolicy#exclude_data_volume_tags}.
	// Experimental.
	ExcludeDataVolumeTags *map[string]*string `field:"optional" json:"excludeDataVolumeTags" yaml:"excludeDataVolumeTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#no_reboot AwsLifecyclePolicy#no_reboot}.
	// Experimental.
	NoReboot interface{} `field:"optional" json:"noReboot" yaml:"noReboot"`
}

