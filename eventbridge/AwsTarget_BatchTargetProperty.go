package eventbridge


// Experimental.
type AwsTarget_BatchTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#job_definition AwsTarget#job_definition}.
	// Experimental.
	JobDefinition *string `field:"required" json:"jobDefinition" yaml:"jobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#job_name AwsTarget#job_name}.
	// Experimental.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#array_size AwsTarget#array_size}.
	// Experimental.
	ArraySize *float64 `field:"optional" json:"arraySize" yaml:"arraySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#job_attempts AwsTarget#job_attempts}.
	// Experimental.
	JobAttempts *float64 `field:"optional" json:"jobAttempts" yaml:"jobAttempts"`
}

