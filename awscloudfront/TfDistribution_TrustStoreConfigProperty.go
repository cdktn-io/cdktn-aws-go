package awscloudfront


// Experimental.
type TfDistribution_TrustStoreConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#trust_store_id TfDistribution#trust_store_id}.
	// Experimental.
	TrustStoreId *string `field:"required" json:"trustStoreId" yaml:"trustStoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#advertise_trust_store_ca_names TfDistribution#advertise_trust_store_ca_names}.
	// Experimental.
	AdvertiseTrustStoreCaNames interface{} `field:"optional" json:"advertiseTrustStoreCaNames" yaml:"advertiseTrustStoreCaNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#ignore_certificate_expiry TfDistribution#ignore_certificate_expiry}.
	// Experimental.
	IgnoreCertificateExpiry interface{} `field:"optional" json:"ignoreCertificateExpiry" yaml:"ignoreCertificateExpiry"`
}

