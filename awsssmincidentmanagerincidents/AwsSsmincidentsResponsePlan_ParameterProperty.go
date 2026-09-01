package awsssmincidentmanagerincidents


// Experimental.
type AwsSsmincidentsResponsePlan_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#name AwsSsmincidentsResponsePlan#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#values AwsSsmincidentsResponsePlan#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

