package quicksight


// Experimental.
type AwsDashboard_DateTimeParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#name AwsDashboard#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#values AwsDashboard#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

