package codepipeline


// Experimental.
type AwsCodepipeline_OnFailureProperty struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#condition AwsCodepipeline#condition}
	// Experimental.
	Condition *AwsCodepipeline_StageOnFailureConditionProperty `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#result AwsCodepipeline#result}.
	// Experimental.
	Result *string `field:"optional" json:"result" yaml:"result"`
	// retry_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#retry_configuration AwsCodepipeline#retry_configuration}
	// Experimental.
	RetryConfiguration *AwsCodepipeline_RetryConfigurationProperty `field:"optional" json:"retryConfiguration" yaml:"retryConfiguration"`
}

