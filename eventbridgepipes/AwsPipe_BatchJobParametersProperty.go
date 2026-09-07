package eventbridgepipes


// Experimental.
type AwsPipe_BatchJobParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#job_definition AwsPipe#job_definition}.
	// Experimental.
	JobDefinition *string `field:"required" json:"jobDefinition" yaml:"jobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#job_name AwsPipe#job_name}.
	// Experimental.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// array_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#array_properties AwsPipe#array_properties}
	// Experimental.
	ArrayProperties *AwsPipe_ArrayPropertiesProperty `field:"optional" json:"arrayProperties" yaml:"arrayProperties"`
	// container_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#container_overrides AwsPipe#container_overrides}
	// Experimental.
	ContainerOverrides *AwsPipe_ContainerOverridesProperty `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// depends_on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#depends_on AwsPipe#depends_on}
	// Experimental.
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#parameters AwsPipe#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// retry_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#retry_strategy AwsPipe#retry_strategy}
	// Experimental.
	RetryStrategy *AwsPipe_RetryStrategyProperty `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
}

