package awswaf


// Experimental.
type TfWebAclRule_CookieProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#name TfWebAclRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#text_transformation TfWebAclRule#text_transformation}
	// Experimental.
	TextTransformation interface{} `field:"optional" json:"textTransformation" yaml:"textTransformation"`
}

