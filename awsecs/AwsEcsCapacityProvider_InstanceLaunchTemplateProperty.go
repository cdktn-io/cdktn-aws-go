package awsecs


// Experimental.
type AwsEcsCapacityProvider_InstanceLaunchTemplateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#ec2_instance_profile_arn AwsEcsCapacityProvider#ec2_instance_profile_arn}.
	// Experimental.
	Ec2InstanceProfileArn *string `field:"required" json:"ec2InstanceProfileArn" yaml:"ec2InstanceProfileArn"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#network_configuration AwsEcsCapacityProvider#network_configuration}
	// Experimental.
	NetworkConfiguration *AwsEcsCapacityProvider_NetworkConfigurationProperty `field:"required" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#capacity_option_type AwsEcsCapacityProvider#capacity_option_type}.
	// Experimental.
	CapacityOptionType *string `field:"optional" json:"capacityOptionType" yaml:"capacityOptionType"`
	// capacity_reservations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#capacity_reservations AwsEcsCapacityProvider#capacity_reservations}
	// Experimental.
	CapacityReservations *AwsEcsCapacityProvider_CapacityReservationsProperty `field:"optional" json:"capacityReservations" yaml:"capacityReservations"`
	// instance_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#instance_requirements AwsEcsCapacityProvider#instance_requirements}
	// Experimental.
	InstanceRequirements *AwsEcsCapacityProvider_InstanceRequirementsProperty `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// local_storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#local_storage_configuration AwsEcsCapacityProvider#local_storage_configuration}
	// Experimental.
	LocalStorageConfiguration *AwsEcsCapacityProvider_LocalStorageConfigurationProperty `field:"optional" json:"localStorageConfiguration" yaml:"localStorageConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#monitoring AwsEcsCapacityProvider#monitoring}.
	// Experimental.
	Monitoring *string `field:"optional" json:"monitoring" yaml:"monitoring"`
	// storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#storage_configuration AwsEcsCapacityProvider#storage_configuration}
	// Experimental.
	StorageConfiguration *AwsEcsCapacityProvider_StorageConfigurationProperty `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
}

