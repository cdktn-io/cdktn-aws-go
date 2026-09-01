package awslightsail


// Experimental.
type AwsLightsailInstance_AddOnProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance#snapshot_time AwsLightsailInstance#snapshot_time}.
	// Experimental.
	SnapshotTime *string `field:"required" json:"snapshotTime" yaml:"snapshotTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance#status AwsLightsailInstance#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance#type AwsLightsailInstance#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

