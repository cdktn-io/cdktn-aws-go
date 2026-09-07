package lakeformation


// Experimental.
type AwsDataLakeSettings_CreateDatabaseDefaultPermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#permissions AwsDataLakeSettings#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#principal AwsDataLakeSettings#principal}.
	// Experimental.
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
}

