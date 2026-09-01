package awssecurityhub


// Experimental.
type AwsSecurityhubAutomationRule_FindingFieldsUpdateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#confidence AwsSecurityhubAutomationRule#confidence}.
	// Experimental.
	Confidence *float64 `field:"optional" json:"confidence" yaml:"confidence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#criticality AwsSecurityhubAutomationRule#criticality}.
	// Experimental.
	Criticality *float64 `field:"optional" json:"criticality" yaml:"criticality"`
	// note block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#note AwsSecurityhubAutomationRule#note}
	// Experimental.
	Note interface{} `field:"optional" json:"note" yaml:"note"`
	// related_findings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#related_findings AwsSecurityhubAutomationRule#related_findings}
	// Experimental.
	RelatedFindings interface{} `field:"optional" json:"relatedFindings" yaml:"relatedFindings"`
	// severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#severity AwsSecurityhubAutomationRule#severity}
	// Experimental.
	Severity interface{} `field:"optional" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#types AwsSecurityhubAutomationRule#types}.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#user_defined_fields AwsSecurityhubAutomationRule#user_defined_fields}.
	// Experimental.
	UserDefinedFields *map[string]*string `field:"optional" json:"userDefinedFields" yaml:"userDefinedFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#verification_state AwsSecurityhubAutomationRule#verification_state}.
	// Experimental.
	VerificationState *string `field:"optional" json:"verificationState" yaml:"verificationState"`
	// workflow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#workflow AwsSecurityhubAutomationRule#workflow}
	// Experimental.
	Workflow interface{} `field:"optional" json:"workflow" yaml:"workflow"`
}

