package eventbridgescheduler


// Experimental.
type AwsSchedule_TargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#arn AwsSchedule#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#role_arn AwsSchedule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#dead_letter_config AwsSchedule#dead_letter_config}
	// Experimental.
	DeadLetterConfig *AwsSchedule_DeadLetterConfigProperty `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// ecs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#ecs_parameters AwsSchedule#ecs_parameters}
	// Experimental.
	EcsParameters *AwsSchedule_EcsParametersProperty `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// eventbridge_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#eventbridge_parameters AwsSchedule#eventbridge_parameters}
	// Experimental.
	EventbridgeParameters *AwsSchedule_EventbridgeParametersProperty `field:"optional" json:"eventbridgeParameters" yaml:"eventbridgeParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#input AwsSchedule#input}.
	// Experimental.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// kinesis_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#kinesis_parameters AwsSchedule#kinesis_parameters}
	// Experimental.
	KinesisParameters *AwsSchedule_KinesisParametersProperty `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#retry_policy AwsSchedule#retry_policy}
	// Experimental.
	RetryPolicy *AwsSchedule_RetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// sagemaker_pipeline_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#sagemaker_pipeline_parameters AwsSchedule#sagemaker_pipeline_parameters}
	// Experimental.
	SagemakerPipelineParameters *AwsSchedule_SagemakerPipelineParametersProperty `field:"optional" json:"sagemakerPipelineParameters" yaml:"sagemakerPipelineParameters"`
	// sqs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#sqs_parameters AwsSchedule#sqs_parameters}
	// Experimental.
	SqsParameters *AwsSchedule_SqsParametersProperty `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
}

