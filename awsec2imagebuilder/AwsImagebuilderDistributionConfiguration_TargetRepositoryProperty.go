package awsec2imagebuilder


// Experimental.
type AwsImagebuilderDistributionConfiguration_TargetRepositoryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#repository_name AwsImagebuilderDistributionConfiguration#repository_name}.
	// Experimental.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#service AwsImagebuilderDistributionConfiguration#service}.
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
}

