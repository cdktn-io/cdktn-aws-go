package quicksight


// Experimental.
type AwsFolder_PermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_folder#actions AwsFolder#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_folder#principal AwsFolder#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

