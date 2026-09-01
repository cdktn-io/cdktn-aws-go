package awsecs


// Experimental.
type AwsEcsService_VolumeConfigurationProperty struct {
	// managed_ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#managed_ebs_volume AwsEcsService#managed_ebs_volume}
	// Experimental.
	ManagedEbsVolume *AwsEcsService_ManagedEbsVolumeProperty `field:"required" json:"managedEbsVolume" yaml:"managedEbsVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#name AwsEcsService#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

