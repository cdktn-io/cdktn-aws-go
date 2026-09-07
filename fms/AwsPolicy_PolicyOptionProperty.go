package fms


// Experimental.
type AwsPolicy_PolicyOptionProperty struct {
	// network_acl_common_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#network_acl_common_policy AwsPolicy#network_acl_common_policy}
	// Experimental.
	NetworkAclCommonPolicy *AwsPolicy_NetworkAclCommonPolicyProperty `field:"optional" json:"networkAclCommonPolicy" yaml:"networkAclCommonPolicy"`
	// network_firewall_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#network_firewall_policy AwsPolicy#network_firewall_policy}
	// Experimental.
	NetworkFirewallPolicy *AwsPolicy_NetworkFirewallPolicyProperty `field:"optional" json:"networkFirewallPolicy" yaml:"networkFirewallPolicy"`
	// third_party_firewall_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#third_party_firewall_policy AwsPolicy#third_party_firewall_policy}
	// Experimental.
	ThirdPartyFirewallPolicy *AwsPolicy_ThirdPartyFirewallPolicyProperty `field:"optional" json:"thirdPartyFirewallPolicy" yaml:"thirdPartyFirewallPolicy"`
}

