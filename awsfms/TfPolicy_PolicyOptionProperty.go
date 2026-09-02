package awsfms


// Experimental.
type TfPolicy_PolicyOptionProperty struct {
	// network_acl_common_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#network_acl_common_policy TfPolicy#network_acl_common_policy}
	// Experimental.
	NetworkAclCommonPolicy *TfPolicy_NetworkAclCommonPolicyProperty `field:"optional" json:"networkAclCommonPolicy" yaml:"networkAclCommonPolicy"`
	// network_firewall_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#network_firewall_policy TfPolicy#network_firewall_policy}
	// Experimental.
	NetworkFirewallPolicy *TfPolicy_NetworkFirewallPolicyProperty `field:"optional" json:"networkFirewallPolicy" yaml:"networkFirewallPolicy"`
	// third_party_firewall_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#third_party_firewall_policy TfPolicy#third_party_firewall_policy}
	// Experimental.
	ThirdPartyFirewallPolicy *TfPolicy_ThirdPartyFirewallPolicyProperty `field:"optional" json:"thirdPartyFirewallPolicy" yaml:"thirdPartyFirewallPolicy"`
}

