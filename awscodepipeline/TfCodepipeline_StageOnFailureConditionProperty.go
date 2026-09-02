package awscodepipeline


// Experimental.
type TfCodepipeline_StageOnFailureConditionProperty struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#rule TfCodepipeline#rule}
	// Experimental.
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#result TfCodepipeline#result}.
	// Experimental.
	Result *string `field:"optional" json:"result" yaml:"result"`
}

