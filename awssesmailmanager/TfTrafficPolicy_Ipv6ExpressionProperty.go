package awssesmailmanager


// Experimental.
type TfTrafficPolicy_Ipv6ExpressionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#operator TfTrafficPolicy#operator}.
	// Experimental.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#values TfTrafficPolicy#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// evaluate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#evaluate TfTrafficPolicy#evaluate}
	// Experimental.
	Evaluate interface{} `field:"optional" json:"evaluate" yaml:"evaluate"`
}

