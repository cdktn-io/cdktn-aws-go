package awseventbridge


// Experimental.
type TfTarget_RetryPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#maximum_event_age_in_seconds TfTarget#maximum_event_age_in_seconds}.
	// Experimental.
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#maximum_retry_attempts TfTarget#maximum_retry_attempts}.
	// Experimental.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

