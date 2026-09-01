package awssagemakerai


// Experimental.
type AwsSagemakerAlgorithm_TrainingChannelsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#name AwsSagemakerAlgorithm#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#supported_content_types AwsSagemakerAlgorithm#supported_content_types}.
	// Experimental.
	SupportedContentTypes *[]*string `field:"required" json:"supportedContentTypes" yaml:"supportedContentTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#supported_input_modes AwsSagemakerAlgorithm#supported_input_modes}.
	// Experimental.
	SupportedInputModes *[]*string `field:"required" json:"supportedInputModes" yaml:"supportedInputModes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#description AwsSagemakerAlgorithm#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#is_required AwsSagemakerAlgorithm#is_required}.
	// Experimental.
	IsRequired interface{} `field:"optional" json:"isRequired" yaml:"isRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#supported_compression_types AwsSagemakerAlgorithm#supported_compression_types}.
	// Experimental.
	SupportedCompressionTypes *[]*string `field:"optional" json:"supportedCompressionTypes" yaml:"supportedCompressionTypes"`
}

