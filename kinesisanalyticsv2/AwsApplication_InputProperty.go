package kinesisanalyticsv2


// Experimental.
type AwsApplication_InputProperty struct {
	// input_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_schema AwsApplication#input_schema}
	// Experimental.
	InputSchema *AwsApplication_InputSchemaProperty `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#name_prefix AwsApplication#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// input_parallelism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_parallelism AwsApplication#input_parallelism}
	// Experimental.
	InputParallelism *AwsApplication_InputParallelismProperty `field:"optional" json:"inputParallelism" yaml:"inputParallelism"`
	// input_processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_processing_configuration AwsApplication#input_processing_configuration}
	// Experimental.
	InputProcessingConfiguration *AwsApplication_InputProcessingConfigurationProperty `field:"optional" json:"inputProcessingConfiguration" yaml:"inputProcessingConfiguration"`
	// input_starting_position_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_starting_position_configuration AwsApplication#input_starting_position_configuration}
	// Experimental.
	InputStartingPositionConfiguration interface{} `field:"optional" json:"inputStartingPositionConfiguration" yaml:"inputStartingPositionConfiguration"`
	// kinesis_firehose_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_firehose_input AwsApplication#kinesis_firehose_input}
	// Experimental.
	KinesisFirehoseInput *AwsApplication_KinesisFirehoseInputProperty `field:"optional" json:"kinesisFirehoseInput" yaml:"kinesisFirehoseInput"`
	// kinesis_streams_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_streams_input AwsApplication#kinesis_streams_input}
	// Experimental.
	KinesisStreamsInput *AwsApplication_KinesisStreamsInputProperty `field:"optional" json:"kinesisStreamsInput" yaml:"kinesisStreamsInput"`
}

