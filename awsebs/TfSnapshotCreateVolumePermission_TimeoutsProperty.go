package awsebs


// Experimental.
type TfSnapshotCreateVolumePermission_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/snapshot_create_volume_permission#create TfSnapshotCreateVolumePermission#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/snapshot_create_volume_permission#delete TfSnapshotCreateVolumePermission#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

