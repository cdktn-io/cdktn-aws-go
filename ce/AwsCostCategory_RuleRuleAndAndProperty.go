package ce


// Experimental.
type AwsCostCategory_RuleRuleAndAndProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#cost_category AwsCostCategory#cost_category}
	// Experimental.
	CostCategory *AwsCostCategory_RuleRuleAndAndCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#dimension AwsCostCategory#dimension}
	// Experimental.
	Dimension *AwsCostCategory_RuleRuleAndAndDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#tags AwsCostCategory#tags}
	// Experimental.
	Tags *AwsCostCategory_RuleRuleAndAndTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

