package awseventbridge


// Experimental.
type TfTarget_PipelineParameterListProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#name TfTarget#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#value TfTarget#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

