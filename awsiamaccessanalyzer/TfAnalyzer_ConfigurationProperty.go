package awsiamaccessanalyzer


// Experimental.
type TfAnalyzer_ConfigurationProperty struct {
	// internal_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#internal_access TfAnalyzer#internal_access}
	// Experimental.
	InternalAccess *TfAnalyzer_InternalAccessProperty `field:"optional" json:"internalAccess" yaml:"internalAccess"`
	// unused_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#unused_access TfAnalyzer#unused_access}
	// Experimental.
	UnusedAccess *TfAnalyzer_UnusedAccessProperty `field:"optional" json:"unusedAccess" yaml:"unusedAccess"`
}

