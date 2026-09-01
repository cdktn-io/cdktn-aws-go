package awswebservicesbudgets


// Experimental.
type AwsBudgetsBudgetAction_ScpActionDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#policy_id AwsBudgetsBudgetAction#policy_id}.
	// Experimental.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#target_ids AwsBudgetsBudgetAction#target_ids}.
	// Experimental.
	TargetIds *[]*string `field:"required" json:"targetIds" yaml:"targetIds"`
}

