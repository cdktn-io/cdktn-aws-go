package awseventbridgescheduler


// Experimental.
type AwsSchedulerSchedule_TargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#arn AwsSchedulerSchedule#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#role_arn AwsSchedulerSchedule#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#dead_letter_config AwsSchedulerSchedule#dead_letter_config}
	// Experimental.
	DeadLetterConfig *AwsSchedulerSchedule_DeadLetterConfigProperty `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// ecs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#ecs_parameters AwsSchedulerSchedule#ecs_parameters}
	// Experimental.
	EcsParameters *AwsSchedulerSchedule_EcsParametersProperty `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// eventbridge_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#eventbridge_parameters AwsSchedulerSchedule#eventbridge_parameters}
	// Experimental.
	EventbridgeParameters *AwsSchedulerSchedule_EventbridgeParametersProperty `field:"optional" json:"eventbridgeParameters" yaml:"eventbridgeParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#input AwsSchedulerSchedule#input}.
	// Experimental.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// kinesis_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#kinesis_parameters AwsSchedulerSchedule#kinesis_parameters}
	// Experimental.
	KinesisParameters *AwsSchedulerSchedule_KinesisParametersProperty `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#retry_policy AwsSchedulerSchedule#retry_policy}
	// Experimental.
	RetryPolicy *AwsSchedulerSchedule_RetryPolicyProperty `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// sagemaker_pipeline_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#sagemaker_pipeline_parameters AwsSchedulerSchedule#sagemaker_pipeline_parameters}
	// Experimental.
	SagemakerPipelineParameters *AwsSchedulerSchedule_SagemakerPipelineParametersProperty `field:"optional" json:"sagemakerPipelineParameters" yaml:"sagemakerPipelineParameters"`
	// sqs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#sqs_parameters AwsSchedulerSchedule#sqs_parameters}
	// Experimental.
	SqsParameters *AwsSchedulerSchedule_SqsParametersProperty `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
}

