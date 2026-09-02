package awssesmailmanager


// Experimental.
type TfRuleSet_RuleConditionStringExpressionEvaluateProperty struct {
	// analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#analysis TfRuleSet#analysis}
	// Experimental.
	Analysis interface{} `field:"optional" json:"analysis" yaml:"analysis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#attribute TfRuleSet#attribute}.
	// Experimental.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#client_certificate_attribute TfRuleSet#client_certificate_attribute}.
	// Experimental.
	ClientCertificateAttribute *string `field:"optional" json:"clientCertificateAttribute" yaml:"clientCertificateAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#mime_header_attribute TfRuleSet#mime_header_attribute}.
	// Experimental.
	MimeHeaderAttribute *string `field:"optional" json:"mimeHeaderAttribute" yaml:"mimeHeaderAttribute"`
}

