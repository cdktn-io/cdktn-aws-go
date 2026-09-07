package eventbridgescheduler


// Experimental.
type AwsSchedule_EventbridgeParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#detail_type AwsSchedule#detail_type}.
	// Experimental.
	DetailType *string `field:"required" json:"detailType" yaml:"detailType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#source AwsSchedule#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
}

