package msk


// Experimental.
type AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationProperty struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#sasl AwsCluster#sasl}
	// Experimental.
	Sasl *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslProperty `field:"optional" json:"sasl" yaml:"sasl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#tls AwsCluster#tls}.
	// Experimental.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

