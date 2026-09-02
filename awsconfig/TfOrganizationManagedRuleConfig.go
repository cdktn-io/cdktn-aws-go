package awsconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOrganizationManagedRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#name TfOrganizationManagedRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#rule_identifier TfOrganizationManagedRule#rule_identifier}.
	// Experimental.
	RuleIdentifier *string `field:"required" json:"ruleIdentifier" yaml:"ruleIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#description TfOrganizationManagedRule#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#excluded_accounts TfOrganizationManagedRule#excluded_accounts}.
	// Experimental.
	ExcludedAccounts *[]*string `field:"optional" json:"excludedAccounts" yaml:"excludedAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#id TfOrganizationManagedRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#input_parameters TfOrganizationManagedRule#input_parameters}.
	// Experimental.
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#maximum_execution_frequency TfOrganizationManagedRule#maximum_execution_frequency}.
	// Experimental.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#region TfOrganizationManagedRule#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#resource_id_scope TfOrganizationManagedRule#resource_id_scope}.
	// Experimental.
	ResourceIdScope *string `field:"optional" json:"resourceIdScope" yaml:"resourceIdScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#resource_types_scope TfOrganizationManagedRule#resource_types_scope}.
	// Experimental.
	ResourceTypesScope *[]*string `field:"optional" json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#tag_key_scope TfOrganizationManagedRule#tag_key_scope}.
	// Experimental.
	TagKeyScope *string `field:"optional" json:"tagKeyScope" yaml:"tagKeyScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#tag_value_scope TfOrganizationManagedRule#tag_value_scope}.
	// Experimental.
	TagValueScope *string `field:"optional" json:"tagValueScope" yaml:"tagValueScope"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#timeouts TfOrganizationManagedRule#timeouts}
	// Experimental.
	Timeouts *TfOrganizationManagedRule_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

