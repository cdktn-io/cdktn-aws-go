package batch


// Experimental.
type AwsJobDefinition_EvaluateOnExitProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#action AwsJobDefinition#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#on_exit_code AwsJobDefinition#on_exit_code}.
	// Experimental.
	OnExitCode *string `field:"optional" json:"onExitCode" yaml:"onExitCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#on_reason AwsJobDefinition#on_reason}.
	// Experimental.
	OnReason *string `field:"optional" json:"onReason" yaml:"onReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#on_status_reason AwsJobDefinition#on_status_reason}.
	// Experimental.
	OnStatusReason *string `field:"optional" json:"onStatusReason" yaml:"onStatusReason"`
}

