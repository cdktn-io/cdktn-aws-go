package codepipeline


// Experimental.
type AwsCodepipeline_StageOnSuccessConditionProperty struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#rule AwsCodepipeline#rule}
	// Experimental.
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#result AwsCodepipeline#result}.
	// Experimental.
	Result *string `field:"optional" json:"result" yaml:"result"`
}

