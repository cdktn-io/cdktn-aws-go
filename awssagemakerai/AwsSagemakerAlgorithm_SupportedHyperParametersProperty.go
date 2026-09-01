package awssagemakerai


// Experimental.
type AwsSagemakerAlgorithm_SupportedHyperParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#name AwsSagemakerAlgorithm#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#type AwsSagemakerAlgorithm#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#default_value AwsSagemakerAlgorithm#default_value}.
	// Experimental.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#description AwsSagemakerAlgorithm#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#is_required AwsSagemakerAlgorithm#is_required}.
	// Experimental.
	IsRequired interface{} `field:"optional" json:"isRequired" yaml:"isRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#is_tunable AwsSagemakerAlgorithm#is_tunable}.
	// Experimental.
	IsTunable interface{} `field:"optional" json:"isTunable" yaml:"isTunable"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#range AwsSagemakerAlgorithm#range}
	// Experimental.
	Range interface{} `field:"optional" json:"range" yaml:"range"`
}

