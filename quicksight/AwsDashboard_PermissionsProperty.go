package quicksight


// Experimental.
type AwsDashboard_PermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#actions AwsDashboard#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#principal AwsDashboard#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

