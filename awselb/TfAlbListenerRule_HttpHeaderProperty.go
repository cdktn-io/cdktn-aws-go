package awselb


// Experimental.
type TfAlbListenerRule_HttpHeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#http_header_name TfAlbListenerRule#http_header_name}.
	// Experimental.
	HttpHeaderName *string `field:"required" json:"httpHeaderName" yaml:"httpHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#regex_values TfAlbListenerRule#regex_values}.
	// Experimental.
	RegexValues *[]*string `field:"optional" json:"regexValues" yaml:"regexValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#values TfAlbListenerRule#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

