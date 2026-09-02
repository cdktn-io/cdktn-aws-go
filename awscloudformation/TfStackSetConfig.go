package awscloudformation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStackSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#name TfStackSet#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#administration_role_arn TfStackSet#administration_role_arn}.
	// Experimental.
	AdministrationRoleArn *string `field:"optional" json:"administrationRoleArn" yaml:"administrationRoleArn"`
	// auto_deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#auto_deployment TfStackSet#auto_deployment}
	// Experimental.
	AutoDeployment *TfStackSet_AutoDeploymentProperty `field:"optional" json:"autoDeployment" yaml:"autoDeployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#call_as TfStackSet#call_as}.
	// Experimental.
	CallAs *string `field:"optional" json:"callAs" yaml:"callAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#capabilities TfStackSet#capabilities}.
	// Experimental.
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#description TfStackSet#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#execution_role_name TfStackSet#execution_role_name}.
	// Experimental.
	ExecutionRoleName *string `field:"optional" json:"executionRoleName" yaml:"executionRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#id TfStackSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// managed_execution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#managed_execution TfStackSet#managed_execution}
	// Experimental.
	ManagedExecution *TfStackSet_ManagedExecutionProperty `field:"optional" json:"managedExecution" yaml:"managedExecution"`
	// operation_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#operation_preferences TfStackSet#operation_preferences}
	// Experimental.
	OperationPreferences *TfStackSet_OperationPreferencesProperty `field:"optional" json:"operationPreferences" yaml:"operationPreferences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#parameters TfStackSet#parameters}.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#permission_model TfStackSet#permission_model}.
	// Experimental.
	PermissionModel *string `field:"optional" json:"permissionModel" yaml:"permissionModel"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#region TfStackSet#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#tags TfStackSet#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#tags_all TfStackSet#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#template_body TfStackSet#template_body}.
	// Experimental.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#template_url TfStackSet#template_url}.
	// Experimental.
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#timeouts TfStackSet#timeouts}
	// Experimental.
	Timeouts *TfStackSet_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

