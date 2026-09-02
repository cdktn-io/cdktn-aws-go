package awsssm

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMaintenanceWindowTaskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#task_arn TfMaintenanceWindowTask#task_arn}.
	// Experimental.
	TaskArn *string `field:"required" json:"taskArn" yaml:"taskArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#task_type TfMaintenanceWindowTask#task_type}.
	// Experimental.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#window_id TfMaintenanceWindowTask#window_id}.
	// Experimental.
	WindowId *string `field:"required" json:"windowId" yaml:"windowId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#cutoff_behavior TfMaintenanceWindowTask#cutoff_behavior}.
	// Experimental.
	CutoffBehavior *string `field:"optional" json:"cutoffBehavior" yaml:"cutoffBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#description TfMaintenanceWindowTask#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#id TfMaintenanceWindowTask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#max_concurrency TfMaintenanceWindowTask#max_concurrency}.
	// Experimental.
	MaxConcurrency *string `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#max_errors TfMaintenanceWindowTask#max_errors}.
	// Experimental.
	MaxErrors *string `field:"optional" json:"maxErrors" yaml:"maxErrors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#name TfMaintenanceWindowTask#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#priority TfMaintenanceWindowTask#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#region TfMaintenanceWindowTask#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#service_role_arn TfMaintenanceWindowTask#service_role_arn}.
	// Experimental.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#targets TfMaintenanceWindowTask#targets}
	// Experimental.
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// task_invocation_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#task_invocation_parameters TfMaintenanceWindowTask#task_invocation_parameters}
	// Experimental.
	TaskInvocationParameters *TfMaintenanceWindowTask_TaskInvocationParametersProperty `field:"optional" json:"taskInvocationParameters" yaml:"taskInvocationParameters"`
}

