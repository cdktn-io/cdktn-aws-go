package awssigner


// Experimental.
type AwsSignerSigningProfile_SigningMaterialProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/signer_signing_profile#certificate_arn AwsSignerSigningProfile#certificate_arn}.
	// Experimental.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
}

