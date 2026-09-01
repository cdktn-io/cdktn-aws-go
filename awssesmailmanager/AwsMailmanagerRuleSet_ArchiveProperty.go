package awssesmailmanager


// Experimental.
type AwsMailmanagerRuleSet_ArchiveProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#target_archive AwsMailmanagerRuleSet#target_archive}.
	// Experimental.
	TargetArchive *string `field:"required" json:"targetArchive" yaml:"targetArchive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#action_failure_policy AwsMailmanagerRuleSet#action_failure_policy}.
	// Experimental.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
}

