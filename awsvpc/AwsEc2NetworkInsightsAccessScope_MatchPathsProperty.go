package awsvpc


// Experimental.
type AwsEc2NetworkInsightsAccessScope_MatchPathsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#destination AwsEc2NetworkInsightsAccessScope#destination}
	// Experimental.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#source AwsEc2NetworkInsightsAccessScope#source}
	// Experimental.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

