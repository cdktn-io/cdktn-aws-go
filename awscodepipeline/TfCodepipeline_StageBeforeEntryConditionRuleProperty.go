package awscodepipeline


// Experimental.
type TfCodepipeline_StageBeforeEntryConditionRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#name TfCodepipeline#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// rule_type_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#rule_type_id TfCodepipeline#rule_type_id}
	// Experimental.
	RuleTypeId *TfCodepipeline_StageBeforeEntryConditionRuleRuleTypeIdProperty `field:"required" json:"ruleTypeId" yaml:"ruleTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#commands TfCodepipeline#commands}.
	// Experimental.
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#configuration TfCodepipeline#configuration}.
	// Experimental.
	Configuration *map[string]*string `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#input_artifacts TfCodepipeline#input_artifacts}.
	// Experimental.
	InputArtifacts *[]*string `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#region TfCodepipeline#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#role_arn TfCodepipeline#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#timeout_in_minutes TfCodepipeline#timeout_in_minutes}.
	// Experimental.
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

