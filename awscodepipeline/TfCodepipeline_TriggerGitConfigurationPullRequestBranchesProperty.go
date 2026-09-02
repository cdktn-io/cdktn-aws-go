package awscodepipeline


// Experimental.
type TfCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#excludes TfCodepipeline#excludes}.
	// Experimental.
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#includes TfCodepipeline#includes}.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
}

