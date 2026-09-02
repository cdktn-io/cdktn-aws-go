package awskinesisanalyticsv2


// Experimental.
type TfApplication_InputProperty struct {
	// input_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_schema TfApplication#input_schema}
	// Experimental.
	InputSchema *TfApplication_InputSchemaProperty `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#name_prefix TfApplication#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// input_parallelism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_parallelism TfApplication#input_parallelism}
	// Experimental.
	InputParallelism *TfApplication_InputParallelismProperty `field:"optional" json:"inputParallelism" yaml:"inputParallelism"`
	// input_processing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_processing_configuration TfApplication#input_processing_configuration}
	// Experimental.
	InputProcessingConfiguration *TfApplication_InputProcessingConfigurationProperty `field:"optional" json:"inputProcessingConfiguration" yaml:"inputProcessingConfiguration"`
	// input_starting_position_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input_starting_position_configuration TfApplication#input_starting_position_configuration}
	// Experimental.
	InputStartingPositionConfiguration interface{} `field:"optional" json:"inputStartingPositionConfiguration" yaml:"inputStartingPositionConfiguration"`
	// kinesis_firehose_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_firehose_input TfApplication#kinesis_firehose_input}
	// Experimental.
	KinesisFirehoseInput *TfApplication_KinesisFirehoseInputProperty `field:"optional" json:"kinesisFirehoseInput" yaml:"kinesisFirehoseInput"`
	// kinesis_streams_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_streams_input TfApplication#kinesis_streams_input}
	// Experimental.
	KinesisStreamsInput *TfApplication_KinesisStreamsInputProperty `field:"optional" json:"kinesisStreamsInput" yaml:"kinesisStreamsInput"`
}

