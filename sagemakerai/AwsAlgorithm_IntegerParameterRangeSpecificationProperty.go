package sagemakerai


// Experimental.
type AwsAlgorithm_IntegerParameterRangeSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#max_value AwsAlgorithm#max_value}.
	// Experimental.
	MaxValue *string `field:"required" json:"maxValue" yaml:"maxValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#min_value AwsAlgorithm#min_value}.
	// Experimental.
	MinValue *string `field:"required" json:"minValue" yaml:"minValue"`
}

