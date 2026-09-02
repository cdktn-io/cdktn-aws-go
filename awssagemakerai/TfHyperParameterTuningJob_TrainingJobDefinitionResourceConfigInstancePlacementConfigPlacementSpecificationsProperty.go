package awssagemakerai


// Experimental.
type TfHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPlacementSpecificationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#instance_count TfHyperParameterTuningJob#instance_count}.
	// Experimental.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#ultra_server_id TfHyperParameterTuningJob#ultra_server_id}.
	// Experimental.
	UltraServerId *string `field:"optional" json:"ultraServerId" yaml:"ultraServerId"`
}

