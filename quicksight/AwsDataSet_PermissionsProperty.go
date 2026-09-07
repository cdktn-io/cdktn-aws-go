package quicksight


// Experimental.
type AwsDataSet_PermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#actions AwsDataSet#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#principal AwsDataSet#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

