package awswaf


// Experimental.
type TfWebAclRuleGroupAssociation_AwsManagedRulesAcfpRuleSetProperty struct {
	// Path to the account creation endpoint on the protected website.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#creation_path TfWebAclRuleGroupAssociation#creation_path}
	// Experimental.
	CreationPath *string `field:"required" json:"creationPath" yaml:"creationPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#registration_page_path TfWebAclRuleGroupAssociation#registration_page_path}.
	// Experimental.
	RegistrationPagePath *string `field:"required" json:"registrationPagePath" yaml:"registrationPagePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#enable_regex_in_path TfWebAclRuleGroupAssociation#enable_regex_in_path}.
	// Experimental.
	EnableRegexInPath interface{} `field:"optional" json:"enableRegexInPath" yaml:"enableRegexInPath"`
	// request_inspection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#request_inspection TfWebAclRuleGroupAssociation#request_inspection}
	// Experimental.
	RequestInspection interface{} `field:"optional" json:"requestInspection" yaml:"requestInspection"`
	// response_inspection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#response_inspection TfWebAclRuleGroupAssociation#response_inspection}
	// Experimental.
	ResponseInspection interface{} `field:"optional" json:"responseInspection" yaml:"responseInspection"`
}

