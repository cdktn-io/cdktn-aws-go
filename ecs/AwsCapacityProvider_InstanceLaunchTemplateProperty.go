package ecs


// Experimental.
type AwsCapacityProvider_InstanceLaunchTemplateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#ec2_instance_profile_arn AwsCapacityProvider#ec2_instance_profile_arn}.
	// Experimental.
	Ec2InstanceProfileArn *string `field:"required" json:"ec2InstanceProfileArn" yaml:"ec2InstanceProfileArn"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#network_configuration AwsCapacityProvider#network_configuration}
	// Experimental.
	NetworkConfiguration *AwsCapacityProvider_NetworkConfigurationProperty `field:"required" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#capacity_option_type AwsCapacityProvider#capacity_option_type}.
	// Experimental.
	CapacityOptionType *string `field:"optional" json:"capacityOptionType" yaml:"capacityOptionType"`
	// capacity_reservations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#capacity_reservations AwsCapacityProvider#capacity_reservations}
	// Experimental.
	CapacityReservations *AwsCapacityProvider_CapacityReservationsProperty `field:"optional" json:"capacityReservations" yaml:"capacityReservations"`
	// instance_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#instance_requirements AwsCapacityProvider#instance_requirements}
	// Experimental.
	InstanceRequirements *AwsCapacityProvider_InstanceRequirementsProperty `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// local_storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#local_storage_configuration AwsCapacityProvider#local_storage_configuration}
	// Experimental.
	LocalStorageConfiguration *AwsCapacityProvider_LocalStorageConfigurationProperty `field:"optional" json:"localStorageConfiguration" yaml:"localStorageConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#monitoring AwsCapacityProvider#monitoring}.
	// Experimental.
	Monitoring *string `field:"optional" json:"monitoring" yaml:"monitoring"`
	// storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#storage_configuration AwsCapacityProvider#storage_configuration}
	// Experimental.
	StorageConfiguration *AwsCapacityProvider_StorageConfigurationProperty `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
}

