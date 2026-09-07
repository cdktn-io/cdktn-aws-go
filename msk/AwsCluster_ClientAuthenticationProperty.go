package msk


// Experimental.
type AwsCluster_ClientAuthenticationProperty struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#sasl AwsCluster#sasl}
	// Experimental.
	Sasl *AwsCluster_ClientAuthenticationSaslProperty `field:"optional" json:"sasl" yaml:"sasl"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#tls AwsCluster#tls}
	// Experimental.
	Tls *AwsCluster_TlsProperty `field:"optional" json:"tls" yaml:"tls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#unauthenticated AwsCluster#unauthenticated}.
	// Experimental.
	Unauthenticated interface{} `field:"optional" json:"unauthenticated" yaml:"unauthenticated"`
}

