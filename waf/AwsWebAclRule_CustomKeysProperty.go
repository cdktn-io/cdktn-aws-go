package waf


// Experimental.
type AwsWebAclRule_CustomKeysProperty struct {
	// asn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#asn AwsWebAclRule#asn}
	// Experimental.
	Asn interface{} `field:"optional" json:"asn" yaml:"asn"`
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#cookie AwsWebAclRule#cookie}
	// Experimental.
	Cookie interface{} `field:"optional" json:"cookie" yaml:"cookie"`
	// forwarded_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#forwarded_ip AwsWebAclRule#forwarded_ip}
	// Experimental.
	ForwardedIp interface{} `field:"optional" json:"forwardedIp" yaml:"forwardedIp"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#header AwsWebAclRule#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// http_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#http_method AwsWebAclRule#http_method}
	// Experimental.
	HttpMethod interface{} `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#ip AwsWebAclRule#ip}
	// Experimental.
	Ip interface{} `field:"optional" json:"ip" yaml:"ip"`
	// ja3_fingerprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#ja3_fingerprint AwsWebAclRule#ja3_fingerprint}
	// Experimental.
	Ja3Fingerprint interface{} `field:"optional" json:"ja3Fingerprint" yaml:"ja3Fingerprint"`
	// ja4_fingerprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#ja4_fingerprint AwsWebAclRule#ja4_fingerprint}
	// Experimental.
	Ja4Fingerprint interface{} `field:"optional" json:"ja4Fingerprint" yaml:"ja4Fingerprint"`
	// label_namespace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#label_namespace AwsWebAclRule#label_namespace}
	// Experimental.
	LabelNamespace interface{} `field:"optional" json:"labelNamespace" yaml:"labelNamespace"`
	// query_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#query_argument AwsWebAclRule#query_argument}
	// Experimental.
	QueryArgument interface{} `field:"optional" json:"queryArgument" yaml:"queryArgument"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#query_string AwsWebAclRule#query_string}
	// Experimental.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#uri_path AwsWebAclRule#uri_path}
	// Experimental.
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

