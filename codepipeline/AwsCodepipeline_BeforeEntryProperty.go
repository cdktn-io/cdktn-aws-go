package codepipeline


// Experimental.
type AwsCodepipeline_BeforeEntryProperty struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#condition AwsCodepipeline#condition}
	// Experimental.
	Condition *AwsCodepipeline_StageBeforeEntryConditionProperty `field:"required" json:"condition" yaml:"condition"`
}

