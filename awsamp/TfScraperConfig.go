package awsamp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfScraperConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#scrape_configuration TfScraper#scrape_configuration}.
	// Experimental.
	ScrapeConfiguration *string `field:"required" json:"scrapeConfiguration" yaml:"scrapeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#alias TfScraper#alias}.
	// Experimental.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#destination TfScraper#destination}
	// Experimental.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// exporter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#exporter TfScraper#exporter}
	// Experimental.
	Exporter interface{} `field:"optional" json:"exporter" yaml:"exporter"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#region TfScraper#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// role_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#role_configuration TfScraper#role_configuration}
	// Experimental.
	RoleConfiguration interface{} `field:"optional" json:"roleConfiguration" yaml:"roleConfiguration"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#source TfScraper#source}
	// Experimental.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#tags TfScraper#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#timeouts TfScraper#timeouts}
	// Experimental.
	Timeouts *TfScraper_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

