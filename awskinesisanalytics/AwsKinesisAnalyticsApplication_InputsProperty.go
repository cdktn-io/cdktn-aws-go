package awskinesisanalytics


// Experimental.
type AwsKinesisAnalyticsApplication_InputsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#name_prefix AwsKinesisAnalyticsApplication#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#schema AwsKinesisAnalyticsApplication#schema}
	// Experimental.
	Schema *AwsKinesisAnalyticsApplication_InputsSchemaProperty `field:"required" json:"schema" yaml:"schema"`
	// kinesis_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_firehose AwsKinesisAnalyticsApplication#kinesis_firehose}
	// Experimental.
	KinesisFirehose *AwsKinesisAnalyticsApplication_InputsKinesisFirehoseProperty `field:"optional" json:"kinesisFirehose" yaml:"kinesisFirehose"`
	// kinesis_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_stream AwsKinesisAnalyticsApplication#kinesis_stream}
	// Experimental.
	KinesisStream *AwsKinesisAnalyticsApplication_InputsKinesisStreamProperty `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// parallelism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#parallelism AwsKinesisAnalyticsApplication#parallelism}
	// Experimental.
	Parallelism *AwsKinesisAnalyticsApplication_ParallelismProperty `field:"optional" json:"parallelism" yaml:"parallelism"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#processing_configuration AwsKinesisAnalyticsApplication#processing_configuration}
	// Experimental.
	ProcessingConfiguration *AwsKinesisAnalyticsApplication_ProcessingConfigurationProperty `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// starting_position_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#starting_position_configuration AwsKinesisAnalyticsApplication#starting_position_configuration}
	// Experimental.
	StartingPositionConfiguration interface{} `field:"optional" json:"startingPositionConfiguration" yaml:"startingPositionConfiguration"`
}

