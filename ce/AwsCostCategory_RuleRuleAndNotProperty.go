package ce


// Experimental.
type AwsCostCategory_RuleRuleAndNotProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#cost_category AwsCostCategory#cost_category}
	// Experimental.
	CostCategory *AwsCostCategory_RuleRuleAndNotCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#dimension AwsCostCategory#dimension}
	// Experimental.
	Dimension *AwsCostCategory_RuleRuleAndNotDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#tags AwsCostCategory#tags}
	// Experimental.
	Tags *AwsCostCategory_RuleRuleAndNotTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

