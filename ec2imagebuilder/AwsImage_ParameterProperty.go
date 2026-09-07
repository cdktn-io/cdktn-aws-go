package ec2imagebuilder


// Experimental.
type AwsImage_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#name AwsImage#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#value AwsImage#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

