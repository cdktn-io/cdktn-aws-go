package sagemakerai


// Experimental.
type AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPlacementSpecificationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#instance_count AwsHyperParameterTuningJob#instance_count}.
	// Experimental.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#ultra_server_id AwsHyperParameterTuningJob#ultra_server_id}.
	// Experimental.
	UltraServerId *string `field:"optional" json:"ultraServerId" yaml:"ultraServerId"`
}

