package awsappfabric


// Experimental.
type TfIngestionDestination_S3BucketProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#bucket_name TfIngestionDestination#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#prefix TfIngestionDestination#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

