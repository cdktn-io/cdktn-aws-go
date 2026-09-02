package awsebs


// Experimental.
type TfSnapshotImport_DiskContainerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#format TfSnapshotImport#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#description TfSnapshotImport#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#url TfSnapshotImport#url}.
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// user_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#user_bucket TfSnapshotImport#user_bucket}
	// Experimental.
	UserBucket *TfSnapshotImport_UserBucketProperty `field:"optional" json:"userBucket" yaml:"userBucket"`
}

