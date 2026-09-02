package awselb


// Experimental.
type TfAlbListenerRule_ConditionProperty struct {
	// host_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#host_header TfAlbListenerRule#host_header}
	// Experimental.
	HostHeader *TfAlbListenerRule_HostHeaderProperty `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#http_header TfAlbListenerRule#http_header}
	// Experimental.
	HttpHeader *TfAlbListenerRule_HttpHeaderProperty `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// http_request_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#http_request_method TfAlbListenerRule#http_request_method}
	// Experimental.
	HttpRequestMethod *TfAlbListenerRule_HttpRequestMethodProperty `field:"optional" json:"httpRequestMethod" yaml:"httpRequestMethod"`
	// path_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#path_pattern TfAlbListenerRule#path_pattern}
	// Experimental.
	PathPattern *TfAlbListenerRule_PathPatternProperty `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#query_string TfAlbListenerRule#query_string}
	// Experimental.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// source_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#source_ip TfAlbListenerRule#source_ip}
	// Experimental.
	SourceIp *TfAlbListenerRule_SourceIpProperty `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

