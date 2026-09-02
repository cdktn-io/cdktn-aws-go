package awssigner


// Experimental.
type TfSigningProfile_SignatureValidityPeriodProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/signer_signing_profile#type TfSigningProfile#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/signer_signing_profile#value TfSigningProfile#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

