package awsarcregionswitch

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPlanConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#execution_role TfPlan#execution_role}.
	// Experimental.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#name TfPlan#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#recovery_approach TfPlan#recovery_approach}.
	// Experimental.
	RecoveryApproach *string `field:"required" json:"recoveryApproach" yaml:"recoveryApproach"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#regions TfPlan#regions}.
	// Experimental.
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// associated_alarms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#associated_alarms TfPlan#associated_alarms}
	// Experimental.
	AssociatedAlarms interface{} `field:"optional" json:"associatedAlarms" yaml:"associatedAlarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#description TfPlan#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#primary_region TfPlan#primary_region}.
	// Experimental.
	PrimaryRegion *string `field:"optional" json:"primaryRegion" yaml:"primaryRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#recovery_time_objective_minutes TfPlan#recovery_time_objective_minutes}.
	// Experimental.
	RecoveryTimeObjectiveMinutes *float64 `field:"optional" json:"recoveryTimeObjectiveMinutes" yaml:"recoveryTimeObjectiveMinutes"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#region TfPlan#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// report_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#report_configuration TfPlan#report_configuration}
	// Experimental.
	ReportConfiguration interface{} `field:"optional" json:"reportConfiguration" yaml:"reportConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#tags TfPlan#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeouts TfPlan#timeouts}
	// Experimental.
	Timeouts *TfPlan_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// triggers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#triggers TfPlan#triggers}
	// Experimental.
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
	// workflow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#workflow TfPlan#workflow}
	// Experimental.
	Workflow interface{} `field:"optional" json:"workflow" yaml:"workflow"`
}

