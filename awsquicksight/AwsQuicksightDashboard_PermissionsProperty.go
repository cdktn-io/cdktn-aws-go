package awsquicksight


// Experimental.
type AwsQuicksightDashboard_PermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#actions AwsQuicksightDashboard#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#principal AwsQuicksightDashboard#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

