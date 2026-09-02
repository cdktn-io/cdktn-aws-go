package awslightsail


// Experimental.
type TfDistribution_OriginProperty struct {
	// The name of the origin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#name TfDistribution#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS Region name of the origin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#region_name TfDistribution#region_name}
	// Experimental.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#protocol_policy TfDistribution#protocol_policy}
	// Experimental.
	ProtocolPolicy *string `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
}

