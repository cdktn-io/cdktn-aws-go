package awsmsk


// Experimental.
type TfCluster_ClientAuthenticationProperty struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#sasl TfCluster#sasl}
	// Experimental.
	Sasl *TfCluster_ClientAuthenticationSaslProperty `field:"optional" json:"sasl" yaml:"sasl"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#tls TfCluster#tls}
	// Experimental.
	Tls *TfCluster_TlsProperty `field:"optional" json:"tls" yaml:"tls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#unauthenticated TfCluster#unauthenticated}.
	// Experimental.
	Unauthenticated interface{} `field:"optional" json:"unauthenticated" yaml:"unauthenticated"`
}

