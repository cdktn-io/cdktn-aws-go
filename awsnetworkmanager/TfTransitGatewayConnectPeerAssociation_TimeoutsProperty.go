package awsnetworkmanager


// Experimental.
type TfTransitGatewayConnectPeerAssociation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_transit_gateway_connect_peer_association#create TfTransitGatewayConnectPeerAssociation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_transit_gateway_connect_peer_association#delete TfTransitGatewayConnectPeerAssociation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

