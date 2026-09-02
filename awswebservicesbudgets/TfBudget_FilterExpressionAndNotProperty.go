package awswebservicesbudgets


// Experimental.
type TfBudget_FilterExpressionAndNotProperty struct {
	// cost_categories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_categories TfBudget#cost_categories}
	// Experimental.
	CostCategories *TfBudget_FilterExpressionAndNotCostCategoriesProperty `field:"optional" json:"costCategories" yaml:"costCategories"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#dimensions TfBudget#dimensions}
	// Experimental.
	Dimensions *TfBudget_FilterExpressionAndNotDimensionsProperty `field:"optional" json:"dimensions" yaml:"dimensions"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags TfBudget#tags}
	// Experimental.
	Tags *TfBudget_FilterExpressionAndNotTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

