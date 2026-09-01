package awssagemakerai


// Experimental.
type AwsSagemakerTrainingJob_ProfilerRuleConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#rule_configuration_name AwsSagemakerTrainingJob#rule_configuration_name}.
	// Experimental.
	RuleConfigurationName *string `field:"required" json:"ruleConfigurationName" yaml:"ruleConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#rule_evaluator_image AwsSagemakerTrainingJob#rule_evaluator_image}.
	// Experimental.
	RuleEvaluatorImage *string `field:"required" json:"ruleEvaluatorImage" yaml:"ruleEvaluatorImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#instance_type AwsSagemakerTrainingJob#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#local_path AwsSagemakerTrainingJob#local_path}.
	// Experimental.
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#rule_parameters AwsSagemakerTrainingJob#rule_parameters}.
	// Experimental.
	RuleParameters *map[string]*string `field:"optional" json:"ruleParameters" yaml:"ruleParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#s3_output_path AwsSagemakerTrainingJob#s3_output_path}.
	// Experimental.
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#volume_size_in_gb AwsSagemakerTrainingJob#volume_size_in_gb}.
	// Experimental.
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

