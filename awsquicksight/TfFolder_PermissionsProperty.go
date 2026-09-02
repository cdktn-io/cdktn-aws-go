package awsquicksight


// Experimental.
type TfFolder_PermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_folder#actions TfFolder#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_folder#principal TfFolder#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

