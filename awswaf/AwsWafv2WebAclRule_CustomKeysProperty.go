package awswaf


// Experimental.
type AwsWafv2WebAclRule_CustomKeysProperty struct {
	// asn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#asn AwsWafv2WebAclRule#asn}
	// Experimental.
	Asn interface{} `field:"optional" json:"asn" yaml:"asn"`
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#cookie AwsWafv2WebAclRule#cookie}
	// Experimental.
	Cookie interface{} `field:"optional" json:"cookie" yaml:"cookie"`
	// forwarded_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#forwarded_ip AwsWafv2WebAclRule#forwarded_ip}
	// Experimental.
	ForwardedIp interface{} `field:"optional" json:"forwardedIp" yaml:"forwardedIp"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#header AwsWafv2WebAclRule#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// http_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#http_method AwsWafv2WebAclRule#http_method}
	// Experimental.
	HttpMethod interface{} `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#ip AwsWafv2WebAclRule#ip}
	// Experimental.
	Ip interface{} `field:"optional" json:"ip" yaml:"ip"`
	// ja3_fingerprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#ja3_fingerprint AwsWafv2WebAclRule#ja3_fingerprint}
	// Experimental.
	Ja3Fingerprint interface{} `field:"optional" json:"ja3Fingerprint" yaml:"ja3Fingerprint"`
	// ja4_fingerprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#ja4_fingerprint AwsWafv2WebAclRule#ja4_fingerprint}
	// Experimental.
	Ja4Fingerprint interface{} `field:"optional" json:"ja4Fingerprint" yaml:"ja4Fingerprint"`
	// label_namespace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#label_namespace AwsWafv2WebAclRule#label_namespace}
	// Experimental.
	LabelNamespace interface{} `field:"optional" json:"labelNamespace" yaml:"labelNamespace"`
	// query_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#query_argument AwsWafv2WebAclRule#query_argument}
	// Experimental.
	QueryArgument interface{} `field:"optional" json:"queryArgument" yaml:"queryArgument"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#query_string AwsWafv2WebAclRule#query_string}
	// Experimental.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#uri_path AwsWafv2WebAclRule#uri_path}
	// Experimental.
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

