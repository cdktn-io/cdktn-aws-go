package awsmsk


// Experimental.
type AwsMskCluster_ClientAuthenticationProperty struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#sasl AwsMskCluster#sasl}
	// Experimental.
	Sasl *AwsMskCluster_ClientAuthenticationSaslProperty `field:"optional" json:"sasl" yaml:"sasl"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#tls AwsMskCluster#tls}
	// Experimental.
	Tls *AwsMskCluster_TlsProperty `field:"optional" json:"tls" yaml:"tls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#unauthenticated AwsMskCluster#unauthenticated}.
	// Experimental.
	Unauthenticated interface{} `field:"optional" json:"unauthenticated" yaml:"unauthenticated"`
}

