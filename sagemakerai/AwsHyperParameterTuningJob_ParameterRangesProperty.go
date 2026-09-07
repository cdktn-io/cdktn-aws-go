package sagemakerai


// Experimental.
type AwsHyperParameterTuningJob_ParameterRangesProperty struct {
	// auto_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#auto_parameters AwsHyperParameterTuningJob#auto_parameters}
	// Experimental.
	AutoParameters interface{} `field:"optional" json:"autoParameters" yaml:"autoParameters"`
	// categorical_parameter_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#categorical_parameter_ranges AwsHyperParameterTuningJob#categorical_parameter_ranges}
	// Experimental.
	CategoricalParameterRanges interface{} `field:"optional" json:"categoricalParameterRanges" yaml:"categoricalParameterRanges"`
	// continuous_parameter_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#continuous_parameter_ranges AwsHyperParameterTuningJob#continuous_parameter_ranges}
	// Experimental.
	ContinuousParameterRanges interface{} `field:"optional" json:"continuousParameterRanges" yaml:"continuousParameterRanges"`
	// integer_parameter_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#integer_parameter_ranges AwsHyperParameterTuningJob#integer_parameter_ranges}
	// Experimental.
	IntegerParameterRanges interface{} `field:"optional" json:"integerParameterRanges" yaml:"integerParameterRanges"`
}

