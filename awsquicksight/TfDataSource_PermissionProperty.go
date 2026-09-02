package awsquicksight


// Experimental.
type TfDataSource_PermissionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#actions TfDataSource#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#principal TfDataSource#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

