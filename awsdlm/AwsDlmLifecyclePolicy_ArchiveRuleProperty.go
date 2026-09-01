package awsdlm


// Experimental.
type AwsDlmLifecyclePolicy_ArchiveRuleProperty struct {
	// archive_retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#archive_retain_rule AwsDlmLifecyclePolicy#archive_retain_rule}
	// Experimental.
	ArchiveRetainRule *AwsDlmLifecyclePolicy_ArchiveRetainRuleProperty `field:"required" json:"archiveRetainRule" yaml:"archiveRetainRule"`
}

