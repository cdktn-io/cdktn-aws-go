package awssagemakerai


// Experimental.
type TfAlgorithm_RangeProperty struct {
	// categorical_parameter_range_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#categorical_parameter_range_specification TfAlgorithm#categorical_parameter_range_specification}
	// Experimental.
	CategoricalParameterRangeSpecification interface{} `field:"optional" json:"categoricalParameterRangeSpecification" yaml:"categoricalParameterRangeSpecification"`
	// continuous_parameter_range_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#continuous_parameter_range_specification TfAlgorithm#continuous_parameter_range_specification}
	// Experimental.
	ContinuousParameterRangeSpecification interface{} `field:"optional" json:"continuousParameterRangeSpecification" yaml:"continuousParameterRangeSpecification"`
	// integer_parameter_range_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#integer_parameter_range_specification TfAlgorithm#integer_parameter_range_specification}
	// Experimental.
	IntegerParameterRangeSpecification interface{} `field:"optional" json:"integerParameterRangeSpecification" yaml:"integerParameterRangeSpecification"`
}

