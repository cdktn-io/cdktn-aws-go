package iamaccessanalyzer


// Experimental.
type AwsAnalyzer_ConfigurationProperty struct {
	// internal_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#internal_access AwsAnalyzer#internal_access}
	// Experimental.
	InternalAccess *AwsAnalyzer_InternalAccessProperty `field:"optional" json:"internalAccess" yaml:"internalAccess"`
	// unused_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#unused_access AwsAnalyzer#unused_access}
	// Experimental.
	UnusedAccess *AwsAnalyzer_UnusedAccessProperty `field:"optional" json:"unusedAccess" yaml:"unusedAccess"`
}

