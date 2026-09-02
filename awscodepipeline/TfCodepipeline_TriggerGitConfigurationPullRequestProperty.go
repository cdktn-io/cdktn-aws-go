package awscodepipeline


// Experimental.
type TfCodepipeline_TriggerGitConfigurationPullRequestProperty struct {
	// branches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#branches TfCodepipeline#branches}
	// Experimental.
	Branches *TfCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty `field:"optional" json:"branches" yaml:"branches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#events TfCodepipeline#events}.
	// Experimental.
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
	// file_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#file_paths TfCodepipeline#file_paths}
	// Experimental.
	FilePaths *TfCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty `field:"optional" json:"filePaths" yaml:"filePaths"`
}

