package awscloudformation


// Experimental.
type TfStackSetInstance_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#create TfStackSetInstance#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#delete TfStackSetInstance#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#update TfStackSetInstance#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

