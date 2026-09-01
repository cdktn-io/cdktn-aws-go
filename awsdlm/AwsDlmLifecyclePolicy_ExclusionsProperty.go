package awsdlm


// Experimental.
type AwsDlmLifecyclePolicy_ExclusionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#exclude_boot_volumes AwsDlmLifecyclePolicy#exclude_boot_volumes}.
	// Experimental.
	ExcludeBootVolumes interface{} `field:"optional" json:"excludeBootVolumes" yaml:"excludeBootVolumes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#exclude_tags AwsDlmLifecyclePolicy#exclude_tags}.
	// Experimental.
	ExcludeTags *map[string]*string `field:"optional" json:"excludeTags" yaml:"excludeTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#exclude_volume_types AwsDlmLifecyclePolicy#exclude_volume_types}.
	// Experimental.
	ExcludeVolumeTypes *[]*string `field:"optional" json:"excludeVolumeTypes" yaml:"excludeVolumeTypes"`
}

