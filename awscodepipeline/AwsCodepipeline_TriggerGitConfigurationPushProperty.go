package awscodepipeline


// Experimental.
type AwsCodepipeline_TriggerGitConfigurationPushProperty struct {
	// branches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#branches AwsCodepipeline#branches}
	// Experimental.
	Branches *AwsCodepipeline_TriggerGitConfigurationPushBranchesProperty `field:"optional" json:"branches" yaml:"branches"`
	// file_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#file_paths AwsCodepipeline#file_paths}
	// Experimental.
	FilePaths *AwsCodepipeline_TriggerGitConfigurationPushFilePathsProperty `field:"optional" json:"filePaths" yaml:"filePaths"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#tags AwsCodepipeline#tags}
	// Experimental.
	Tags *AwsCodepipeline_TriggerGitConfigurationPushTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

