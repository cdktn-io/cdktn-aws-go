package webservicesbudgets


// Experimental.
type AwsBudget_AutoAdjustDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#auto_adjust_type AwsBudget#auto_adjust_type}.
	// Experimental.
	AutoAdjustType *string `field:"required" json:"autoAdjustType" yaml:"autoAdjustType"`
	// historical_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#historical_options AwsBudget#historical_options}
	// Experimental.
	HistoricalOptions *AwsBudget_HistoricalOptionsProperty `field:"optional" json:"historicalOptions" yaml:"historicalOptions"`
}

