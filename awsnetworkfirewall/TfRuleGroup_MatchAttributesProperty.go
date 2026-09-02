package awsnetworkfirewall


// Experimental.
type TfRuleGroup_MatchAttributesProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#destination TfRuleGroup#destination}
	// Experimental.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// destination_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#destination_port TfRuleGroup#destination_port}
	// Experimental.
	DestinationPort interface{} `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#protocols TfRuleGroup#protocols}.
	// Experimental.
	Protocols *[]*float64 `field:"optional" json:"protocols" yaml:"protocols"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#source TfRuleGroup#source}
	// Experimental.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// source_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#source_port TfRuleGroup#source_port}
	// Experimental.
	SourcePort interface{} `field:"optional" json:"sourcePort" yaml:"sourcePort"`
	// tcp_flag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#tcp_flag TfRuleGroup#tcp_flag}
	// Experimental.
	TcpFlag interface{} `field:"optional" json:"tcpFlag" yaml:"tcpFlag"`
}

