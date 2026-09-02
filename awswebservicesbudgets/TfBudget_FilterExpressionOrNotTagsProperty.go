package awswebservicesbudgets


// Experimental.
type TfBudget_FilterExpressionOrNotTagsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#key TfBudget#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#match_options TfBudget#match_options}.
	// Experimental.
	MatchOptions *[]*string `field:"optional" json:"matchOptions" yaml:"matchOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#values TfBudget#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

