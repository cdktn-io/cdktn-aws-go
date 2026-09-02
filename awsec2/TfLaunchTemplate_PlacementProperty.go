package awsec2


// Experimental.
type TfLaunchTemplate_PlacementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#affinity TfLaunchTemplate#affinity}.
	// Experimental.
	Affinity *string `field:"optional" json:"affinity" yaml:"affinity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#availability_zone TfLaunchTemplate#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#group_id TfLaunchTemplate#group_id}.
	// Experimental.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#group_name TfLaunchTemplate#group_name}.
	// Experimental.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#host_id TfLaunchTemplate#host_id}.
	// Experimental.
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#host_resource_group_arn TfLaunchTemplate#host_resource_group_arn}.
	// Experimental.
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#partition_number TfLaunchTemplate#partition_number}.
	// Experimental.
	PartitionNumber *float64 `field:"optional" json:"partitionNumber" yaml:"partitionNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#spread_domain TfLaunchTemplate#spread_domain}.
	// Experimental.
	SpreadDomain *string `field:"optional" json:"spreadDomain" yaml:"spreadDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#tenancy TfLaunchTemplate#tenancy}.
	// Experimental.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

