package awsebs


// Experimental.
type AwsEbsSnapshotImport_DiskContainerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#format AwsEbsSnapshotImport#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#description AwsEbsSnapshotImport#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#url AwsEbsSnapshotImport#url}.
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// user_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#user_bucket AwsEbsSnapshotImport#user_bucket}
	// Experimental.
	UserBucket *AwsEbsSnapshotImport_UserBucketProperty `field:"optional" json:"userBucket" yaml:"userBucket"`
}

