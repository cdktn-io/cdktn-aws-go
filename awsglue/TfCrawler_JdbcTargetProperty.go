package awsglue


// Experimental.
type TfCrawler_JdbcTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#connection_name TfCrawler#connection_name}.
	// Experimental.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#path TfCrawler#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#enable_additional_metadata TfCrawler#enable_additional_metadata}.
	// Experimental.
	EnableAdditionalMetadata *[]*string `field:"optional" json:"enableAdditionalMetadata" yaml:"enableAdditionalMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#exclusions TfCrawler#exclusions}.
	// Experimental.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
}

