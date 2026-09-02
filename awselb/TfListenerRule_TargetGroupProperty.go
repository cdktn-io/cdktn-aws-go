package awselb


// Experimental.
type TfListenerRule_TargetGroupProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#arn TfListenerRule#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#weight TfListenerRule#weight}.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

