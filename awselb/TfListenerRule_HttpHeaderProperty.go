package awselb


// Experimental.
type TfListenerRule_HttpHeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#http_header_name TfListenerRule#http_header_name}.
	// Experimental.
	HttpHeaderName *string `field:"required" json:"httpHeaderName" yaml:"httpHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#regex_values TfListenerRule#regex_values}.
	// Experimental.
	RegexValues *[]*string `field:"optional" json:"regexValues" yaml:"regexValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#values TfListenerRule#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

