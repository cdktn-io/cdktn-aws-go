package awss3tables


// Experimental.
type TfTableReplication_DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_replication#destination_table_bucket_arn TfTableReplication#destination_table_bucket_arn}.
	// Experimental.
	DestinationTableBucketArn *string `field:"required" json:"destinationTableBucketArn" yaml:"destinationTableBucketArn"`
}

