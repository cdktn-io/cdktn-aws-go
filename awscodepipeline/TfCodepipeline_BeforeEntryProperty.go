package awscodepipeline


// Experimental.
type TfCodepipeline_BeforeEntryProperty struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#condition TfCodepipeline#condition}
	// Experimental.
	Condition *TfCodepipeline_StageBeforeEntryConditionProperty `field:"required" json:"condition" yaml:"condition"`
}

