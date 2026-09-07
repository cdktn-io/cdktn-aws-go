package kinesisanalytics


// Experimental.
type AwsApplication_InputsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#name_prefix AwsApplication#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#schema AwsApplication#schema}
	// Experimental.
	Schema *AwsApplication_InputsSchemaProperty `field:"required" json:"schema" yaml:"schema"`
	// kinesis_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_firehose AwsApplication#kinesis_firehose}
	// Experimental.
	KinesisFirehose *AwsApplication_InputsKinesisFirehoseProperty `field:"optional" json:"kinesisFirehose" yaml:"kinesisFirehose"`
	// kinesis_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_stream AwsApplication#kinesis_stream}
	// Experimental.
	KinesisStream *AwsApplication_InputsKinesisStreamProperty `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// parallelism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#parallelism AwsApplication#parallelism}
	// Experimental.
	Parallelism *AwsApplication_ParallelismProperty `field:"optional" json:"parallelism" yaml:"parallelism"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#processing_configuration AwsApplication#processing_configuration}
	// Experimental.
	ProcessingConfiguration *AwsApplication_ProcessingConfigurationProperty `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// starting_position_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#starting_position_configuration AwsApplication#starting_position_configuration}
	// Experimental.
	StartingPositionConfiguration interface{} `field:"optional" json:"startingPositionConfiguration" yaml:"startingPositionConfiguration"`
}

