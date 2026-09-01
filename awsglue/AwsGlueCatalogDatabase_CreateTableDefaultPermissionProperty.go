package awsglue


// Experimental.
type AwsGlueCatalogDatabase_CreateTableDefaultPermissionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_database#permissions AwsGlueCatalogDatabase#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_database#principal AwsGlueCatalogDatabase#principal}
	// Experimental.
	Principal *AwsGlueCatalogDatabase_PrincipalProperty `field:"optional" json:"principal" yaml:"principal"`
}

