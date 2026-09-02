package awsamp


// Experimental.
type TfScraper_DestinationProperty struct {
	// amp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#amp TfScraper#amp}
	// Experimental.
	Amp interface{} `field:"optional" json:"amp" yaml:"amp"`
	// cloudwatch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#cloudwatch TfScraper#cloudwatch}
	// Experimental.
	Cloudwatch interface{} `field:"optional" json:"cloudwatch" yaml:"cloudwatch"`
}

