package awseventbridgescheduler


// Experimental.
type TfSchedule_PipelineParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#name TfSchedule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#value TfSchedule#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

