package eventbridge


// Experimental.
type AwsTarget_SagemakerPipelineTargetProperty struct {
	// pipeline_parameter_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#pipeline_parameter_list AwsTarget#pipeline_parameter_list}
	// Experimental.
	PipelineParameterList interface{} `field:"optional" json:"pipelineParameterList" yaml:"pipelineParameterList"`
}

