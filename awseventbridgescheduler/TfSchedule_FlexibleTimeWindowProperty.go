package awseventbridgescheduler


// Experimental.
type TfSchedule_FlexibleTimeWindowProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#mode TfSchedule#mode}.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#maximum_window_in_minutes TfSchedule#maximum_window_in_minutes}.
	// Experimental.
	MaximumWindowInMinutes *float64 `field:"optional" json:"maximumWindowInMinutes" yaml:"maximumWindowInMinutes"`
}

