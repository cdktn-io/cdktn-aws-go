package awsiamaccessanalyzer


// Experimental.
type AwsAccessanalyzerAnalyzer_ConfigurationProperty struct {
	// internal_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#internal_access AwsAccessanalyzerAnalyzer#internal_access}
	// Experimental.
	InternalAccess *AwsAccessanalyzerAnalyzer_InternalAccessProperty `field:"optional" json:"internalAccess" yaml:"internalAccess"`
	// unused_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#unused_access AwsAccessanalyzerAnalyzer#unused_access}
	// Experimental.
	UnusedAccess *AwsAccessanalyzerAnalyzer_UnusedAccessProperty `field:"optional" json:"unusedAccess" yaml:"unusedAccess"`
}

