package config

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOrganizationCustomPolicyRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#name AwsOrganizationCustomPolicyRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#policy_runtime AwsOrganizationCustomPolicyRule#policy_runtime}.
	// Experimental.
	PolicyRuntime *string `field:"required" json:"policyRuntime" yaml:"policyRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#policy_text AwsOrganizationCustomPolicyRule#policy_text}.
	// Experimental.
	PolicyText *string `field:"required" json:"policyText" yaml:"policyText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#trigger_types AwsOrganizationCustomPolicyRule#trigger_types}.
	// Experimental.
	TriggerTypes *[]*string `field:"required" json:"triggerTypes" yaml:"triggerTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#debug_log_delivery_accounts AwsOrganizationCustomPolicyRule#debug_log_delivery_accounts}.
	// Experimental.
	DebugLogDeliveryAccounts *[]*string `field:"optional" json:"debugLogDeliveryAccounts" yaml:"debugLogDeliveryAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#description AwsOrganizationCustomPolicyRule#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#excluded_accounts AwsOrganizationCustomPolicyRule#excluded_accounts}.
	// Experimental.
	ExcludedAccounts *[]*string `field:"optional" json:"excludedAccounts" yaml:"excludedAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#id AwsOrganizationCustomPolicyRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#input_parameters AwsOrganizationCustomPolicyRule#input_parameters}.
	// Experimental.
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#maximum_execution_frequency AwsOrganizationCustomPolicyRule#maximum_execution_frequency}.
	// Experimental.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#region AwsOrganizationCustomPolicyRule#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#resource_id_scope AwsOrganizationCustomPolicyRule#resource_id_scope}.
	// Experimental.
	ResourceIdScope *string `field:"optional" json:"resourceIdScope" yaml:"resourceIdScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#resource_types_scope AwsOrganizationCustomPolicyRule#resource_types_scope}.
	// Experimental.
	ResourceTypesScope *[]*string `field:"optional" json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#tag_key_scope AwsOrganizationCustomPolicyRule#tag_key_scope}.
	// Experimental.
	TagKeyScope *string `field:"optional" json:"tagKeyScope" yaml:"tagKeyScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#tag_value_scope AwsOrganizationCustomPolicyRule#tag_value_scope}.
	// Experimental.
	TagValueScope *string `field:"optional" json:"tagValueScope" yaml:"tagValueScope"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#timeouts AwsOrganizationCustomPolicyRule#timeouts}
	// Experimental.
	Timeouts *AwsOrganizationCustomPolicyRule_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

