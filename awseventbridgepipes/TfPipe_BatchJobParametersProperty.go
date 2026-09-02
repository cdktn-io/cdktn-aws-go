package awseventbridgepipes


// Experimental.
type TfPipe_BatchJobParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#job_definition TfPipe#job_definition}.
	// Experimental.
	JobDefinition *string `field:"required" json:"jobDefinition" yaml:"jobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#job_name TfPipe#job_name}.
	// Experimental.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// array_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#array_properties TfPipe#array_properties}
	// Experimental.
	ArrayProperties *TfPipe_ArrayPropertiesProperty `field:"optional" json:"arrayProperties" yaml:"arrayProperties"`
	// container_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#container_overrides TfPipe#container_overrides}
	// Experimental.
	ContainerOverrides *TfPipe_ContainerOverridesProperty `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// depends_on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#depends_on TfPipe#depends_on}
	// Experimental.
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#parameters TfPipe#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// retry_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#retry_strategy TfPipe#retry_strategy}
	// Experimental.
	RetryStrategy *TfPipe_RetryStrategyProperty `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
}

