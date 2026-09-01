package awsvpc


// Experimental.
type AwsVpcRouteServerPeer_BgpOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_route_server_peer#peer_asn AwsVpcRouteServerPeer#peer_asn}.
	// Experimental.
	PeerAsn *float64 `field:"required" json:"peerAsn" yaml:"peerAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_route_server_peer#peer_liveness_detection AwsVpcRouteServerPeer#peer_liveness_detection}.
	// Experimental.
	PeerLivenessDetection *string `field:"optional" json:"peerLivenessDetection" yaml:"peerLivenessDetection"`
}

