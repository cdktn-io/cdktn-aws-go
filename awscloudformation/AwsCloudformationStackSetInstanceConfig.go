package awscloudformation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudformationStackSetInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#stack_set_name AwsCloudformationStackSetInstance#stack_set_name}.
	// Experimental.
	StackSetName *string `field:"required" json:"stackSetName" yaml:"stackSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#account_id AwsCloudformationStackSetInstance#account_id}.
	// Experimental.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#call_as AwsCloudformationStackSetInstance#call_as}.
	// Experimental.
	CallAs *string `field:"optional" json:"callAs" yaml:"callAs"`
	// deployment_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#deployment_targets AwsCloudformationStackSetInstance#deployment_targets}
	// Experimental.
	DeploymentTargets *AwsCloudformationStackSetInstance_DeploymentTargetsProperty `field:"optional" json:"deploymentTargets" yaml:"deploymentTargets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#id AwsCloudformationStackSetInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// operation_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#operation_preferences AwsCloudformationStackSetInstance#operation_preferences}
	// Experimental.
	OperationPreferences *AwsCloudformationStackSetInstance_OperationPreferencesProperty `field:"optional" json:"operationPreferences" yaml:"operationPreferences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#parameter_overrides AwsCloudformationStackSetInstance#parameter_overrides}.
	// Experimental.
	ParameterOverrides *map[string]*string `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#region AwsCloudformationStackSetInstance#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#retain_stack AwsCloudformationStackSetInstance#retain_stack}.
	// Experimental.
	RetainStack interface{} `field:"optional" json:"retainStack" yaml:"retainStack"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#stack_set_instance_region AwsCloudformationStackSetInstance#stack_set_instance_region}.
	// Experimental.
	StackSetInstanceRegion *string `field:"optional" json:"stackSetInstanceRegion" yaml:"stackSetInstanceRegion"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set_instance#timeouts AwsCloudformationStackSetInstance#timeouts}
	// Experimental.
	Timeouts *AwsCloudformationStackSetInstance_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

