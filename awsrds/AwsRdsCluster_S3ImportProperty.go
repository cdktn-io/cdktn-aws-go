package awsrds


// Experimental.
type AwsRdsCluster_S3ImportProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#bucket_name AwsRdsCluster#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#ingestion_role AwsRdsCluster#ingestion_role}.
	// Experimental.
	IngestionRole *string `field:"required" json:"ingestionRole" yaml:"ingestionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#source_engine AwsRdsCluster#source_engine}.
	// Experimental.
	SourceEngine *string `field:"required" json:"sourceEngine" yaml:"sourceEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#source_engine_version AwsRdsCluster#source_engine_version}.
	// Experimental.
	SourceEngineVersion *string `field:"required" json:"sourceEngineVersion" yaml:"sourceEngineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#bucket_prefix AwsRdsCluster#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

