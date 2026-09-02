package awssesmailmanager


// Experimental.
type TfTrafficPolicy_ConditionProperty struct {
	// boolean_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#boolean_expression TfTrafficPolicy#boolean_expression}
	// Experimental.
	BooleanExpression interface{} `field:"optional" json:"booleanExpression" yaml:"booleanExpression"`
	// ip_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#ip_expression TfTrafficPolicy#ip_expression}
	// Experimental.
	IpExpression interface{} `field:"optional" json:"ipExpression" yaml:"ipExpression"`
	// ipv6_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#ipv6_expression TfTrafficPolicy#ipv6_expression}
	// Experimental.
	Ipv6Expression interface{} `field:"optional" json:"ipv6Expression" yaml:"ipv6Expression"`
	// string_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#string_expression TfTrafficPolicy#string_expression}
	// Experimental.
	StringExpression interface{} `field:"optional" json:"stringExpression" yaml:"stringExpression"`
	// tls_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#tls_expression TfTrafficPolicy#tls_expression}
	// Experimental.
	TlsExpression interface{} `field:"optional" json:"tlsExpression" yaml:"tlsExpression"`
}

