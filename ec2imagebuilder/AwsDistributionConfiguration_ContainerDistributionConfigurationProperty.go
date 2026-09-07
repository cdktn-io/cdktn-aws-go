package ec2imagebuilder


// Experimental.
type AwsDistributionConfiguration_ContainerDistributionConfigurationProperty struct {
	// target_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#target_repository AwsDistributionConfiguration#target_repository}
	// Experimental.
	TargetRepository *AwsDistributionConfiguration_TargetRepositoryProperty `field:"required" json:"targetRepository" yaml:"targetRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#container_tags AwsDistributionConfiguration#container_tags}.
	// Experimental.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#description AwsDistributionConfiguration#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

