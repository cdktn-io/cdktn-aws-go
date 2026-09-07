package kinesisanalytics


// Experimental.
type AwsApplication_CloudwatchLoggingOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#log_stream_arn AwsApplication#log_stream_arn}.
	// Experimental.
	LogStreamArn *string `field:"required" json:"logStreamArn" yaml:"logStreamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#role_arn AwsApplication#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

