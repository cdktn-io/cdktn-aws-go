package awsappflow


// Experimental.
type TfFlow_ScheduledProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#schedule_expression TfFlow#schedule_expression}.
	// Experimental.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#data_pull_mode TfFlow#data_pull_mode}.
	// Experimental.
	DataPullMode *string `field:"optional" json:"dataPullMode" yaml:"dataPullMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#first_execution_from TfFlow#first_execution_from}.
	// Experimental.
	FirstExecutionFrom *string `field:"optional" json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#schedule_end_time TfFlow#schedule_end_time}.
	// Experimental.
	ScheduleEndTime *string `field:"optional" json:"scheduleEndTime" yaml:"scheduleEndTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#schedule_offset TfFlow#schedule_offset}.
	// Experimental.
	ScheduleOffset *float64 `field:"optional" json:"scheduleOffset" yaml:"scheduleOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#schedule_start_time TfFlow#schedule_start_time}.
	// Experimental.
	ScheduleStartTime *string `field:"optional" json:"scheduleStartTime" yaml:"scheduleStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#timezone TfFlow#timezone}.
	// Experimental.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

