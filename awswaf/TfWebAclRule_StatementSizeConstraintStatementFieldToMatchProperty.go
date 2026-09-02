package awswaf


// Experimental.
type TfWebAclRule_StatementSizeConstraintStatementFieldToMatchProperty struct {
	// all_query_arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#all_query_arguments TfWebAclRule#all_query_arguments}
	// Experimental.
	AllQueryArguments interface{} `field:"optional" json:"allQueryArguments" yaml:"allQueryArguments"`
	// body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#body TfWebAclRule#body}
	// Experimental.
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#cookies TfWebAclRule#cookies}
	// Experimental.
	Cookies interface{} `field:"optional" json:"cookies" yaml:"cookies"`
	// header_order block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#header_order TfWebAclRule#header_order}
	// Experimental.
	HeaderOrder interface{} `field:"optional" json:"headerOrder" yaml:"headerOrder"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#headers TfWebAclRule#headers}
	// Experimental.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// ja3_fingerprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#ja3_fingerprint TfWebAclRule#ja3_fingerprint}
	// Experimental.
	Ja3Fingerprint interface{} `field:"optional" json:"ja3Fingerprint" yaml:"ja3Fingerprint"`
	// ja4_fingerprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#ja4_fingerprint TfWebAclRule#ja4_fingerprint}
	// Experimental.
	Ja4Fingerprint interface{} `field:"optional" json:"ja4Fingerprint" yaml:"ja4Fingerprint"`
	// json_body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#json_body TfWebAclRule#json_body}
	// Experimental.
	JsonBody interface{} `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#method TfWebAclRule#method}
	// Experimental.
	Method interface{} `field:"optional" json:"method" yaml:"method"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#query_string TfWebAclRule#query_string}
	// Experimental.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#single_header TfWebAclRule#single_header}
	// Experimental.
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// single_query_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#single_query_argument TfWebAclRule#single_query_argument}
	// Experimental.
	SingleQueryArgument interface{} `field:"optional" json:"singleQueryArgument" yaml:"singleQueryArgument"`
	// uri_fragment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#uri_fragment TfWebAclRule#uri_fragment}
	// Experimental.
	UriFragment interface{} `field:"optional" json:"uriFragment" yaml:"uriFragment"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#uri_path TfWebAclRule#uri_path}
	// Experimental.
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

