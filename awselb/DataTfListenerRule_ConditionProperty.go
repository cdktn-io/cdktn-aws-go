package awselb


// Experimental.
type DataTfListenerRule_ConditionProperty struct {
	// host_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#host_header DataTfListenerRule#host_header}
	// Experimental.
	HostHeader interface{} `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#http_header DataTfListenerRule#http_header}
	// Experimental.
	HttpHeader interface{} `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// http_request_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#http_request_method DataTfListenerRule#http_request_method}
	// Experimental.
	HttpRequestMethod interface{} `field:"optional" json:"httpRequestMethod" yaml:"httpRequestMethod"`
	// path_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#path_pattern DataTfListenerRule#path_pattern}
	// Experimental.
	PathPattern interface{} `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#query_string DataTfListenerRule#query_string}
	// Experimental.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// source_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#source_ip DataTfListenerRule#source_ip}
	// Experimental.
	SourceIp interface{} `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

