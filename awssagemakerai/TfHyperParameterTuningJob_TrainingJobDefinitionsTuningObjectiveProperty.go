package awssagemakerai


// Experimental.
type TfHyperParameterTuningJob_TrainingJobDefinitionsTuningObjectiveProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#metric_name TfHyperParameterTuningJob#metric_name}.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#type TfHyperParameterTuningJob#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

