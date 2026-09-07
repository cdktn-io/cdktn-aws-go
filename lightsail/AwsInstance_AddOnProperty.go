package lightsail


// Experimental.
type AwsInstance_AddOnProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance#snapshot_time AwsInstance#snapshot_time}.
	// Experimental.
	SnapshotTime *string `field:"required" json:"snapshotTime" yaml:"snapshotTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance#status AwsInstance#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance#type AwsInstance#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

