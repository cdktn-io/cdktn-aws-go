package awsebs


// Experimental.
type TfSnapshotImport_UserBucketProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#s3_bucket TfSnapshotImport#s3_bucket}.
	// Experimental.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#s3_key TfSnapshotImport#s3_key}.
	// Experimental.
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}

