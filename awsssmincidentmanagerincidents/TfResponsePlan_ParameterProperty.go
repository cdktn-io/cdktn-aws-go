package awsssmincidentmanagerincidents


// Experimental.
type TfResponsePlan_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#name TfResponsePlan#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#values TfResponsePlan#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

