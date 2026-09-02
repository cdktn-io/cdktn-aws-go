package awswebservicesbudgets


// Experimental.
type TfBudget_FilterExpressionNotProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#and TfBudget#and}
	// Experimental.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_categories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_categories TfBudget#cost_categories}
	// Experimental.
	CostCategories *TfBudget_FilterExpressionNotCostCategoriesProperty `field:"optional" json:"costCategories" yaml:"costCategories"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#dimensions TfBudget#dimensions}
	// Experimental.
	Dimensions *TfBudget_FilterExpressionNotDimensionsProperty `field:"optional" json:"dimensions" yaml:"dimensions"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#not TfBudget#not}
	// Experimental.
	Not *TfBudget_FilterExpressionNotNotProperty `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#or TfBudget#or}
	// Experimental.
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags TfBudget#tags}
	// Experimental.
	Tags *TfBudget_FilterExpressionNotTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

