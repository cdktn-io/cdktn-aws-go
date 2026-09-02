package awsec2imagebuilder


// Experimental.
type TfImagePipeline_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#name TfImagePipeline#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#value TfImagePipeline#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

