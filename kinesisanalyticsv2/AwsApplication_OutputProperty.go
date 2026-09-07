package kinesisanalyticsv2


// Experimental.
type AwsApplication_OutputProperty struct {
	// destination_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#destination_schema AwsApplication#destination_schema}
	// Experimental.
	DestinationSchema *AwsApplication_DestinationSchemaProperty `field:"required" json:"destinationSchema" yaml:"destinationSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#name AwsApplication#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// kinesis_firehose_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_firehose_output AwsApplication#kinesis_firehose_output}
	// Experimental.
	KinesisFirehoseOutput *AwsApplication_KinesisFirehoseOutputProperty `field:"optional" json:"kinesisFirehoseOutput" yaml:"kinesisFirehoseOutput"`
	// kinesis_streams_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#kinesis_streams_output AwsApplication#kinesis_streams_output}
	// Experimental.
	KinesisStreamsOutput *AwsApplication_KinesisStreamsOutputProperty `field:"optional" json:"kinesisStreamsOutput" yaml:"kinesisStreamsOutput"`
	// lambda_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#lambda_output AwsApplication#lambda_output}
	// Experimental.
	LambdaOutput *AwsApplication_LambdaOutputProperty `field:"optional" json:"lambdaOutput" yaml:"lambdaOutput"`
}

