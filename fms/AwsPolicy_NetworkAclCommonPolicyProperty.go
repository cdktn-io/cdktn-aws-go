package fms


// Experimental.
type AwsPolicy_NetworkAclCommonPolicyProperty struct {
	// network_acl_entry_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#network_acl_entry_set AwsPolicy#network_acl_entry_set}
	// Experimental.
	NetworkAclEntrySet *AwsPolicy_NetworkAclEntrySetProperty `field:"optional" json:"networkAclEntrySet" yaml:"networkAclEntrySet"`
}

