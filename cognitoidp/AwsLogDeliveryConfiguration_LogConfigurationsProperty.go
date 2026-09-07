package cognitoidp


// Experimental.
type AwsLogDeliveryConfiguration_LogConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_log_delivery_configuration#event_source AwsLogDeliveryConfiguration#event_source}.
	// Experimental.
	EventSource *string `field:"required" json:"eventSource" yaml:"eventSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_log_delivery_configuration#log_level AwsLogDeliveryConfiguration#log_level}.
	// Experimental.
	LogLevel *string `field:"required" json:"logLevel" yaml:"logLevel"`
	// cloud_watch_logs_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_log_delivery_configuration#cloud_watch_logs_configuration AwsLogDeliveryConfiguration#cloud_watch_logs_configuration}
	// Experimental.
	CloudWatchLogsConfiguration interface{} `field:"optional" json:"cloudWatchLogsConfiguration" yaml:"cloudWatchLogsConfiguration"`
	// firehose_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_log_delivery_configuration#firehose_configuration AwsLogDeliveryConfiguration#firehose_configuration}
	// Experimental.
	FirehoseConfiguration interface{} `field:"optional" json:"firehoseConfiguration" yaml:"firehoseConfiguration"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_log_delivery_configuration#s3_configuration AwsLogDeliveryConfiguration#s3_configuration}
	// Experimental.
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

