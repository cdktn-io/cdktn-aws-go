package awsglue


// Experimental.
type AwsGlueCatalogTable_PartitionIndexProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#index_name AwsGlueCatalogTable#index_name}.
	// Experimental.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#keys AwsGlueCatalogTable#keys}.
	// Experimental.
	Keys *[]*string `field:"required" json:"keys" yaml:"keys"`
}

