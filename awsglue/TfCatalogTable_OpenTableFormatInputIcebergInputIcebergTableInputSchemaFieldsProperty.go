package awsglue


// Experimental.
type TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#id TfCatalogTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#name TfCatalogTable#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#required TfCatalogTable#required}.
	// Experimental.
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#type TfCatalogTable#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#doc TfCatalogTable#doc}.
	// Experimental.
	Doc *string `field:"optional" json:"doc" yaml:"doc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#initial_default TfCatalogTable#initial_default}.
	// Experimental.
	InitialDefault *string `field:"optional" json:"initialDefault" yaml:"initialDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#write_default TfCatalogTable#write_default}.
	// Experimental.
	WriteDefault *string `field:"optional" json:"writeDefault" yaml:"writeDefault"`
}

