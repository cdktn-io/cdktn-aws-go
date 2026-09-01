package awsec2imagebuilder


// Experimental.
type AwsImagebuilderDistributionConfiguration_ContainerDistributionConfigurationProperty struct {
	// target_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#target_repository AwsImagebuilderDistributionConfiguration#target_repository}
	// Experimental.
	TargetRepository *AwsImagebuilderDistributionConfiguration_TargetRepositoryProperty `field:"required" json:"targetRepository" yaml:"targetRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#container_tags AwsImagebuilderDistributionConfiguration#container_tags}.
	// Experimental.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#description AwsImagebuilderDistributionConfiguration#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

