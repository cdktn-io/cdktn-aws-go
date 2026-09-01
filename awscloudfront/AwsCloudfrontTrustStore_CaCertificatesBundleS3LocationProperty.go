package awscloudfront


// Experimental.
type AwsCloudfrontTrustStore_CaCertificatesBundleS3LocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_trust_store#bucket AwsCloudfrontTrustStore#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_trust_store#key AwsCloudfrontTrustStore#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_trust_store#region AwsCloudfrontTrustStore#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_trust_store#version AwsCloudfrontTrustStore#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

