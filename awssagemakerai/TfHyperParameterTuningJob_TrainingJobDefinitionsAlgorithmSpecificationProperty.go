package awssagemakerai


// Experimental.
type TfHyperParameterTuningJob_TrainingJobDefinitionsAlgorithmSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#training_input_mode TfHyperParameterTuningJob#training_input_mode}.
	// Experimental.
	TrainingInputMode *string `field:"required" json:"trainingInputMode" yaml:"trainingInputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#algorithm_name TfHyperParameterTuningJob#algorithm_name}.
	// Experimental.
	AlgorithmName *string `field:"optional" json:"algorithmName" yaml:"algorithmName"`
	// metric_definitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#metric_definitions TfHyperParameterTuningJob#metric_definitions}
	// Experimental.
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#training_image TfHyperParameterTuningJob#training_image}.
	// Experimental.
	TrainingImage *string `field:"optional" json:"trainingImage" yaml:"trainingImage"`
}

