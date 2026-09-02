package awselb


// Experimental.
type TfListener_MutualAuthenticationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#mode TfListener#mode}.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#advertise_trust_store_ca_names TfListener#advertise_trust_store_ca_names}.
	// Experimental.
	AdvertiseTrustStoreCaNames *string `field:"optional" json:"advertiseTrustStoreCaNames" yaml:"advertiseTrustStoreCaNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#ignore_client_certificate_expiry TfListener#ignore_client_certificate_expiry}.
	// Experimental.
	IgnoreClientCertificateExpiry interface{} `field:"optional" json:"ignoreClientCertificateExpiry" yaml:"ignoreClientCertificateExpiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#trust_store_arn TfListener#trust_store_arn}.
	// Experimental.
	TrustStoreArn *string `field:"optional" json:"trustStoreArn" yaml:"trustStoreArn"`
}

