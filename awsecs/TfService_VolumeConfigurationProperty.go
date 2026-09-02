package awsecs


// Experimental.
type TfService_VolumeConfigurationProperty struct {
	// managed_ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#managed_ebs_volume TfService#managed_ebs_volume}
	// Experimental.
	ManagedEbsVolume *TfService_ManagedEbsVolumeProperty `field:"required" json:"managedEbsVolume" yaml:"managedEbsVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#name TfService#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

