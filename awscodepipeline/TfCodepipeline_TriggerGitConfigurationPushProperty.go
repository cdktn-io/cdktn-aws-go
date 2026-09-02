package awscodepipeline


// Experimental.
type TfCodepipeline_TriggerGitConfigurationPushProperty struct {
	// branches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#branches TfCodepipeline#branches}
	// Experimental.
	Branches *TfCodepipeline_TriggerGitConfigurationPushBranchesProperty `field:"optional" json:"branches" yaml:"branches"`
	// file_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#file_paths TfCodepipeline#file_paths}
	// Experimental.
	FilePaths *TfCodepipeline_TriggerGitConfigurationPushFilePathsProperty `field:"optional" json:"filePaths" yaml:"filePaths"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#tags TfCodepipeline#tags}
	// Experimental.
	Tags *TfCodepipeline_TriggerGitConfigurationPushTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

