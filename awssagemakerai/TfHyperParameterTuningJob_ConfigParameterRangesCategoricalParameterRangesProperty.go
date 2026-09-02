package awssagemakerai


// Experimental.
type TfHyperParameterTuningJob_ConfigParameterRangesCategoricalParameterRangesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#name TfHyperParameterTuningJob#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#values TfHyperParameterTuningJob#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

