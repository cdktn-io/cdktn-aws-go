package eventbridgepipes


// Experimental.
type AwsPipe_EnrichmentParametersProperty struct {
	// http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#http_parameters AwsPipe#http_parameters}
	// Experimental.
	HttpParameters *AwsPipe_EnrichmentParametersHttpParametersProperty `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#input_template AwsPipe#input_template}.
	// Experimental.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
}

