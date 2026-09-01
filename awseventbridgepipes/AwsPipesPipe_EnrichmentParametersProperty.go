package awseventbridgepipes


// Experimental.
type AwsPipesPipe_EnrichmentParametersProperty struct {
	// http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#http_parameters AwsPipesPipe#http_parameters}
	// Experimental.
	HttpParameters *AwsPipesPipe_EnrichmentParametersHttpParametersProperty `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#input_template AwsPipesPipe#input_template}.
	// Experimental.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
}

