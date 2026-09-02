package awseventbridgepipes


// Experimental.
type TfPipe_TargetParametersProperty struct {
	// batch_job_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#batch_job_parameters TfPipe#batch_job_parameters}
	// Experimental.
	BatchJobParameters *TfPipe_BatchJobParametersProperty `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// cloudwatch_logs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#cloudwatch_logs_parameters TfPipe#cloudwatch_logs_parameters}
	// Experimental.
	CloudwatchLogsParameters *TfPipe_CloudwatchLogsParametersProperty `field:"optional" json:"cloudwatchLogsParameters" yaml:"cloudwatchLogsParameters"`
	// ecs_task_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#ecs_task_parameters TfPipe#ecs_task_parameters}
	// Experimental.
	EcsTaskParameters *TfPipe_EcsTaskParametersProperty `field:"optional" json:"ecsTaskParameters" yaml:"ecsTaskParameters"`
	// eventbridge_event_bus_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#eventbridge_event_bus_parameters TfPipe#eventbridge_event_bus_parameters}
	// Experimental.
	EventbridgeEventBusParameters *TfPipe_EventbridgeEventBusParametersProperty `field:"optional" json:"eventbridgeEventBusParameters" yaml:"eventbridgeEventBusParameters"`
	// http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#http_parameters TfPipe#http_parameters}
	// Experimental.
	HttpParameters *TfPipe_TargetParametersHttpParametersProperty `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#input_template TfPipe#input_template}.
	// Experimental.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
	// kinesis_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#kinesis_stream_parameters TfPipe#kinesis_stream_parameters}
	// Experimental.
	KinesisStreamParameters *TfPipe_TargetParametersKinesisStreamParametersProperty `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// lambda_function_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#lambda_function_parameters TfPipe#lambda_function_parameters}
	// Experimental.
	LambdaFunctionParameters *TfPipe_LambdaFunctionParametersProperty `field:"optional" json:"lambdaFunctionParameters" yaml:"lambdaFunctionParameters"`
	// redshift_data_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#redshift_data_parameters TfPipe#redshift_data_parameters}
	// Experimental.
	RedshiftDataParameters *TfPipe_RedshiftDataParametersProperty `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// sagemaker_pipeline_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sagemaker_pipeline_parameters TfPipe#sagemaker_pipeline_parameters}
	// Experimental.
	SagemakerPipelineParameters *TfPipe_SagemakerPipelineParametersProperty `field:"optional" json:"sagemakerPipelineParameters" yaml:"sagemakerPipelineParameters"`
	// sqs_queue_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sqs_queue_parameters TfPipe#sqs_queue_parameters}
	// Experimental.
	SqsQueueParameters *TfPipe_TargetParametersSqsQueueParametersProperty `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
	// step_function_state_machine_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#step_function_state_machine_parameters TfPipe#step_function_state_machine_parameters}
	// Experimental.
	StepFunctionStateMachineParameters *TfPipe_StepFunctionStateMachineParametersProperty `field:"optional" json:"stepFunctionStateMachineParameters" yaml:"stepFunctionStateMachineParameters"`
}

