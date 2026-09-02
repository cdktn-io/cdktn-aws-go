package awsquicksight


// Experimental.
type TfTheme_PermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#actions TfTheme#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#principal TfTheme#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

