package eventbridgescheduler


// Experimental.
type AwsSchedule_PipelineParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#name AwsSchedule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/scheduler_schedule#value AwsSchedule#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

