package awsiamaccessanalyzer


// Experimental.
type TfAnalyzer_InclusionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#account_ids TfAnalyzer#account_ids}.
	// Experimental.
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#resource_arns TfAnalyzer#resource_arns}.
	// Experimental.
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_analyzer#resource_types TfAnalyzer#resource_types}.
	// Experimental.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

