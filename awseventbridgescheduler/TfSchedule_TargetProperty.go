package awseventbridgescheduler


// Experimental.
type TfSchedule_TargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#arn TfSchedule#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#role_arn TfSchedule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#dead_letter_config TfSchedule#dead_letter_config}
	// Experimental.
	DeadLetterConfig *TfSchedule_DeadLetterConfigProperty `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// ecs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#ecs_parameters TfSchedule#ecs_parameters}
	// Experimental.
	EcsParameters *TfSchedule_EcsParametersProperty `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// eventbridge_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#eventbridge_parameters TfSchedule#eventbridge_parameters}
	// Experimental.
	EventbridgeParameters *TfSchedule_EventbridgeParametersProperty `field:"optional" json:"eventbridgeParameters" yaml:"eventbridgeParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#input TfSchedule#input}.
	// Experimental.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// kinesis_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#kinesis_parameters TfSchedule#kinesis_parameters}
	// Experimental.
	KinesisParameters *TfSchedule_KinesisParametersProperty `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#retry_policy TfSchedule#retry_policy}
	// Experimental.
	RetryPolicy *TfSchedule_RetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// sagemaker_pipeline_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#sagemaker_pipeline_parameters TfSchedule#sagemaker_pipeline_parameters}
	// Experimental.
	SagemakerPipelineParameters *TfSchedule_SagemakerPipelineParametersProperty `field:"optional" json:"sagemakerPipelineParameters" yaml:"sagemakerPipelineParameters"`
	// sqs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#sqs_parameters TfSchedule#sqs_parameters}
	// Experimental.
	SqsParameters *TfSchedule_SqsParametersProperty `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
}

