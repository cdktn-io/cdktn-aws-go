package awskendra


// Experimental.
type AwsKendraExperience_ContentSourceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_experience#data_source_ids AwsKendraExperience#data_source_ids}.
	// Experimental.
	DataSourceIds *[]*string `field:"optional" json:"dataSourceIds" yaml:"dataSourceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_experience#direct_put_content AwsKendraExperience#direct_put_content}.
	// Experimental.
	DirectPutContent interface{} `field:"optional" json:"directPutContent" yaml:"directPutContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_experience#faq_ids AwsKendraExperience#faq_ids}.
	// Experimental.
	FaqIds *[]*string `field:"optional" json:"faqIds" yaml:"faqIds"`
}

