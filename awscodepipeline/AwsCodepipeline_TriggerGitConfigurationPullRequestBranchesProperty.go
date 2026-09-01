package awscodepipeline


// Experimental.
type AwsCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#excludes AwsCodepipeline#excludes}.
	// Experimental.
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#includes AwsCodepipeline#includes}.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
}

