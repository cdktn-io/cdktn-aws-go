package awselb


// Experimental.
type AwsLbListenerRule_ConditionProperty struct {
	// host_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#host_header AwsLbListenerRule#host_header}
	// Experimental.
	HostHeader *AwsLbListenerRule_HostHeaderProperty `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#http_header AwsLbListenerRule#http_header}
	// Experimental.
	HttpHeader *AwsLbListenerRule_HttpHeaderProperty `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// http_request_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#http_request_method AwsLbListenerRule#http_request_method}
	// Experimental.
	HttpRequestMethod *AwsLbListenerRule_HttpRequestMethodProperty `field:"optional" json:"httpRequestMethod" yaml:"httpRequestMethod"`
	// path_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#path_pattern AwsLbListenerRule#path_pattern}
	// Experimental.
	PathPattern *AwsLbListenerRule_PathPatternProperty `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#query_string AwsLbListenerRule#query_string}
	// Experimental.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// source_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#source_ip AwsLbListenerRule#source_ip}
	// Experimental.
	SourceIp *AwsLbListenerRule_SourceIpProperty `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

