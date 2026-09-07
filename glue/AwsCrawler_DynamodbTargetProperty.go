package glue


// Experimental.
type AwsCrawler_DynamodbTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#path AwsCrawler#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#scan_all AwsCrawler#scan_all}.
	// Experimental.
	ScanAll interface{} `field:"optional" json:"scanAll" yaml:"scanAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#scan_rate AwsCrawler#scan_rate}.
	// Experimental.
	ScanRate *float64 `field:"optional" json:"scanRate" yaml:"scanRate"`
}

