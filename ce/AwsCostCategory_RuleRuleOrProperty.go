package ce


// Experimental.
type AwsCostCategory_RuleRuleOrProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#and AwsCostCategory#and}
	// Experimental.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#cost_category AwsCostCategory#cost_category}
	// Experimental.
	CostCategory *AwsCostCategory_RuleRuleOrCostCategoryProperty `field:"optional" json:"costCategory" yaml:"costCategory"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#dimension AwsCostCategory#dimension}
	// Experimental.
	Dimension *AwsCostCategory_RuleRuleOrDimensionProperty `field:"optional" json:"dimension" yaml:"dimension"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#not AwsCostCategory#not}
	// Experimental.
	Not *AwsCostCategory_RuleRuleOrNotProperty `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#or AwsCostCategory#or}
	// Experimental.
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ce_cost_category#tags AwsCostCategory#tags}
	// Experimental.
	Tags *AwsCostCategory_RuleRuleOrTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

