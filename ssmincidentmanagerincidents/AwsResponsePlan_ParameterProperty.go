package ssmincidentmanagerincidents


// Experimental.
type AwsResponsePlan_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#name AwsResponsePlan#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmincidents_response_plan#values AwsResponsePlan#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

