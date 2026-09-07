package webservicesbudgets


// Experimental.
type AwsBudget_FilterExpressionOrOrProperty struct {
	// cost_categories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_categories AwsBudget#cost_categories}
	// Experimental.
	CostCategories *AwsBudget_FilterExpressionOrOrCostCategoriesProperty `field:"optional" json:"costCategories" yaml:"costCategories"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#dimensions AwsBudget#dimensions}
	// Experimental.
	Dimensions *AwsBudget_FilterExpressionOrOrDimensionsProperty `field:"optional" json:"dimensions" yaml:"dimensions"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags AwsBudget#tags}
	// Experimental.
	Tags *AwsBudget_FilterExpressionOrOrTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

