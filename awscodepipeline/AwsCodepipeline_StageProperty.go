package awscodepipeline


// Experimental.
type AwsCodepipeline_StageProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#action AwsCodepipeline#action}
	// Experimental.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#name AwsCodepipeline#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// before_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#before_entry AwsCodepipeline#before_entry}
	// Experimental.
	BeforeEntry *AwsCodepipeline_BeforeEntryProperty `field:"optional" json:"beforeEntry" yaml:"beforeEntry"`
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#on_failure AwsCodepipeline#on_failure}
	// Experimental.
	OnFailure *AwsCodepipeline_OnFailureProperty `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#on_success AwsCodepipeline#on_success}
	// Experimental.
	OnSuccess *AwsCodepipeline_OnSuccessProperty `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

