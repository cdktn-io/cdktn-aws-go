package awseventbridgepipes


// Experimental.
type AwsPipesPipe_BatchJobParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#job_definition AwsPipesPipe#job_definition}.
	// Experimental.
	JobDefinition *string `field:"required" json:"jobDefinition" yaml:"jobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#job_name AwsPipesPipe#job_name}.
	// Experimental.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// array_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#array_properties AwsPipesPipe#array_properties}
	// Experimental.
	ArrayProperties *AwsPipesPipe_ArrayPropertiesProperty `field:"optional" json:"arrayProperties" yaml:"arrayProperties"`
	// container_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#container_overrides AwsPipesPipe#container_overrides}
	// Experimental.
	ContainerOverrides *AwsPipesPipe_ContainerOverridesProperty `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// depends_on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#depends_on AwsPipesPipe#depends_on}
	// Experimental.
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#parameters AwsPipesPipe#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// retry_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#retry_strategy AwsPipesPipe#retry_strategy}
	// Experimental.
	RetryStrategy *AwsPipesPipe_RetryStrategyProperty `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
}

