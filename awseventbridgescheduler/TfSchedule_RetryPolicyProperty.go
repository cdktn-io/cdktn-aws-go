package awseventbridgescheduler


// Experimental.
type TfSchedule_RetryPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#maximum_event_age_in_seconds TfSchedule#maximum_event_age_in_seconds}.
	// Experimental.
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#maximum_retry_attempts TfSchedule#maximum_retry_attempts}.
	// Experimental.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

