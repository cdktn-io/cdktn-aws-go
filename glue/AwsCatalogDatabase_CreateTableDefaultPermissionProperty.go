package glue


// Experimental.
type AwsCatalogDatabase_CreateTableDefaultPermissionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_database#permissions AwsCatalogDatabase#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_database#principal AwsCatalogDatabase#principal}
	// Experimental.
	Principal *AwsCatalogDatabase_PrincipalProperty `field:"optional" json:"principal" yaml:"principal"`
}

