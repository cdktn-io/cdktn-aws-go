package quicksight


// Experimental.
type AwsRefreshSchedule_ScheduleFrequencyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#interval AwsRefreshSchedule#interval}.
	// Experimental.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// refresh_on_day block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#refresh_on_day AwsRefreshSchedule#refresh_on_day}
	// Experimental.
	RefreshOnDay interface{} `field:"optional" json:"refreshOnDay" yaml:"refreshOnDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#time_of_the_day AwsRefreshSchedule#time_of_the_day}.
	// Experimental.
	TimeOfTheDay *string `field:"optional" json:"timeOfTheDay" yaml:"timeOfTheDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#timezone AwsRefreshSchedule#timezone}.
	// Experimental.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

