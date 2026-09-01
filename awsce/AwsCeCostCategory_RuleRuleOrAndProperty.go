package awsce


// Experimental.
type AwsCeCostCategory_RuleRuleOrAndProperty struct {
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#cost_category AwsCeCostCategory#cost_category}
	// Experimental.
	CostCategory *AwsCeCostCategory_RuleRuleOrAndCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#dimension AwsCeCostCategory#dimension}
	// Experimental.
	Dimension *AwsCeCostCategory_RuleRuleOrAndDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#tags AwsCeCostCategory#tags}
	// Experimental.
	Tags *AwsCeCostCategory_RuleRuleOrAndTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

