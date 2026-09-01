package awssigner


// Experimental.
type AwsSignerSigningProfile_SignatureValidityPeriodProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/signer_signing_profile#type AwsSignerSigningProfile#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/signer_signing_profile#value AwsSignerSigningProfile#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

