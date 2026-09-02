package awseventbridge


// Experimental.
type TfTarget_BatchTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#job_definition TfTarget#job_definition}.
	// Experimental.
	JobDefinition *string `field:"required" json:"jobDefinition" yaml:"jobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#job_name TfTarget#job_name}.
	// Experimental.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#array_size TfTarget#array_size}.
	// Experimental.
	ArraySize *float64 `field:"optional" json:"arraySize" yaml:"arraySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#job_attempts TfTarget#job_attempts}.
	// Experimental.
	JobAttempts *float64 `field:"optional" json:"jobAttempts" yaml:"jobAttempts"`
}

