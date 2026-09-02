package awscodepipeline


// Experimental.
type TfCodepipeline_TriggerProperty struct {
	// git_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#git_configuration TfCodepipeline#git_configuration}
	// Experimental.
	GitConfiguration *TfCodepipeline_TriggerGitConfigurationProperty `field:"required" json:"gitConfiguration" yaml:"gitConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#provider_type TfCodepipeline#provider_type}.
	// Experimental.
	ProviderType *string `field:"required" json:"providerType" yaml:"providerType"`
}

