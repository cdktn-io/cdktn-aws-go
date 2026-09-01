package awscloudwatchsynthetics


// Experimental.
type AwsSyntheticsCanary_ScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#expression AwsSyntheticsCanary#expression}.
	// Experimental.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#duration_in_seconds AwsSyntheticsCanary#duration_in_seconds}.
	// Experimental.
	DurationInSeconds *float64 `field:"optional" json:"durationInSeconds" yaml:"durationInSeconds"`
	// retry_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#retry_config AwsSyntheticsCanary#retry_config}
	// Experimental.
	RetryConfig *AwsSyntheticsCanary_RetryConfigProperty `field:"optional" json:"retryConfig" yaml:"retryConfig"`
}

