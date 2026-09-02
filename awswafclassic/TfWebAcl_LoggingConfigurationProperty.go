package awswafclassic


// Experimental.
type TfWebAcl_LoggingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_web_acl#log_destination TfWebAcl#log_destination}.
	// Experimental.
	LogDestination *string `field:"required" json:"logDestination" yaml:"logDestination"`
	// redacted_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_web_acl#redacted_fields TfWebAcl#redacted_fields}
	// Experimental.
	RedactedFields *TfWebAcl_RedactedFieldsProperty `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

