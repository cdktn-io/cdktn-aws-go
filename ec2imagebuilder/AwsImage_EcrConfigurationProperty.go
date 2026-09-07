package ec2imagebuilder


// Experimental.
type AwsImage_EcrConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#container_tags AwsImage#container_tags}.
	// Experimental.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image#repository_name AwsImage#repository_name}.
	// Experimental.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}

