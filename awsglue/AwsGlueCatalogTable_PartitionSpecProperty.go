package awsglue


// Experimental.
type AwsGlueCatalogTable_PartitionSpecProperty struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#fields AwsGlueCatalogTable#fields}
	// Experimental.
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#spec_id AwsGlueCatalogTable#spec_id}.
	// Experimental.
	SpecId *float64 `field:"optional" json:"specId" yaml:"specId"`
}

