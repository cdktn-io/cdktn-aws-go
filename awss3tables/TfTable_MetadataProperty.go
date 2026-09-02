package awss3tables


// Experimental.
type TfTable_MetadataProperty struct {
	// iceberg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#iceberg TfTable#iceberg}
	// Experimental.
	Iceberg interface{} `field:"optional" json:"iceberg" yaml:"iceberg"`
}

