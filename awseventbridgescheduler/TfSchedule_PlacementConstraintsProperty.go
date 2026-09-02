package awseventbridgescheduler


// Experimental.
type TfSchedule_PlacementConstraintsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#type TfSchedule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#expression TfSchedule#expression}.
	// Experimental.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

