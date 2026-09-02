package awsroute53recoverycontrolconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSafetyRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#control_panel_arn TfSafetyRule#control_panel_arn}.
	// Experimental.
	ControlPanelArn *string `field:"required" json:"controlPanelArn" yaml:"controlPanelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#name TfSafetyRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// rule_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#rule_config TfSafetyRule#rule_config}
	// Experimental.
	RuleConfig *TfSafetyRule_RuleConfigProperty `field:"required" json:"ruleConfig" yaml:"ruleConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#wait_period_ms TfSafetyRule#wait_period_ms}.
	// Experimental.
	WaitPeriodMs *float64 `field:"required" json:"waitPeriodMs" yaml:"waitPeriodMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#asserted_controls TfSafetyRule#asserted_controls}.
	// Experimental.
	AssertedControls *[]*string `field:"optional" json:"assertedControls" yaml:"assertedControls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#gating_controls TfSafetyRule#gating_controls}.
	// Experimental.
	GatingControls *[]*string `field:"optional" json:"gatingControls" yaml:"gatingControls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#id TfSafetyRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#tags TfSafetyRule#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#tags_all TfSafetyRule#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53recoverycontrolconfig_safety_rule#target_controls TfSafetyRule#target_controls}.
	// Experimental.
	TargetControls *[]*string `field:"optional" json:"targetControls" yaml:"targetControls"`
}

