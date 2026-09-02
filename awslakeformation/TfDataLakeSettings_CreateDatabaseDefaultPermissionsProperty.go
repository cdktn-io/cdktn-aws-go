package awslakeformation


// Experimental.
type TfDataLakeSettings_CreateDatabaseDefaultPermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#permissions TfDataLakeSettings#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#principal TfDataLakeSettings#principal}.
	// Experimental.
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
}

