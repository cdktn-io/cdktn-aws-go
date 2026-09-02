package awswebservicesbudgets


// Experimental.
type TfBudgetAction_SubscriberProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#address TfBudgetAction#address}.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#subscription_type TfBudgetAction#subscription_type}.
	// Experimental.
	SubscriptionType *string `field:"required" json:"subscriptionType" yaml:"subscriptionType"`
}

