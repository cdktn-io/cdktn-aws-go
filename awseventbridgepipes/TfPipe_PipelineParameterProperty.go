package awseventbridgepipes


// Experimental.
type TfPipe_PipelineParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#name TfPipe#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#value TfPipe#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

