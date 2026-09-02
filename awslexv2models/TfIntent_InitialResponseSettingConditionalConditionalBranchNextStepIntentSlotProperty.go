package awslexv2models


// Experimental.
type TfIntent_InitialResponseSettingConditionalConditionalBranchNextStepIntentSlotProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#map_block_key TfIntent#map_block_key}.
	// Experimental.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#shape TfIntent#shape}.
	// Experimental.
	Shape *string `field:"optional" json:"shape" yaml:"shape"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#value TfIntent#value}
	// Experimental.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

