package awsvpnclient


// Experimental.
type TfEndpoint_TransitGatewayConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#availability_zone_ids TfEndpoint#availability_zone_ids}.
	// Experimental.
	AvailabilityZoneIds *[]*string `field:"optional" json:"availabilityZoneIds" yaml:"availabilityZoneIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#availability_zones TfEndpoint#availability_zones}.
	// Experimental.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#transit_gateway_id TfEndpoint#transit_gateway_id}.
	// Experimental.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
}

