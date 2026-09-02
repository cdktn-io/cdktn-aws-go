package awsmsk


// Experimental.
type TfCluster_ClientAuthenticationSaslProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#iam TfCluster#iam}.
	// Experimental.
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#scram TfCluster#scram}.
	// Experimental.
	Scram interface{} `field:"optional" json:"scram" yaml:"scram"`
}

