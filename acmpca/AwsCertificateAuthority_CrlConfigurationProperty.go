package acmpca


// Experimental.
type AwsCertificateAuthority_CrlConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#custom_cname AwsCertificateAuthority#custom_cname}.
	// Experimental.
	CustomCname *string `field:"optional" json:"customCname" yaml:"customCname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#custom_path AwsCertificateAuthority#custom_path}.
	// Experimental.
	CustomPath *string `field:"optional" json:"customPath" yaml:"customPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#enabled AwsCertificateAuthority#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#expiration_in_days AwsCertificateAuthority#expiration_in_days}.
	// Experimental.
	ExpirationInDays *float64 `field:"optional" json:"expirationInDays" yaml:"expirationInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#s3_bucket_name AwsCertificateAuthority#s3_bucket_name}.
	// Experimental.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#s3_object_acl AwsCertificateAuthority#s3_object_acl}.
	// Experimental.
	S3ObjectAcl *string `field:"optional" json:"s3ObjectAcl" yaml:"s3ObjectAcl"`
}

