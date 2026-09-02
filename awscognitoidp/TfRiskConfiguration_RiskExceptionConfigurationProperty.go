package awscognitoidp


// Experimental.
type TfRiskConfiguration_RiskExceptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#blocked_ip_range_list TfRiskConfiguration#blocked_ip_range_list}.
	// Experimental.
	BlockedIpRangeList *[]*string `field:"optional" json:"blockedIpRangeList" yaml:"blockedIpRangeList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#skipped_ip_range_list TfRiskConfiguration#skipped_ip_range_list}.
	// Experimental.
	SkippedIpRangeList *[]*string `field:"optional" json:"skippedIpRangeList" yaml:"skippedIpRangeList"`
}

