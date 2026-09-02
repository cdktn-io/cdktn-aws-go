package awsvpc


// Experimental.
type TfEc2NetworkInsightsAccessScope_ExcludePathsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#destination TfEc2NetworkInsightsAccessScope#destination}
	// Experimental.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#source TfEc2NetworkInsightsAccessScope#source}
	// Experimental.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// through_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#through_resources TfEc2NetworkInsightsAccessScope#through_resources}
	// Experimental.
	ThroughResources interface{} `field:"optional" json:"throughResources" yaml:"throughResources"`
}

