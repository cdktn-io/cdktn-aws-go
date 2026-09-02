package awscodepipeline


// Experimental.
type TfCodepipeline_StageProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#action TfCodepipeline#action}
	// Experimental.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#name TfCodepipeline#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// before_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#before_entry TfCodepipeline#before_entry}
	// Experimental.
	BeforeEntry *TfCodepipeline_BeforeEntryProperty `field:"optional" json:"beforeEntry" yaml:"beforeEntry"`
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#on_failure TfCodepipeline#on_failure}
	// Experimental.
	OnFailure *TfCodepipeline_OnFailureProperty `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#on_success TfCodepipeline#on_success}
	// Experimental.
	OnSuccess *TfCodepipeline_OnSuccessProperty `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

