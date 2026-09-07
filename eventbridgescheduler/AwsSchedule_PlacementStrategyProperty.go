package eventbridgescheduler


// Experimental.
type AwsSchedule_PlacementStrategyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#type AwsSchedule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#field AwsSchedule#field}.
	// Experimental.
	Field *string `field:"optional" json:"field" yaml:"field"`
}

