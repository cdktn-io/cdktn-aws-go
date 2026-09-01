package awslightsail


// Experimental.
type AwsLightsailDistribution_OriginProperty struct {
	// The name of the origin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#name AwsLightsailDistribution#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS Region name of the origin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#region_name AwsLightsailDistribution#region_name}
	// Experimental.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#protocol_policy AwsLightsailDistribution#protocol_policy}
	// Experimental.
	ProtocolPolicy *string `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
}

