package awsglue


// Experimental.
type TfCrawler_IcebergTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#maximum_traversal_depth TfCrawler#maximum_traversal_depth}.
	// Experimental.
	MaximumTraversalDepth *float64 `field:"required" json:"maximumTraversalDepth" yaml:"maximumTraversalDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#paths TfCrawler#paths}.
	// Experimental.
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#connection_name TfCrawler#connection_name}.
	// Experimental.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#exclusions TfCrawler#exclusions}.
	// Experimental.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
}

