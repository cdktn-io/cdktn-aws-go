package awsce


// Experimental.
type TfCostCategory_RuleRuleOrAndProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#cost_category TfCostCategory#cost_category}
	// Experimental.
	CostCategory *TfCostCategory_RuleRuleOrAndCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#dimension TfCostCategory#dimension}
	// Experimental.
	Dimension *TfCostCategory_RuleRuleOrAndDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#tags TfCostCategory#tags}
	// Experimental.
	Tags *TfCostCategory_RuleRuleOrAndTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

