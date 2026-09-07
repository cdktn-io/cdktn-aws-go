package glue


// Experimental.
type AwsCatalogTable_PartitionSpecProperty struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#fields AwsCatalogTable#fields}
	// Experimental.
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#spec_id AwsCatalogTable#spec_id}.
	// Experimental.
	SpecId *float64 `field:"optional" json:"specId" yaml:"specId"`
}

