package awsdatasync


// Experimental.
type AwsDatasyncTask_ScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#schedule_expression AwsDatasyncTask#schedule_expression}.
	// Experimental.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#status AwsDatasyncTask#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

