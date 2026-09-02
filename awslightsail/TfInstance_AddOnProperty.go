package awslightsail


// Experimental.
type TfInstance_AddOnProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance#snapshot_time TfInstance#snapshot_time}.
	// Experimental.
	SnapshotTime *string `field:"required" json:"snapshotTime" yaml:"snapshotTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance#status TfInstance#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance#type TfInstance#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

