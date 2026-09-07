package networkfirewall


// Experimental.
type AwsRuleGroup_TcpFlagProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#flags AwsRuleGroup#flags}.
	// Experimental.
	Flags *[]*string `field:"required" json:"flags" yaml:"flags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#masks AwsRuleGroup#masks}.
	// Experimental.
	Masks *[]*string `field:"optional" json:"masks" yaml:"masks"`
}

