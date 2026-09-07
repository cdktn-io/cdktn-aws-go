package sagemakerai


// Experimental.
type AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesAutoParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#name AwsHyperParameterTuningJob#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#value_hint AwsHyperParameterTuningJob#value_hint}.
	// Experimental.
	ValueHint *string `field:"required" json:"valueHint" yaml:"valueHint"`
}

