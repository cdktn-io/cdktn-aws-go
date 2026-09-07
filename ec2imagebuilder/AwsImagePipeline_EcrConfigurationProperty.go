package ec2imagebuilder


// Experimental.
type AwsImagePipeline_EcrConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#container_tags AwsImagePipeline#container_tags}.
	// Experimental.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline#repository_name AwsImagePipeline#repository_name}.
	// Experimental.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}

