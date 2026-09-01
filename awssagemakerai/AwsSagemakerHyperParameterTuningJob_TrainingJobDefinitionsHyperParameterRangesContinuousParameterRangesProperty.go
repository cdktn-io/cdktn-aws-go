package awssagemakerai


// Experimental.
type AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesContinuousParameterRangesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_value AwsSagemakerHyperParameterTuningJob#max_value}.
	// Experimental.
	MaxValue *string `field:"required" json:"maxValue" yaml:"maxValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#min_value AwsSagemakerHyperParameterTuningJob#min_value}.
	// Experimental.
	MinValue *string `field:"required" json:"minValue" yaml:"minValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#name AwsSagemakerHyperParameterTuningJob#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#scaling_type AwsSagemakerHyperParameterTuningJob#scaling_type}.
	// Experimental.
	ScalingType *string `field:"optional" json:"scalingType" yaml:"scalingType"`
}

