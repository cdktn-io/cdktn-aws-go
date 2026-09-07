package s3tables


// Experimental.
type AwsTable_IcebergProperty struct {
	// A map of configuration properties for the Iceberg table, for example `write.distribution-mode` and `write.sort-order`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#properties AwsTable#properties}
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#schema AwsTable#schema}
	// Experimental.
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

