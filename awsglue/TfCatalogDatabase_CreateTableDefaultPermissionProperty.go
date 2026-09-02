package awsglue


// Experimental.
type TfCatalogDatabase_CreateTableDefaultPermissionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_database#permissions TfCatalogDatabase#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_database#principal TfCatalogDatabase#principal}
	// Experimental.
	Principal *TfCatalogDatabase_PrincipalProperty `field:"optional" json:"principal" yaml:"principal"`
}

