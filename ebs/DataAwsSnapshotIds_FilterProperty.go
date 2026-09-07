package ebs


// Experimental.
type DataAwsSnapshotIds_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ebs_snapshot_ids#name DataAwsSnapshotIds#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ebs_snapshot_ids#values DataAwsSnapshotIds#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

