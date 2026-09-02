package awseventbridgepipes


// Experimental.
type TfPipe_EventbridgeEventBusParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#detail_type TfPipe#detail_type}.
	// Experimental.
	DetailType *string `field:"optional" json:"detailType" yaml:"detailType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#endpoint_id TfPipe#endpoint_id}.
	// Experimental.
	EndpointId *string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#resources TfPipe#resources}.
	// Experimental.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#source TfPipe#source}.
	// Experimental.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#time TfPipe#time}.
	// Experimental.
	Time *string `field:"optional" json:"time" yaml:"time"`
}

