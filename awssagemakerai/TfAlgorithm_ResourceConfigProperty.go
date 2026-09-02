package awssagemakerai


// Experimental.
type TfAlgorithm_ResourceConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#instance_count TfAlgorithm#instance_count}.
	// Experimental.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// instance_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#instance_groups TfAlgorithm#instance_groups}
	// Experimental.
	InstanceGroups interface{} `field:"optional" json:"instanceGroups" yaml:"instanceGroups"`
	// instance_placement_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#instance_placement_config TfAlgorithm#instance_placement_config}
	// Experimental.
	InstancePlacementConfig interface{} `field:"optional" json:"instancePlacementConfig" yaml:"instancePlacementConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#instance_type TfAlgorithm#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#keep_alive_period_in_seconds TfAlgorithm#keep_alive_period_in_seconds}.
	// Experimental.
	KeepAlivePeriodInSeconds *float64 `field:"optional" json:"keepAlivePeriodInSeconds" yaml:"keepAlivePeriodInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#training_plan_arn TfAlgorithm#training_plan_arn}.
	// Experimental.
	TrainingPlanArn *string `field:"optional" json:"trainingPlanArn" yaml:"trainingPlanArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#volume_kms_key_id TfAlgorithm#volume_kms_key_id}.
	// Experimental.
	VolumeKmsKeyId *string `field:"optional" json:"volumeKmsKeyId" yaml:"volumeKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#volume_size_in_gb TfAlgorithm#volume_size_in_gb}.
	// Experimental.
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

