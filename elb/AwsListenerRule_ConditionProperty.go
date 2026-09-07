package elb


// Experimental.
type AwsListenerRule_ConditionProperty struct {
	// host_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#host_header AwsListenerRule#host_header}
	// Experimental.
	HostHeader *AwsListenerRule_HostHeaderProperty `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#http_header AwsListenerRule#http_header}
	// Experimental.
	HttpHeader *AwsListenerRule_HttpHeaderProperty `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// http_request_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#http_request_method AwsListenerRule#http_request_method}
	// Experimental.
	HttpRequestMethod *AwsListenerRule_HttpRequestMethodProperty `field:"optional" json:"httpRequestMethod" yaml:"httpRequestMethod"`
	// path_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#path_pattern AwsListenerRule#path_pattern}
	// Experimental.
	PathPattern *AwsListenerRule_PathPatternProperty `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#query_string AwsListenerRule#query_string}
	// Experimental.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// source_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#source_ip AwsListenerRule#source_ip}
	// Experimental.
	SourceIp *AwsListenerRule_SourceIpProperty `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

