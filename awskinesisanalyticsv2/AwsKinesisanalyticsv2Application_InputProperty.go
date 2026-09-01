package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_InputProperty struct {
	// input_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_schema AwsKinesisanalyticsv2Application#input_schema}
	// Experimental.
	InputSchema *AwsKinesisanalyticsv2Application_InputSchemaProperty `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#name_prefix AwsKinesisanalyticsv2Application#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// input_parallelism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_parallelism AwsKinesisanalyticsv2Application#input_parallelism}
	// Experimental.
	InputParallelism *AwsKinesisanalyticsv2Application_InputParallelismProperty `field:"optional" json:"inputParallelism" yaml:"inputParallelism"`
	// input_processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_processing_configuration AwsKinesisanalyticsv2Application#input_processing_configuration}
	// Experimental.
	InputProcessingConfiguration *AwsKinesisanalyticsv2Application_InputProcessingConfigurationProperty `field:"optional" json:"inputProcessingConfiguration" yaml:"inputProcessingConfiguration"`
	// input_starting_position_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_starting_position_configuration AwsKinesisanalyticsv2Application#input_starting_position_configuration}
	// Experimental.
	InputStartingPositionConfiguration interface{} `field:"optional" json:"inputStartingPositionConfiguration" yaml:"inputStartingPositionConfiguration"`
	// kinesis_firehose_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_firehose_input AwsKinesisanalyticsv2Application#kinesis_firehose_input}
	// Experimental.
	KinesisFirehoseInput *AwsKinesisanalyticsv2Application_KinesisFirehoseInputProperty `field:"optional" json:"kinesisFirehoseInput" yaml:"kinesisFirehoseInput"`
	// kinesis_streams_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_streams_input AwsKinesisanalyticsv2Application#kinesis_streams_input}
	// Experimental.
	KinesisStreamsInput *AwsKinesisanalyticsv2Application_KinesisStreamsInputProperty `field:"optional" json:"kinesisStreamsInput" yaml:"kinesisStreamsInput"`
}

