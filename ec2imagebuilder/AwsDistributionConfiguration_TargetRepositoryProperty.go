package ec2imagebuilder


// Experimental.
type AwsDistributionConfiguration_TargetRepositoryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#repository_name AwsDistributionConfiguration#repository_name}.
	// Experimental.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#service AwsDistributionConfiguration#service}.
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
}

