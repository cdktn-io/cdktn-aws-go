package awssesmailmanager


// Experimental.
type TfTrafficPolicy_TlsExpressionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#operator TfTrafficPolicy#operator}.
	// Experimental.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#value TfTrafficPolicy#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// evaluate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#evaluate TfTrafficPolicy#evaluate}
	// Experimental.
	Evaluate interface{} `field:"optional" json:"evaluate" yaml:"evaluate"`
}

