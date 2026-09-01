package awsquicksight


// Experimental.
type AwsQuicksightDataSource_PermissionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#actions AwsQuicksightDataSource#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#principal AwsQuicksightDataSource#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

