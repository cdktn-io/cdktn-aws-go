package awsglue


// Experimental.
type TfCatalog_CreateTableDefaultPermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#permissions TfCatalog#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#principal TfCatalog#principal}
	// Experimental.
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
}

