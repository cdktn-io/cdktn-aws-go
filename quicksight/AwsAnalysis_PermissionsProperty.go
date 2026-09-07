package quicksight


// Experimental.
type AwsAnalysis_PermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#actions AwsAnalysis#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#principal AwsAnalysis#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

