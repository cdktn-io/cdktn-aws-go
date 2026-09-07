package webservicesbudgets


// Experimental.
type AwsBudget_FilterExpressionNotProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#and AwsBudget#and}
	// Experimental.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_categories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_categories AwsBudget#cost_categories}
	// Experimental.
	CostCategories *AwsBudget_FilterExpressionNotCostCategoriesProperty `field:"optional" json:"costCategories" yaml:"costCategories"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#dimensions AwsBudget#dimensions}
	// Experimental.
	Dimensions *AwsBudget_FilterExpressionNotDimensionsProperty `field:"optional" json:"dimensions" yaml:"dimensions"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#not AwsBudget#not}
	// Experimental.
	Not *AwsBudget_FilterExpressionNotNotProperty `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#or AwsBudget#or}
	// Experimental.
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags AwsBudget#tags}
	// Experimental.
	Tags *AwsBudget_FilterExpressionNotTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

