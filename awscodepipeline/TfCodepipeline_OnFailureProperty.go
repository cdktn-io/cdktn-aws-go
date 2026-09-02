package awscodepipeline


// Experimental.
type TfCodepipeline_OnFailureProperty struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#condition TfCodepipeline#condition}
	// Experimental.
	Condition *TfCodepipeline_StageOnFailureConditionProperty `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#result TfCodepipeline#result}.
	// Experimental.
	Result *string `field:"optional" json:"result" yaml:"result"`
	// retry_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#retry_configuration TfCodepipeline#retry_configuration}
	// Experimental.
	RetryConfiguration *TfCodepipeline_RetryConfigurationProperty `field:"optional" json:"retryConfiguration" yaml:"retryConfiguration"`
}

