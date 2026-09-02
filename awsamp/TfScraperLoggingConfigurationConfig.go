package awsamp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfScraperLoggingConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper_logging_configuration#scraper_id TfScraperLoggingConfiguration#scraper_id}.
	// Experimental.
	ScraperId *string `field:"required" json:"scraperId" yaml:"scraperId"`
	// logging_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper_logging_configuration#logging_destination TfScraperLoggingConfiguration#logging_destination}
	// Experimental.
	LoggingDestination interface{} `field:"optional" json:"loggingDestination" yaml:"loggingDestination"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper_logging_configuration#region TfScraperLoggingConfiguration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper_logging_configuration#scraper_components TfScraperLoggingConfiguration#scraper_components}.
	// Experimental.
	ScraperComponents *[]*string `field:"optional" json:"scraperComponents" yaml:"scraperComponents"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper_logging_configuration#timeouts TfScraperLoggingConfiguration#timeouts}
	// Experimental.
	Timeouts *TfScraperLoggingConfiguration_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

