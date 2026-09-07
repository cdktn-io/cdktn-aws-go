package amp


// Experimental.
type AwsScraper_DestinationProperty struct {
	// amp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#amp AwsScraper#amp}
	// Experimental.
	Amp interface{} `field:"optional" json:"amp" yaml:"amp"`
	// cloudwatch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#cloudwatch AwsScraper#cloudwatch}
	// Experimental.
	Cloudwatch interface{} `field:"optional" json:"cloudwatch" yaml:"cloudwatch"`
}

