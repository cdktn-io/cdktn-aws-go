package awsbatch


// Experimental.
type TfJobDefinition_TimeoutProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#attempt_duration_seconds TfJobDefinition#attempt_duration_seconds}.
	// Experimental.
	AttemptDurationSeconds *float64 `field:"optional" json:"attemptDurationSeconds" yaml:"attemptDurationSeconds"`
}

