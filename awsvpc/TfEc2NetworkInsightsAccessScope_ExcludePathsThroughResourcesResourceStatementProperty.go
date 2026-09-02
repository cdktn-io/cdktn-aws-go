package awsvpc


// Experimental.
type TfEc2NetworkInsightsAccessScope_ExcludePathsThroughResourcesResourceStatementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#resources TfEc2NetworkInsightsAccessScope#resources}.
	// Experimental.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#resource_types TfEc2NetworkInsightsAccessScope#resource_types}.
	// Experimental.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

