package awswaf


// Experimental.
type TfWebAclRule_AwsManagedRulesBotControlRuleSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#inspection_level TfWebAclRule#inspection_level}.
	// Experimental.
	InspectionLevel *string `field:"required" json:"inspectionLevel" yaml:"inspectionLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#enable_machine_learning TfWebAclRule#enable_machine_learning}.
	// Experimental.
	EnableMachineLearning interface{} `field:"optional" json:"enableMachineLearning" yaml:"enableMachineLearning"`
}

