package awseventbridge


// Experimental.
type TfConnection_InvocationHttpParametersProperty struct {
	// body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#body TfConnection#body}
	// Experimental.
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#header TfConnection#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#query_string TfConnection#query_string}
	// Experimental.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
}

