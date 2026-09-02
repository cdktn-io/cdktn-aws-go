package awsec2imagebuilder


// Experimental.
type TfInfrastructureConfiguration_PlacementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_infrastructure_configuration#availability_zone TfInfrastructureConfiguration#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_infrastructure_configuration#host_id TfInfrastructureConfiguration#host_id}.
	// Experimental.
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_infrastructure_configuration#host_resource_group_arn TfInfrastructureConfiguration#host_resource_group_arn}.
	// Experimental.
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_infrastructure_configuration#tenancy TfInfrastructureConfiguration#tenancy}.
	// Experimental.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

