package awsrds


// Experimental.
type TfDbInstance_S3ImportProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#bucket_name TfDbInstance#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#ingestion_role TfDbInstance#ingestion_role}.
	// Experimental.
	IngestionRole *string `field:"required" json:"ingestionRole" yaml:"ingestionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#source_engine TfDbInstance#source_engine}.
	// Experimental.
	SourceEngine *string `field:"required" json:"sourceEngine" yaml:"sourceEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#source_engine_version TfDbInstance#source_engine_version}.
	// Experimental.
	SourceEngineVersion *string `field:"required" json:"sourceEngineVersion" yaml:"sourceEngineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#bucket_prefix TfDbInstance#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

