package awskinesisanalytics


// Experimental.
type TfApplication_InputsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#name_prefix TfApplication#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#schema TfApplication#schema}
	// Experimental.
	Schema *TfApplication_InputsSchemaProperty `field:"required" json:"schema" yaml:"schema"`
	// kinesis_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_firehose TfApplication#kinesis_firehose}
	// Experimental.
	KinesisFirehose *TfApplication_InputsKinesisFirehoseProperty `field:"optional" json:"kinesisFirehose" yaml:"kinesisFirehose"`
	// kinesis_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#kinesis_stream TfApplication#kinesis_stream}
	// Experimental.
	KinesisStream *TfApplication_InputsKinesisStreamProperty `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// parallelism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#parallelism TfApplication#parallelism}
	// Experimental.
	Parallelism *TfApplication_ParallelismProperty `field:"optional" json:"parallelism" yaml:"parallelism"`
	// processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#processing_configuration TfApplication#processing_configuration}
	// Experimental.
	ProcessingConfiguration *TfApplication_ProcessingConfigurationProperty `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// starting_position_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_analytics_application#starting_position_configuration TfApplication#starting_position_configuration}
	// Experimental.
	StartingPositionConfiguration interface{} `field:"optional" json:"startingPositionConfiguration" yaml:"startingPositionConfiguration"`
}

