package awsmsk


// Experimental.
type AwsMskCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#iam AwsMskCluster#iam}.
	// Experimental.
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#scram AwsMskCluster#scram}.
	// Experimental.
	Scram interface{} `field:"optional" json:"scram" yaml:"scram"`
}

