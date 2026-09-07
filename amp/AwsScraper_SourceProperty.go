package amp


// Experimental.
type AwsScraper_SourceProperty struct {
	// eks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#eks AwsScraper#eks}
	// Experimental.
	Eks interface{} `field:"optional" json:"eks" yaml:"eks"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#vpc AwsScraper#vpc}
	// Experimental.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

