package awscognitoidp


// Experimental.
type TfUserPool_UserPoolAddOnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#advanced_security_mode TfUserPool#advanced_security_mode}.
	// Experimental.
	AdvancedSecurityMode *string `field:"required" json:"advancedSecurityMode" yaml:"advancedSecurityMode"`
	// advanced_security_additional_flows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#advanced_security_additional_flows TfUserPool#advanced_security_additional_flows}
	// Experimental.
	AdvancedSecurityAdditionalFlows *TfUserPool_AdvancedSecurityAdditionalFlowsProperty `field:"optional" json:"advancedSecurityAdditionalFlows" yaml:"advancedSecurityAdditionalFlows"`
}

