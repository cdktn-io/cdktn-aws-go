package awsec2imagebuilder


// Experimental.
type TfDistributionConfiguration_ContainerDistributionConfigurationProperty struct {
	// target_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#target_repository TfDistributionConfiguration#target_repository}
	// Experimental.
	TargetRepository *TfDistributionConfiguration_TargetRepositoryProperty `field:"required" json:"targetRepository" yaml:"targetRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#container_tags TfDistributionConfiguration#container_tags}.
	// Experimental.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#description TfDistributionConfiguration#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

