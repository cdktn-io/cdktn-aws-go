package awselb


// Experimental.
type TfListenerRule_ConditionProperty struct {
	// host_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#host_header TfListenerRule#host_header}
	// Experimental.
	HostHeader *TfListenerRule_HostHeaderProperty `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#http_header TfListenerRule#http_header}
	// Experimental.
	HttpHeader *TfListenerRule_HttpHeaderProperty `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// http_request_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#http_request_method TfListenerRule#http_request_method}
	// Experimental.
	HttpRequestMethod *TfListenerRule_HttpRequestMethodProperty `field:"optional" json:"httpRequestMethod" yaml:"httpRequestMethod"`
	// path_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#path_pattern TfListenerRule#path_pattern}
	// Experimental.
	PathPattern *TfListenerRule_PathPatternProperty `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#query_string TfListenerRule#query_string}
	// Experimental.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// source_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#source_ip TfListenerRule#source_ip}
	// Experimental.
	SourceIp *TfListenerRule_SourceIpProperty `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

