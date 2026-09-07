package webservicesbudgets


// Experimental.
type AwsBudget_FilterExpressionAndOrProperty struct {
	// cost_categories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_categories AwsBudget#cost_categories}
	// Experimental.
	CostCategories *AwsBudget_FilterExpressionAndOrCostCategoriesProperty `field:"optional" json:"costCategories" yaml:"costCategories"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#dimensions AwsBudget#dimensions}
	// Experimental.
	Dimensions *AwsBudget_FilterExpressionAndOrDimensionsProperty `field:"optional" json:"dimensions" yaml:"dimensions"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags AwsBudget#tags}
	// Experimental.
	Tags *AwsBudget_FilterExpressionAndOrTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

