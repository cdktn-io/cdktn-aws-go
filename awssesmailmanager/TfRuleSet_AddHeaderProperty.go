package awssesmailmanager


// Experimental.
type TfRuleSet_AddHeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#header_name TfRuleSet#header_name}.
	// Experimental.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#header_value TfRuleSet#header_value}.
	// Experimental.
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
}

