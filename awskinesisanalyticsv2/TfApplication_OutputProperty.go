package awskinesisanalyticsv2


// Experimental.
type TfApplication_OutputProperty struct {
	// destination_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#destination_schema TfApplication#destination_schema}
	// Experimental.
	DestinationSchema *TfApplication_DestinationSchemaProperty `field:"required" json:"destinationSchema" yaml:"destinationSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#name TfApplication#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// kinesis_firehose_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_firehose_output TfApplication#kinesis_firehose_output}
	// Experimental.
	KinesisFirehoseOutput *TfApplication_KinesisFirehoseOutputProperty `field:"optional" json:"kinesisFirehoseOutput" yaml:"kinesisFirehoseOutput"`
	// kinesis_streams_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_streams_output TfApplication#kinesis_streams_output}
	// Experimental.
	KinesisStreamsOutput *TfApplication_KinesisStreamsOutputProperty `field:"optional" json:"kinesisStreamsOutput" yaml:"kinesisStreamsOutput"`
	// lambda_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#lambda_output TfApplication#lambda_output}
	// Experimental.
	LambdaOutput *TfApplication_LambdaOutputProperty `field:"optional" json:"lambdaOutput" yaml:"lambdaOutput"`
}

