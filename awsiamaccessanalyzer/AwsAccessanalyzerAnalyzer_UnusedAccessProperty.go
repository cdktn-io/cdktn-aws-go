package awsiamaccessanalyzer


// Experimental.
type AwsAccessanalyzerAnalyzer_UnusedAccessProperty struct {
	// analysis_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#analysis_rule AwsAccessanalyzerAnalyzer#analysis_rule}
	// Experimental.
	AnalysisRule *AwsAccessanalyzerAnalyzer_ConfigurationUnusedAccessAnalysisRuleProperty `field:"optional" json:"analysisRule" yaml:"analysisRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#unused_access_age AwsAccessanalyzerAnalyzer#unused_access_age}.
	// Experimental.
	UnusedAccessAge *float64 `field:"optional" json:"unusedAccessAge" yaml:"unusedAccessAge"`
}

