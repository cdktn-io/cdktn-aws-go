package ebs


// Experimental.
type AwsSnapshotImport_ClientDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#comment AwsSnapshotImport#comment}.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#upload_end AwsSnapshotImport#upload_end}.
	// Experimental.
	UploadEnd *string `field:"optional" json:"uploadEnd" yaml:"uploadEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#upload_size AwsSnapshotImport#upload_size}.
	// Experimental.
	UploadSize *float64 `field:"optional" json:"uploadSize" yaml:"uploadSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#upload_start AwsSnapshotImport#upload_start}.
	// Experimental.
	UploadStart *string `field:"optional" json:"uploadStart" yaml:"uploadStart"`
}

