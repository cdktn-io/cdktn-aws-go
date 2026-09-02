package awsquicksight


// Experimental.
type TfRefreshSchedule_RefreshOnDayProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#day_of_month TfRefreshSchedule#day_of_month}.
	// Experimental.
	DayOfMonth *string `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#day_of_week TfRefreshSchedule#day_of_week}.
	// Experimental.
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
}

