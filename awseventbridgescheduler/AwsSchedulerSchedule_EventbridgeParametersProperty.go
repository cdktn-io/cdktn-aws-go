package awseventbridgescheduler


// Experimental.
type AwsSchedulerSchedule_EventbridgeParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#detail_type AwsSchedulerSchedule#detail_type}.
	// Experimental.
	DetailType *string `field:"required" json:"detailType" yaml:"detailType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#source AwsSchedulerSchedule#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
}

