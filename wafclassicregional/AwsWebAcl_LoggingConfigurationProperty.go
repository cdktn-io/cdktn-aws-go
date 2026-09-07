package wafclassicregional


// Experimental.
type AwsWebAcl_LoggingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#log_destination AwsWebAcl#log_destination}.
	// Experimental.
	LogDestination *string `field:"required" json:"logDestination" yaml:"logDestination"`
	// redacted_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_web_acl#redacted_fields AwsWebAcl#redacted_fields}
	// Experimental.
	RedactedFields *AwsWebAcl_RedactedFieldsProperty `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

