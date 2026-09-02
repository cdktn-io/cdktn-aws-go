package awsdlm


// Experimental.
type TfLifecyclePolicy_ArchiveRuleProperty struct {
	// archive_retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#archive_retain_rule TfLifecyclePolicy#archive_retain_rule}
	// Experimental.
	ArchiveRetainRule *TfLifecyclePolicy_ArchiveRetainRuleProperty `field:"required" json:"archiveRetainRule" yaml:"archiveRetainRule"`
}

