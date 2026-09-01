package awseventbridgescheduler


// Experimental.
type AwsSchedulerSchedule_PlacementStrategyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#type AwsSchedulerSchedule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#field AwsSchedulerSchedule#field}.
	// Experimental.
	Field *string `field:"optional" json:"field" yaml:"field"`
}

