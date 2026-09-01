package awswebservicesbudgets


// Experimental.
type AwsBudgetsBudgetAction_SsmActionDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#action_sub_type AwsBudgetsBudgetAction#action_sub_type}.
	// Experimental.
	ActionSubType *string `field:"required" json:"actionSubType" yaml:"actionSubType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#instance_ids AwsBudgetsBudgetAction#instance_ids}.
	// Experimental.
	InstanceIds *[]*string `field:"required" json:"instanceIds" yaml:"instanceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#region AwsBudgetsBudgetAction#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
}

