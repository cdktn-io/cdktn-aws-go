package fms


// Experimental.
type AwsPolicy_NetworkAclEntrySetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#force_remediate_for_first_entries AwsPolicy#force_remediate_for_first_entries}.
	// Experimental.
	ForceRemediateForFirstEntries interface{} `field:"required" json:"forceRemediateForFirstEntries" yaml:"forceRemediateForFirstEntries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#force_remediate_for_last_entries AwsPolicy#force_remediate_for_last_entries}.
	// Experimental.
	ForceRemediateForLastEntries interface{} `field:"required" json:"forceRemediateForLastEntries" yaml:"forceRemediateForLastEntries"`
	// first_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#first_entry AwsPolicy#first_entry}
	// Experimental.
	FirstEntry interface{} `field:"optional" json:"firstEntry" yaml:"firstEntry"`
	// last_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#last_entry AwsPolicy#last_entry}
	// Experimental.
	LastEntry interface{} `field:"optional" json:"lastEntry" yaml:"lastEntry"`
}

