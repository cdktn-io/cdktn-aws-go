package glue


// Experimental.
type AwsCatalog_CreateDatabaseDefaultPermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#permissions AwsCatalog#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#principal AwsCatalog#principal}
	// Experimental.
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
}

