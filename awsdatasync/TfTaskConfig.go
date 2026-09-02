package awsdatasync

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTaskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#destination_location_arn TfTask#destination_location_arn}.
	// Experimental.
	DestinationLocationArn *string `field:"required" json:"destinationLocationArn" yaml:"destinationLocationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#source_location_arn TfTask#source_location_arn}.
	// Experimental.
	SourceLocationArn *string `field:"required" json:"sourceLocationArn" yaml:"sourceLocationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#cloudwatch_log_group_arn TfTask#cloudwatch_log_group_arn}.
	// Experimental.
	CloudwatchLogGroupArn *string `field:"optional" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// excludes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#excludes TfTask#excludes}
	// Experimental.
	Excludes *TfTask_ExcludesProperty `field:"optional" json:"excludes" yaml:"excludes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#id TfTask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// includes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#includes TfTask#includes}
	// Experimental.
	Includes *TfTask_IncludesProperty `field:"optional" json:"includes" yaml:"includes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#name TfTask#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#options TfTask#options}
	// Experimental.
	Options *TfTask_OptionsProperty `field:"optional" json:"options" yaml:"options"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#region TfTask#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#schedule TfTask#schedule}
	// Experimental.
	Schedule *TfTask_ScheduleProperty `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#tags TfTask#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#tags_all TfTask#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#task_mode TfTask#task_mode}.
	// Experimental.
	TaskMode *string `field:"optional" json:"taskMode" yaml:"taskMode"`
	// task_report_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#task_report_config TfTask#task_report_config}
	// Experimental.
	TaskReportConfig *TfTask_TaskReportConfigProperty `field:"optional" json:"taskReportConfig" yaml:"taskReportConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#timeouts TfTask#timeouts}
	// Experimental.
	Timeouts *TfTask_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

