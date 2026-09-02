package awselb


// Experimental.
type TfAlbListener_TargetGroupProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#arn TfAlbListener#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#weight TfAlbListener#weight}.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

