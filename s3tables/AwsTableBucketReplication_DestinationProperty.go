package s3tables


// Experimental.
type AwsTableBucketReplication_DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket_replication#destination_table_bucket_arn AwsTableBucketReplication#destination_table_bucket_arn}.
	// Experimental.
	DestinationTableBucketArn *string `field:"required" json:"destinationTableBucketArn" yaml:"destinationTableBucketArn"`
}

