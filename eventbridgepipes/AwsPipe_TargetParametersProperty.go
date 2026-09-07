package eventbridgepipes


// Experimental.
type AwsPipe_TargetParametersProperty struct {
	// batch_job_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#batch_job_parameters AwsPipe#batch_job_parameters}
	// Experimental.
	BatchJobParameters *AwsPipe_BatchJobParametersProperty `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// cloudwatch_logs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#cloudwatch_logs_parameters AwsPipe#cloudwatch_logs_parameters}
	// Experimental.
	CloudwatchLogsParameters *AwsPipe_CloudwatchLogsParametersProperty `field:"optional" json:"cloudwatchLogsParameters" yaml:"cloudwatchLogsParameters"`
	// ecs_task_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#ecs_task_parameters AwsPipe#ecs_task_parameters}
	// Experimental.
	EcsTaskParameters *AwsPipe_EcsTaskParametersProperty `field:"optional" json:"ecsTaskParameters" yaml:"ecsTaskParameters"`
	// eventbridge_event_bus_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#eventbridge_event_bus_parameters AwsPipe#eventbridge_event_bus_parameters}
	// Experimental.
	EventbridgeEventBusParameters *AwsPipe_EventbridgeEventBusParametersProperty `field:"optional" json:"eventbridgeEventBusParameters" yaml:"eventbridgeEventBusParameters"`
	// http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#http_parameters AwsPipe#http_parameters}
	// Experimental.
	HttpParameters *AwsPipe_TargetParametersHttpParametersProperty `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#input_template AwsPipe#input_template}.
	// Experimental.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
	// kinesis_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#kinesis_stream_parameters AwsPipe#kinesis_stream_parameters}
	// Experimental.
	KinesisStreamParameters *AwsPipe_TargetParametersKinesisStreamParametersProperty `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// lambda_function_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#lambda_function_parameters AwsPipe#lambda_function_parameters}
	// Experimental.
	LambdaFunctionParameters *AwsPipe_LambdaFunctionParametersProperty `field:"optional" json:"lambdaFunctionParameters" yaml:"lambdaFunctionParameters"`
	// redshift_data_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#redshift_data_parameters AwsPipe#redshift_data_parameters}
	// Experimental.
	RedshiftDataParameters *AwsPipe_RedshiftDataParametersProperty `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// sagemaker_pipeline_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sagemaker_pipeline_parameters AwsPipe#sagemaker_pipeline_parameters}
	// Experimental.
	SagemakerPipelineParameters *AwsPipe_SagemakerPipelineParametersProperty `field:"optional" json:"sagemakerPipelineParameters" yaml:"sagemakerPipelineParameters"`
	// sqs_queue_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sqs_queue_parameters AwsPipe#sqs_queue_parameters}
	// Experimental.
	SqsQueueParameters *AwsPipe_TargetParametersSqsQueueParametersProperty `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
	// step_function_state_machine_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#step_function_state_machine_parameters AwsPipe#step_function_state_machine_parameters}
	// Experimental.
	StepFunctionStateMachineParameters *AwsPipe_StepFunctionStateMachineParametersProperty `field:"optional" json:"stepFunctionStateMachineParameters" yaml:"stepFunctionStateMachineParameters"`
}

