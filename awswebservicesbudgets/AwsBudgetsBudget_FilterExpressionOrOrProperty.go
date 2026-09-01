package awswebservicesbudgets


// Experimental.
type AwsBudgetsBudget_FilterExpressionOrOrProperty struct {
	// cost_categories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_categories AwsBudgetsBudget#cost_categories}
	// Experimental.
	CostCategories *AwsBudgetsBudget_FilterExpressionOrOrCostCategoriesProperty `field:"optional" json:"costCategories" yaml:"costCategories"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#dimensions AwsBudgetsBudget#dimensions}
	// Experimental.
	Dimensions *AwsBudgetsBudget_FilterExpressionOrOrDimensionsProperty `field:"optional" json:"dimensions" yaml:"dimensions"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags AwsBudgetsBudget#tags}
	// Experimental.
	Tags *AwsBudgetsBudget_FilterExpressionOrOrTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

