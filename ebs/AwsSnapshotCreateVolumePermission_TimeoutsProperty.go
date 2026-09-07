package ebs


// Experimental.
type AwsSnapshotCreateVolumePermission_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/snapshot_create_volume_permission#create AwsSnapshotCreateVolumePermission#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/snapshot_create_volume_permission#delete AwsSnapshotCreateVolumePermission#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

