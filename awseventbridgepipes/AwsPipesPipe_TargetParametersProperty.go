package awseventbridgepipes


// Experimental.
type AwsPipesPipe_TargetParametersProperty struct {
	// batch_job_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#batch_job_parameters AwsPipesPipe#batch_job_parameters}
	// Experimental.
	BatchJobParameters *AwsPipesPipe_BatchJobParametersProperty `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// cloudwatch_logs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#cloudwatch_logs_parameters AwsPipesPipe#cloudwatch_logs_parameters}
	// Experimental.
	CloudwatchLogsParameters *AwsPipesPipe_CloudwatchLogsParametersProperty `field:"optional" json:"cloudwatchLogsParameters" yaml:"cloudwatchLogsParameters"`
	// ecs_task_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#ecs_task_parameters AwsPipesPipe#ecs_task_parameters}
	// Experimental.
	EcsTaskParameters *AwsPipesPipe_EcsTaskParametersProperty `field:"optional" json:"ecsTaskParameters" yaml:"ecsTaskParameters"`
	// eventbridge_event_bus_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#eventbridge_event_bus_parameters AwsPipesPipe#eventbridge_event_bus_parameters}
	// Experimental.
	EventbridgeEventBusParameters *AwsPipesPipe_EventbridgeEventBusParametersProperty `field:"optional" json:"eventbridgeEventBusParameters" yaml:"eventbridgeEventBusParameters"`
	// http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#http_parameters AwsPipesPipe#http_parameters}
	// Experimental.
	HttpParameters *AwsPipesPipe_TargetParametersHttpParametersProperty `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#input_template AwsPipesPipe#input_template}.
	// Experimental.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
	// kinesis_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#kinesis_stream_parameters AwsPipesPipe#kinesis_stream_parameters}
	// Experimental.
	KinesisStreamParameters *AwsPipesPipe_TargetParametersKinesisStreamParametersProperty `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// lambda_function_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#lambda_function_parameters AwsPipesPipe#lambda_function_parameters}
	// Experimental.
	LambdaFunctionParameters *AwsPipesPipe_LambdaFunctionParametersProperty `field:"optional" json:"lambdaFunctionParameters" yaml:"lambdaFunctionParameters"`
	// redshift_data_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#redshift_data_parameters AwsPipesPipe#redshift_data_parameters}
	// Experimental.
	RedshiftDataParameters *AwsPipesPipe_RedshiftDataParametersProperty `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// sagemaker_pipeline_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sagemaker_pipeline_parameters AwsPipesPipe#sagemaker_pipeline_parameters}
	// Experimental.
	SagemakerPipelineParameters *AwsPipesPipe_SagemakerPipelineParametersProperty `field:"optional" json:"sagemakerPipelineParameters" yaml:"sagemakerPipelineParameters"`
	// sqs_queue_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sqs_queue_parameters AwsPipesPipe#sqs_queue_parameters}
	// Experimental.
	SqsQueueParameters *AwsPipesPipe_TargetParametersSqsQueueParametersProperty `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
	// step_function_state_machine_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#step_function_state_machine_parameters AwsPipesPipe#step_function_state_machine_parameters}
	// Experimental.
	StepFunctionStateMachineParameters *AwsPipesPipe_StepFunctionStateMachineParametersProperty `field:"optional" json:"stepFunctionStateMachineParameters" yaml:"stepFunctionStateMachineParameters"`
}

