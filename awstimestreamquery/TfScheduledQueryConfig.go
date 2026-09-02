package awstimestreamquery

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfScheduledQueryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#execution_role_arn TfScheduledQuery#execution_role_arn}.
	// Experimental.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#name TfScheduledQuery#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#query_string TfScheduledQuery#query_string}.
	// Experimental.
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// error_report_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#error_report_configuration TfScheduledQuery#error_report_configuration}
	// Experimental.
	ErrorReportConfiguration interface{} `field:"optional" json:"errorReportConfiguration" yaml:"errorReportConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#kms_key_id TfScheduledQuery#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// last_run_summary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#last_run_summary TfScheduledQuery#last_run_summary}
	// Experimental.
	LastRunSummary interface{} `field:"optional" json:"lastRunSummary" yaml:"lastRunSummary"`
	// notification_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#notification_configuration TfScheduledQuery#notification_configuration}
	// Experimental.
	NotificationConfiguration interface{} `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// recently_failed_runs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#recently_failed_runs TfScheduledQuery#recently_failed_runs}
	// Experimental.
	RecentlyFailedRuns interface{} `field:"optional" json:"recentlyFailedRuns" yaml:"recentlyFailedRuns"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#region TfScheduledQuery#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// schedule_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#schedule_configuration TfScheduledQuery#schedule_configuration}
	// Experimental.
	ScheduleConfiguration interface{} `field:"optional" json:"scheduleConfiguration" yaml:"scheduleConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#tags TfScheduledQuery#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// target_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#target_configuration TfScheduledQuery#target_configuration}
	// Experimental.
	TargetConfiguration interface{} `field:"optional" json:"targetConfiguration" yaml:"targetConfiguration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#timeouts TfScheduledQuery#timeouts}
	// Experimental.
	Timeouts *TfScheduledQuery_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

