package codepipeline


// Experimental.
type AwsCodepipeline_OnSuccessProperty struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#condition AwsCodepipeline#condition}
	// Experimental.
	Condition *AwsCodepipeline_StageOnSuccessConditionProperty `field:"required" json:"condition" yaml:"condition"`
}

