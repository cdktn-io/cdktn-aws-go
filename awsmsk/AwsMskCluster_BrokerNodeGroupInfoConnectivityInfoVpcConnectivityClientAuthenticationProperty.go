package awsmsk


// Experimental.
type AwsMskCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationProperty struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#sasl AwsMskCluster#sasl}
	// Experimental.
	Sasl *AwsMskCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslProperty `field:"optional" json:"sasl" yaml:"sasl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#tls AwsMskCluster#tls}.
	// Experimental.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

