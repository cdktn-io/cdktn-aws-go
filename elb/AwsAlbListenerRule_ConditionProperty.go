package elb


// Experimental.
type AwsAlbListenerRule_ConditionProperty struct {
	// host_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#host_header AwsAlbListenerRule#host_header}
	// Experimental.
	HostHeader *AwsAlbListenerRule_HostHeaderProperty `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#http_header AwsAlbListenerRule#http_header}
	// Experimental.
	HttpHeader *AwsAlbListenerRule_HttpHeaderProperty `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// http_request_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#http_request_method AwsAlbListenerRule#http_request_method}
	// Experimental.
	HttpRequestMethod *AwsAlbListenerRule_HttpRequestMethodProperty `field:"optional" json:"httpRequestMethod" yaml:"httpRequestMethod"`
	// path_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#path_pattern AwsAlbListenerRule#path_pattern}
	// Experimental.
	PathPattern *AwsAlbListenerRule_PathPatternProperty `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#query_string AwsAlbListenerRule#query_string}
	// Experimental.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// source_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#source_ip AwsAlbListenerRule#source_ip}
	// Experimental.
	SourceIp *AwsAlbListenerRule_SourceIpProperty `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}

