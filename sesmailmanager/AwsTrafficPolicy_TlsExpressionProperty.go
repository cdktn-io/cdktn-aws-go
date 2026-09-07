package sesmailmanager


// Experimental.
type AwsTrafficPolicy_TlsExpressionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#operator AwsTrafficPolicy#operator}.
	// Experimental.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#value AwsTrafficPolicy#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// evaluate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#evaluate AwsTrafficPolicy#evaluate}
	// Experimental.
	Evaluate interface{} `field:"optional" json:"evaluate" yaml:"evaluate"`
}

