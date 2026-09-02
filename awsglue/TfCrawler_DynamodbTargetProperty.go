package awsglue


// Experimental.
type TfCrawler_DynamodbTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#path TfCrawler#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#scan_all TfCrawler#scan_all}.
	// Experimental.
	ScanAll interface{} `field:"optional" json:"scanAll" yaml:"scanAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#scan_rate TfCrawler#scan_rate}.
	// Experimental.
	ScanRate *float64 `field:"optional" json:"scanRate" yaml:"scanRate"`
}

