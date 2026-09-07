package iamaccessanalyzer


// Experimental.
type AwsAnalyzer_ExclusionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#account_ids AwsAnalyzer#account_ids}.
	// Experimental.
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#resource_tags AwsAnalyzer#resource_tags}.
	// Experimental.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

