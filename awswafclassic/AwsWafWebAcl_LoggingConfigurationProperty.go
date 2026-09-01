package awswafclassic


// Experimental.
type AwsWafWebAcl_LoggingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_web_acl#log_destination AwsWafWebAcl#log_destination}.
	// Experimental.
	LogDestination *string `field:"required" json:"logDestination" yaml:"logDestination"`
	// redacted_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_web_acl#redacted_fields AwsWafWebAcl#redacted_fields}
	// Experimental.
	RedactedFields *AwsWafWebAcl_RedactedFieldsProperty `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

