package awsce


// Experimental.
type TfCostCategory_RuleRuleProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#and TfCostCategory#and}
	// Experimental.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#cost_category TfCostCategory#cost_category}
	// Experimental.
	CostCategory *TfCostCategory_RuleRuleCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#dimension TfCostCategory#dimension}
	// Experimental.
	Dimension *TfCostCategory_RuleRuleDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#not TfCostCategory#not}
	// Experimental.
	Not *TfCostCategory_RuleRuleNotProperty `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#or TfCostCategory#or}
	// Experimental.
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#tags TfCostCategory#tags}
	// Experimental.
	Tags *TfCostCategory_RuleRuleTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

