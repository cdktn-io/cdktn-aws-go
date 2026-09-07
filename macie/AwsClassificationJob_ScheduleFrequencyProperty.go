package macie


// Experimental.
type AwsClassificationJob_ScheduleFrequencyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#daily_schedule AwsClassificationJob#daily_schedule}.
	// Experimental.
	DailySchedule interface{} `field:"optional" json:"dailySchedule" yaml:"dailySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#monthly_schedule AwsClassificationJob#monthly_schedule}.
	// Experimental.
	MonthlySchedule *float64 `field:"optional" json:"monthlySchedule" yaml:"monthlySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/macie2_classification_job#weekly_schedule AwsClassificationJob#weekly_schedule}.
	// Experimental.
	WeeklySchedule *string `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

