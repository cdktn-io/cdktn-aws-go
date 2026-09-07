package codepipeline


// Experimental.
type AwsCodepipeline_TriggerProperty struct {
	// git_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#git_configuration AwsCodepipeline#git_configuration}
	// Experimental.
	GitConfiguration *AwsCodepipeline_TriggerGitConfigurationProperty `field:"required" json:"gitConfiguration" yaml:"gitConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#provider_type AwsCodepipeline#provider_type}.
	// Experimental.
	ProviderType *string `field:"required" json:"providerType" yaml:"providerType"`
}

