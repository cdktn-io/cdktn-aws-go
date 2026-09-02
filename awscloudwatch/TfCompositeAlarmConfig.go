package awscloudwatch

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCompositeAlarmConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#alarm_name TfCompositeAlarm#alarm_name}.
	// Experimental.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#alarm_rule TfCompositeAlarm#alarm_rule}.
	// Experimental.
	AlarmRule *string `field:"required" json:"alarmRule" yaml:"alarmRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#actions_enabled TfCompositeAlarm#actions_enabled}.
	// Experimental.
	ActionsEnabled interface{} `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// actions_suppressor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#actions_suppressor TfCompositeAlarm#actions_suppressor}
	// Experimental.
	ActionsSuppressor *TfCompositeAlarm_ActionsSuppressorProperty `field:"optional" json:"actionsSuppressor" yaml:"actionsSuppressor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#alarm_actions TfCompositeAlarm#alarm_actions}.
	// Experimental.
	AlarmActions *[]*string `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#alarm_description TfCompositeAlarm#alarm_description}.
	// Experimental.
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#id TfCompositeAlarm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#insufficient_data_actions TfCompositeAlarm#insufficient_data_actions}.
	// Experimental.
	InsufficientDataActions *[]*string `field:"optional" json:"insufficientDataActions" yaml:"insufficientDataActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#ok_actions TfCompositeAlarm#ok_actions}.
	// Experimental.
	OkActions *[]*string `field:"optional" json:"okActions" yaml:"okActions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#region TfCompositeAlarm#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#tags TfCompositeAlarm#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_composite_alarm#tags_all TfCompositeAlarm#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

