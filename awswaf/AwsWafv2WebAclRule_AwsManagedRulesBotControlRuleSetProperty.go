package awswaf


// Experimental.
type AwsWafv2WebAclRule_AwsManagedRulesBotControlRuleSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#inspection_level AwsWafv2WebAclRule#inspection_level}.
	// Experimental.
	InspectionLevel *string `field:"required" json:"inspectionLevel" yaml:"inspectionLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#enable_machine_learning AwsWafv2WebAclRule#enable_machine_learning}.
	// Experimental.
	EnableMachineLearning interface{} `field:"optional" json:"enableMachineLearning" yaml:"enableMachineLearning"`
}

