package awsce


// Experimental.
type TfCostCategory_RuleProperty struct {
	// inherited_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#inherited_value TfCostCategory#inherited_value}
	// Experimental.
	InheritedValue *TfCostCategory_InheritedValueProperty `field:"optional" json:"inheritedValue" yaml:"inheritedValue"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#rule TfCostCategory#rule}
	// Experimental.
	Rule *TfCostCategory_RuleRuleProperty `field:"optional" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#type TfCostCategory#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#value TfCostCategory#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

