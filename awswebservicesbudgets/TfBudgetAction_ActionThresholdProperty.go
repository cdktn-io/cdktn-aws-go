package awswebservicesbudgets


// Experimental.
type TfBudgetAction_ActionThresholdProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#action_threshold_type TfBudgetAction#action_threshold_type}.
	// Experimental.
	ActionThresholdType *string `field:"required" json:"actionThresholdType" yaml:"actionThresholdType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#action_threshold_value TfBudgetAction#action_threshold_value}.
	// Experimental.
	ActionThresholdValue *float64 `field:"required" json:"actionThresholdValue" yaml:"actionThresholdValue"`
}

