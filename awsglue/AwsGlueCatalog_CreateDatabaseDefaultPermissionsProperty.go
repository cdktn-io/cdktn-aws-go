package awsglue


// Experimental.
type AwsGlueCatalog_CreateDatabaseDefaultPermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#permissions AwsGlueCatalog#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#principal AwsGlueCatalog#principal}
	// Experimental.
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
}

