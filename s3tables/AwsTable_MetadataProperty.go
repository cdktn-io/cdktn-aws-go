package s3tables


// Experimental.
type AwsTable_MetadataProperty struct {
	// iceberg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#iceberg AwsTable#iceberg}
	// Experimental.
	Iceberg interface{} `field:"optional" json:"iceberg" yaml:"iceberg"`
}

