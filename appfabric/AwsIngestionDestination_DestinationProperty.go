package appfabric


// Experimental.
type AwsIngestionDestination_DestinationProperty struct {
	// firehose_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#firehose_stream AwsIngestionDestination#firehose_stream}
	// Experimental.
	FirehoseStream interface{} `field:"optional" json:"firehoseStream" yaml:"firehoseStream"`
	// s3_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#s3_bucket AwsIngestionDestination#s3_bucket}
	// Experimental.
	S3Bucket interface{} `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

