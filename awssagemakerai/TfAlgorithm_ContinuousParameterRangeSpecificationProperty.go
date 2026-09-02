package awssagemakerai


// Experimental.
type TfAlgorithm_ContinuousParameterRangeSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#max_value TfAlgorithm#max_value}.
	// Experimental.
	MaxValue *string `field:"required" json:"maxValue" yaml:"maxValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#min_value TfAlgorithm#min_value}.
	// Experimental.
	MinValue *string `field:"required" json:"minValue" yaml:"minValue"`
}

