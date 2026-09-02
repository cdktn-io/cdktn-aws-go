package awsce


// Experimental.
type TfCostCategory_RuleRuleAndAndProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#cost_category TfCostCategory#cost_category}
	// Experimental.
	CostCategory *TfCostCategory_RuleRuleAndAndCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#dimension TfCostCategory#dimension}
	// Experimental.
	Dimension *TfCostCategory_RuleRuleAndAndDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#tags TfCostCategory#tags}
	// Experimental.
	Tags *TfCostCategory_RuleRuleAndAndTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

