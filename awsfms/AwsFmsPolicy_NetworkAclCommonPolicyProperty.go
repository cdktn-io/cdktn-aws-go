package awsfms


// Experimental.
type AwsFmsPolicy_NetworkAclCommonPolicyProperty struct {
	// network_acl_entry_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#network_acl_entry_set AwsFmsPolicy#network_acl_entry_set}
	// Experimental.
	NetworkAclEntrySet *AwsFmsPolicy_NetworkAclEntrySetProperty `field:"optional" json:"networkAclEntrySet" yaml:"networkAclEntrySet"`
}

