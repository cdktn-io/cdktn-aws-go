package webservicesbudgets

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBudgetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#budget_type AwsBudget#budget_type}.
	// Experimental.
	BudgetType *string `field:"required" json:"budgetType" yaml:"budgetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#time_unit AwsBudget#time_unit}.
	// Experimental.
	TimeUnit *string `field:"required" json:"timeUnit" yaml:"timeUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#account_id AwsBudget#account_id}.
	// Experimental.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// auto_adjust_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#auto_adjust_data AwsBudget#auto_adjust_data}
	// Experimental.
	AutoAdjustData *AwsBudget_AutoAdjustDataProperty `field:"optional" json:"autoAdjustData" yaml:"autoAdjustData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#billing_view_arn AwsBudget#billing_view_arn}.
	// Experimental.
	BillingViewArn *string `field:"optional" json:"billingViewArn" yaml:"billingViewArn"`
	// cost_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_filter AwsBudget#cost_filter}
	// Experimental.
	CostFilter interface{} `field:"optional" json:"costFilter" yaml:"costFilter"`
	// cost_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_types AwsBudget#cost_types}
	// Experimental.
	CostTypes *AwsBudget_CostTypesProperty `field:"optional" json:"costTypes" yaml:"costTypes"`
	// filter_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#filter_expression AwsBudget#filter_expression}
	// Experimental.
	FilterExpression *AwsBudget_FilterExpressionProperty `field:"optional" json:"filterExpression" yaml:"filterExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#id AwsBudget#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#limit_amount AwsBudget#limit_amount}.
	// Experimental.
	LimitAmount *string `field:"optional" json:"limitAmount" yaml:"limitAmount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#limit_unit AwsBudget#limit_unit}.
	// Experimental.
	LimitUnit *string `field:"optional" json:"limitUnit" yaml:"limitUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#metrics AwsBudget#metrics}.
	// Experimental.
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#name AwsBudget#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#name_prefix AwsBudget#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#notification AwsBudget#notification}
	// Experimental.
	Notification interface{} `field:"optional" json:"notification" yaml:"notification"`
	// planned_limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#planned_limit AwsBudget#planned_limit}
	// Experimental.
	PlannedLimit interface{} `field:"optional" json:"plannedLimit" yaml:"plannedLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags AwsBudget#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags_all AwsBudget#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#time_period_end AwsBudget#time_period_end}.
	// Experimental.
	TimePeriodEnd *string `field:"optional" json:"timePeriodEnd" yaml:"timePeriodEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#time_period_start AwsBudget#time_period_start}.
	// Experimental.
	TimePeriodStart *string `field:"optional" json:"timePeriodStart" yaml:"timePeriodStart"`
}

