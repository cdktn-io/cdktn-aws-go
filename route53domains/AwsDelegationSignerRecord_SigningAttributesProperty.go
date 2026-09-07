package route53domains


// Experimental.
type AwsDelegationSignerRecord_SigningAttributesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_delegation_signer_record#algorithm AwsDelegationSignerRecord#algorithm}.
	// Experimental.
	Algorithm *float64 `field:"required" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_delegation_signer_record#flags AwsDelegationSignerRecord#flags}.
	// Experimental.
	Flags *float64 `field:"required" json:"flags" yaml:"flags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_delegation_signer_record#public_key AwsDelegationSignerRecord#public_key}.
	// Experimental.
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}

