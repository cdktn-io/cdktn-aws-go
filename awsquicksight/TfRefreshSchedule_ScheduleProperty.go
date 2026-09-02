package awsquicksight


// Experimental.
type TfRefreshSchedule_ScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#refresh_type TfRefreshSchedule#refresh_type}.
	// Experimental.
	RefreshType *string `field:"required" json:"refreshType" yaml:"refreshType"`
	// schedule_frequency block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#schedule_frequency TfRefreshSchedule#schedule_frequency}
	// Experimental.
	ScheduleFrequency interface{} `field:"optional" json:"scheduleFrequency" yaml:"scheduleFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#start_after_date_time TfRefreshSchedule#start_after_date_time}.
	// Experimental.
	StartAfterDateTime *string `field:"optional" json:"startAfterDateTime" yaml:"startAfterDateTime"`
}

