package awswebservicesbudgets

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBudgetsBudgetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#budget_type AwsBudgetsBudget#budget_type}.
	// Experimental.
	BudgetType *string `field:"required" json:"budgetType" yaml:"budgetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#time_unit AwsBudgetsBudget#time_unit}.
	// Experimental.
	TimeUnit *string `field:"required" json:"timeUnit" yaml:"timeUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#account_id AwsBudgetsBudget#account_id}.
	// Experimental.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// auto_adjust_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#auto_adjust_data AwsBudgetsBudget#auto_adjust_data}
	// Experimental.
	AutoAdjustData *AwsBudgetsBudget_AutoAdjustDataProperty `field:"optional" json:"autoAdjustData" yaml:"autoAdjustData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#billing_view_arn AwsBudgetsBudget#billing_view_arn}.
	// Experimental.
	BillingViewArn *string `field:"optional" json:"billingViewArn" yaml:"billingViewArn"`
	// cost_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_filter AwsBudgetsBudget#cost_filter}
	// Experimental.
	CostFilter interface{} `field:"optional" json:"costFilter" yaml:"costFilter"`
	// cost_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_types AwsBudgetsBudget#cost_types}
	// Experimental.
	CostTypes *AwsBudgetsBudget_CostTypesProperty `field:"optional" json:"costTypes" yaml:"costTypes"`
	// filter_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#filter_expression AwsBudgetsBudget#filter_expression}
	// Experimental.
	FilterExpression *AwsBudgetsBudget_FilterExpressionProperty `field:"optional" json:"filterExpression" yaml:"filterExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#id AwsBudgetsBudget#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#limit_amount AwsBudgetsBudget#limit_amount}.
	// Experimental.
	LimitAmount *string `field:"optional" json:"limitAmount" yaml:"limitAmount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#limit_unit AwsBudgetsBudget#limit_unit}.
	// Experimental.
	LimitUnit *string `field:"optional" json:"limitUnit" yaml:"limitUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#metrics AwsBudgetsBudget#metrics}.
	// Experimental.
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#name AwsBudgetsBudget#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#name_prefix AwsBudgetsBudget#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#notification AwsBudgetsBudget#notification}
	// Experimental.
	Notification interface{} `field:"optional" json:"notification" yaml:"notification"`
	// planned_limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#planned_limit AwsBudgetsBudget#planned_limit}
	// Experimental.
	PlannedLimit interface{} `field:"optional" json:"plannedLimit" yaml:"plannedLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags AwsBudgetsBudget#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags_all AwsBudgetsBudget#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#time_period_end AwsBudgetsBudget#time_period_end}.
	// Experimental.
	TimePeriodEnd *string `field:"optional" json:"timePeriodEnd" yaml:"timePeriodEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#time_period_start AwsBudgetsBudget#time_period_start}.
	// Experimental.
	TimePeriodStart *string `field:"optional" json:"timePeriodStart" yaml:"timePeriodStart"`
}

