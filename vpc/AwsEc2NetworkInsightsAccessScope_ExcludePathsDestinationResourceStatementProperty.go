package vpc


// Experimental.
type AwsEc2NetworkInsightsAccessScope_ExcludePathsDestinationResourceStatementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#resources AwsEc2NetworkInsightsAccessScope#resources}.
	// Experimental.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#resource_types AwsEc2NetworkInsightsAccessScope#resource_types}.
	// Experimental.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

