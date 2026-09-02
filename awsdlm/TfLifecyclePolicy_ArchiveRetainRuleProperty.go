package awsdlm


// Experimental.
type TfLifecyclePolicy_ArchiveRetainRuleProperty struct {
	// retention_archive_tier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#retention_archive_tier TfLifecyclePolicy#retention_archive_tier}
	// Experimental.
	RetentionArchiveTier *TfLifecyclePolicy_RetentionArchiveTierProperty `field:"required" json:"retentionArchiveTier" yaml:"retentionArchiveTier"`
}

