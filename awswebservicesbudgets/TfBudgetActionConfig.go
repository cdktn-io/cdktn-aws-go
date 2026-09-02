package awswebservicesbudgets

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBudgetActionConfig struct {
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
	// action_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#action_threshold TfBudgetAction#action_threshold}
	// Experimental.
	ActionThreshold *TfBudgetAction_ActionThresholdProperty `field:"required" json:"actionThreshold" yaml:"actionThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#action_type TfBudgetAction#action_type}.
	// Experimental.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#approval_model TfBudgetAction#approval_model}.
	// Experimental.
	ApprovalModel *string `field:"required" json:"approvalModel" yaml:"approvalModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#budget_name TfBudgetAction#budget_name}.
	// Experimental.
	BudgetName *string `field:"required" json:"budgetName" yaml:"budgetName"`
	// definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#definition TfBudgetAction#definition}
	// Experimental.
	Definition *TfBudgetAction_DefinitionProperty `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#execution_role_arn TfBudgetAction#execution_role_arn}.
	// Experimental.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#notification_type TfBudgetAction#notification_type}.
	// Experimental.
	NotificationType *string `field:"required" json:"notificationType" yaml:"notificationType"`
	// subscriber block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#subscriber TfBudgetAction#subscriber}
	// Experimental.
	Subscriber interface{} `field:"required" json:"subscriber" yaml:"subscriber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#account_id TfBudgetAction#account_id}.
	// Experimental.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#id TfBudgetAction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#tags TfBudgetAction#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#tags_all TfBudgetAction#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#timeouts TfBudgetAction#timeouts}
	// Experimental.
	Timeouts *TfBudgetAction_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

