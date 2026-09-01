package awss3tables


// Experimental.
type AwsS3TablesTable_IcebergProperty struct {
	// A map of configuration properties for the Iceberg table, for example `write.distribution-mode` and `write.sort-order`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#properties AwsS3TablesTable#properties}
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#schema AwsS3TablesTable#schema}
	// Experimental.
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

