package awsamp


// Experimental.
type AwsPrometheusScraper_DestinationProperty struct {
	// amp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#amp AwsPrometheusScraper#amp}
	// Experimental.
	Amp interface{} `field:"optional" json:"amp" yaml:"amp"`
	// cloudwatch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#cloudwatch AwsPrometheusScraper#cloudwatch}
	// Experimental.
	Cloudwatch interface{} `field:"optional" json:"cloudwatch" yaml:"cloudwatch"`
}

