package awsdatasync


// Experimental.
type TfTask_ScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#schedule_expression TfTask#schedule_expression}.
	// Experimental.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#status TfTask#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

