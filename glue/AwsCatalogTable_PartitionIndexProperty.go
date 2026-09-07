package glue


// Experimental.
type AwsCatalogTable_PartitionIndexProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#index_name AwsCatalogTable#index_name}.
	// Experimental.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#keys AwsCatalogTable#keys}.
	// Experimental.
	Keys *[]*string `field:"required" json:"keys" yaml:"keys"`
}

