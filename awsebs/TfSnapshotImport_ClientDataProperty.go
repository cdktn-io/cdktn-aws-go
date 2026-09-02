package awsebs


// Experimental.
type TfSnapshotImport_ClientDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#comment TfSnapshotImport#comment}.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#upload_end TfSnapshotImport#upload_end}.
	// Experimental.
	UploadEnd *string `field:"optional" json:"uploadEnd" yaml:"uploadEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#upload_size TfSnapshotImport#upload_size}.
	// Experimental.
	UploadSize *float64 `field:"optional" json:"uploadSize" yaml:"uploadSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#upload_start TfSnapshotImport#upload_start}.
	// Experimental.
	UploadStart *string `field:"optional" json:"uploadStart" yaml:"uploadStart"`
}

