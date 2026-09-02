package awscloudwatchsynthetics


// Experimental.
type TfCanary_ScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#expression TfCanary#expression}.
	// Experimental.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#duration_in_seconds TfCanary#duration_in_seconds}.
	// Experimental.
	DurationInSeconds *float64 `field:"optional" json:"durationInSeconds" yaml:"durationInSeconds"`
	// retry_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#retry_config TfCanary#retry_config}
	// Experimental.
	RetryConfig *TfCanary_RetryConfigProperty `field:"optional" json:"retryConfig" yaml:"retryConfig"`
}

