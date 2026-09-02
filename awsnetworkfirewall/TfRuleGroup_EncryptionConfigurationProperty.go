package awsnetworkfirewall


// Experimental.
type TfRuleGroup_EncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#type TfRuleGroup#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#key_id TfRuleGroup#key_id}.
	// Experimental.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

