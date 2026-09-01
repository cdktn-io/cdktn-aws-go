package awswebservicesbudgets


// Experimental.
type AwsBudgetsBudgetAction_SubscriberProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#address AwsBudgetsBudgetAction#address}.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#subscription_type AwsBudgetsBudgetAction#subscription_type}.
	// Experimental.
	SubscriptionType *string `field:"required" json:"subscriptionType" yaml:"subscriptionType"`
}

