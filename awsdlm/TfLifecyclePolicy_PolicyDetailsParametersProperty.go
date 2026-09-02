package awsdlm


// Experimental.
type TfLifecyclePolicy_PolicyDetailsParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#exclude_boot_volume TfLifecyclePolicy#exclude_boot_volume}.
	// Experimental.
	ExcludeBootVolume interface{} `field:"optional" json:"excludeBootVolume" yaml:"excludeBootVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#exclude_data_volume_tags TfLifecyclePolicy#exclude_data_volume_tags}.
	// Experimental.
	ExcludeDataVolumeTags *map[string]*string `field:"optional" json:"excludeDataVolumeTags" yaml:"excludeDataVolumeTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#no_reboot TfLifecyclePolicy#no_reboot}.
	// Experimental.
	NoReboot interface{} `field:"optional" json:"noReboot" yaml:"noReboot"`
}

