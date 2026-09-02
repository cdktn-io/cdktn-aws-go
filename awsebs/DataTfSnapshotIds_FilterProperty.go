package awsebs


// Experimental.
type DataTfSnapshotIds_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ebs_snapshot_ids#name DataTfSnapshotIds#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ebs_snapshot_ids#values DataTfSnapshotIds#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

