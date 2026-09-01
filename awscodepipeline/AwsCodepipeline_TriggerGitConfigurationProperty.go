package awscodepipeline


// Experimental.
type AwsCodepipeline_TriggerGitConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#source_action_name AwsCodepipeline#source_action_name}.
	// Experimental.
	SourceActionName *string `field:"required" json:"sourceActionName" yaml:"sourceActionName"`
	// pull_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#pull_request AwsCodepipeline#pull_request}
	// Experimental.
	PullRequest interface{} `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// push block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#push AwsCodepipeline#push}
	// Experimental.
	Push interface{} `field:"optional" json:"push" yaml:"push"`
}

