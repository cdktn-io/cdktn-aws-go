package awsfsx


// Experimental.
type AwsFsxOpenzfsVolume_OriginSnapshotProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#copy_strategy AwsFsxOpenzfsVolume#copy_strategy}.
	// Experimental.
	CopyStrategy *string `field:"required" json:"copyStrategy" yaml:"copyStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#snapshot_arn AwsFsxOpenzfsVolume#snapshot_arn}.
	// Experimental.
	SnapshotArn *string `field:"required" json:"snapshotArn" yaml:"snapshotArn"`
}

