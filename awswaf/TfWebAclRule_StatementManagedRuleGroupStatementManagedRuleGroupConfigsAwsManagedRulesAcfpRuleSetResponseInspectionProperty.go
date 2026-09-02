package awswaf


// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetResponseInspectionProperty struct {
	// body_contains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#body_contains TfWebAclRule#body_contains}
	// Experimental.
	BodyContains interface{} `field:"optional" json:"bodyContains" yaml:"bodyContains"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#header TfWebAclRule#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#json TfWebAclRule#json}
	// Experimental.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// status_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#status_code TfWebAclRule#status_code}
	// Experimental.
	StatusCode interface{} `field:"optional" json:"statusCode" yaml:"statusCode"`
}

