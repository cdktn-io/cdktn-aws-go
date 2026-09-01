package awsfms


// Experimental.
type AwsFmsPolicy_PolicyOptionProperty struct {
	// network_acl_common_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#network_acl_common_policy AwsFmsPolicy#network_acl_common_policy}
	// Experimental.
	NetworkAclCommonPolicy *AwsFmsPolicy_NetworkAclCommonPolicyProperty `field:"optional" json:"networkAclCommonPolicy" yaml:"networkAclCommonPolicy"`
	// network_firewall_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#network_firewall_policy AwsFmsPolicy#network_firewall_policy}
	// Experimental.
	NetworkFirewallPolicy *AwsFmsPolicy_NetworkFirewallPolicyProperty `field:"optional" json:"networkFirewallPolicy" yaml:"networkFirewallPolicy"`
	// third_party_firewall_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#third_party_firewall_policy AwsFmsPolicy#third_party_firewall_policy}
	// Experimental.
	ThirdPartyFirewallPolicy *AwsFmsPolicy_ThirdPartyFirewallPolicyProperty `field:"optional" json:"thirdPartyFirewallPolicy" yaml:"thirdPartyFirewallPolicy"`
}

