package awssecurityhub

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAutomationRuleV2Config struct {
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
	// A description of the automation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#description TfAutomationRuleV2#description}
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the automation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#rule_name TfAutomationRuleV2#rule_name}
	// Experimental.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// The priority of the rule (lower values = higher priority).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#rule_order TfAutomationRuleV2#rule_order}
	// Experimental.
	RuleOrder *float64 `field:"required" json:"ruleOrder" yaml:"ruleOrder"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#action TfAutomationRuleV2#action}
	// Experimental.
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#criteria TfAutomationRuleV2#criteria}
	// Experimental.
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#region TfAutomationRuleV2#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The status of the rule: ENABLED or DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#rule_status TfAutomationRuleV2#rule_status}
	// Experimental.
	RuleStatus *string `field:"optional" json:"ruleStatus" yaml:"ruleStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#tags TfAutomationRuleV2#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

