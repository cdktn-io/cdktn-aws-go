package sagemakerai


// Experimental.
type AwsTrainingJob_PlacementSpecificationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#instance_count AwsTrainingJob#instance_count}.
	// Experimental.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#ultra_server_id AwsTrainingJob#ultra_server_id}.
	// Experimental.
	UltraServerId *string `field:"optional" json:"ultraServerId" yaml:"ultraServerId"`
}

