package awsdlm


// Experimental.
type AwsDlmLifecyclePolicy_ArchiveRetainRuleProperty struct {
	// retention_archive_tier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#retention_archive_tier AwsDlmLifecyclePolicy#retention_archive_tier}
	// Experimental.
	RetentionArchiveTier *AwsDlmLifecyclePolicy_RetentionArchiveTierProperty `field:"required" json:"retentionArchiveTier" yaml:"retentionArchiveTier"`
}

