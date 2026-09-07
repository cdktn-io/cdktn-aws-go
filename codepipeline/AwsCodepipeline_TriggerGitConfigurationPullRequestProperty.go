package codepipeline


// Experimental.
type AwsCodepipeline_TriggerGitConfigurationPullRequestProperty struct {
	// branches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#branches AwsCodepipeline#branches}
	// Experimental.
	Branches *AwsCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty `field:"optional" json:"branches" yaml:"branches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#events AwsCodepipeline#events}.
	// Experimental.
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
	// file_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#file_paths AwsCodepipeline#file_paths}
	// Experimental.
	FilePaths *AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty `field:"optional" json:"filePaths" yaml:"filePaths"`
}

