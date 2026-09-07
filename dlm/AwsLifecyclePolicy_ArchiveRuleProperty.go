package dlm


// Experimental.
type AwsLifecyclePolicy_ArchiveRuleProperty struct {
	// archive_retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#archive_retain_rule AwsLifecyclePolicy#archive_retain_rule}
	// Experimental.
	ArchiveRetainRule *AwsLifecyclePolicy_ArchiveRetainRuleProperty `field:"required" json:"archiveRetainRule" yaml:"archiveRetainRule"`
}

