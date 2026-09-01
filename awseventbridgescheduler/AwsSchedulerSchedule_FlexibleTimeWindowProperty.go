package awseventbridgescheduler


// Experimental.
type AwsSchedulerSchedule_FlexibleTimeWindowProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#mode AwsSchedulerSchedule#mode}.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#maximum_window_in_minutes AwsSchedulerSchedule#maximum_window_in_minutes}.
	// Experimental.
	MaximumWindowInMinutes *float64 `field:"optional" json:"maximumWindowInMinutes" yaml:"maximumWindowInMinutes"`
}

