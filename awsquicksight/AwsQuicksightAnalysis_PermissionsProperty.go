package awsquicksight


// Experimental.
type AwsQuicksightAnalysis_PermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#actions AwsQuicksightAnalysis#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#principal AwsQuicksightAnalysis#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

