package awscognitoidp


// Experimental.
type AwsCognitoLogDeliveryConfiguration_LogConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_log_delivery_configuration#event_source AwsCognitoLogDeliveryConfiguration#event_source}.
	// Experimental.
	EventSource *string `field:"required" json:"eventSource" yaml:"eventSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_log_delivery_configuration#log_level AwsCognitoLogDeliveryConfiguration#log_level}.
	// Experimental.
	LogLevel *string `field:"required" json:"logLevel" yaml:"logLevel"`
	// cloud_watch_logs_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_log_delivery_configuration#cloud_watch_logs_configuration AwsCognitoLogDeliveryConfiguration#cloud_watch_logs_configuration}
	// Experimental.
	CloudWatchLogsConfiguration interface{} `field:"optional" json:"cloudWatchLogsConfiguration" yaml:"cloudWatchLogsConfiguration"`
	// firehose_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_log_delivery_configuration#firehose_configuration AwsCognitoLogDeliveryConfiguration#firehose_configuration}
	// Experimental.
	FirehoseConfiguration interface{} `field:"optional" json:"firehoseConfiguration" yaml:"firehoseConfiguration"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_log_delivery_configuration#s3_configuration AwsCognitoLogDeliveryConfiguration#s3_configuration}
	// Experimental.
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

