package waf


// Experimental.
type AwsWebAclRuleGroupAssociation_AddressFieldsProperty struct {
	// Identifiers of the address fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#identifiers AwsWebAclRuleGroupAssociation#identifiers}
	// Experimental.
	Identifiers *[]*string `field:"required" json:"identifiers" yaml:"identifiers"`
}

