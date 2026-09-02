package awsquicksight


// Experimental.
type TfAnalysis_PermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#actions TfAnalysis#actions}.
	// Experimental.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#principal TfAnalysis#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

