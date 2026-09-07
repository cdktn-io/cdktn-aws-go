package autoscaling


// Experimental.
type AwsGroup_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#delete AwsGroup#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#update AwsGroup#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

