package awslakeformation


// Experimental.
type AwsLakeformationDataLakeSettings_CreateDatabaseDefaultPermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#permissions AwsLakeformationDataLakeSettings#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_lake_settings#principal AwsLakeformationDataLakeSettings#principal}.
	// Experimental.
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
}

