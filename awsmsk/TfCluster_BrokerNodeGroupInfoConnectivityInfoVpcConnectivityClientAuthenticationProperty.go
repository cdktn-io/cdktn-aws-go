package awsmsk


// Experimental.
type TfCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationProperty struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#sasl TfCluster#sasl}
	// Experimental.
	Sasl *TfCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslProperty `field:"optional" json:"sasl" yaml:"sasl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#tls TfCluster#tls}.
	// Experimental.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

