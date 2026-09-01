package awskinesisanalytics


// Experimental.
type AwsKinesisAnalyticsApplication_OutputsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#name AwsKinesisAnalyticsApplication#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#schema AwsKinesisAnalyticsApplication#schema}
	// Experimental.
	Schema *AwsKinesisAnalyticsApplication_OutputsSchemaProperty `field:"required" json:"schema" yaml:"schema"`
	// kinesis_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_firehose AwsKinesisAnalyticsApplication#kinesis_firehose}
	// Experimental.
	KinesisFirehose *AwsKinesisAnalyticsApplication_OutputsKinesisFirehoseProperty `field:"optional" json:"kinesisFirehose" yaml:"kinesisFirehose"`
	// kinesis_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_stream AwsKinesisAnalyticsApplication#kinesis_stream}
	// Experimental.
	KinesisStream *AwsKinesisAnalyticsApplication_OutputsKinesisStreamProperty `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#lambda AwsKinesisAnalyticsApplication#lambda}
	// Experimental.
	Lambda *AwsKinesisAnalyticsApplication_OutputsLambdaProperty `field:"optional" json:"lambda" yaml:"lambda"`
}

