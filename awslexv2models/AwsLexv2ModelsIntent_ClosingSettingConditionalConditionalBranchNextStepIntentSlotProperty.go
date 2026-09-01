package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_ClosingSettingConditionalConditionalBranchNextStepIntentSlotProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#map_block_key AwsLexv2ModelsIntent#map_block_key}.
	// Experimental.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#shape AwsLexv2ModelsIntent#shape}.
	// Experimental.
	Shape *string `field:"optional" json:"shape" yaml:"shape"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#value AwsLexv2ModelsIntent#value}
	// Experimental.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

