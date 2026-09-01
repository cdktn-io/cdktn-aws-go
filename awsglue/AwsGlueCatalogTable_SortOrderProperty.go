package awsglue


// Experimental.
type AwsGlueCatalogTable_SortOrderProperty struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#fields AwsGlueCatalogTable#fields}
	// Experimental.
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#order_id AwsGlueCatalogTable#order_id}.
	// Experimental.
	OrderId *float64 `field:"required" json:"orderId" yaml:"orderId"`
}

