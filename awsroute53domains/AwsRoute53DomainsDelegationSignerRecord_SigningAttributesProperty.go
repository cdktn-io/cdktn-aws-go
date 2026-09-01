package awsroute53domains


// Experimental.
type AwsRoute53DomainsDelegationSignerRecord_SigningAttributesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_delegation_signer_record#algorithm AwsRoute53DomainsDelegationSignerRecord#algorithm}.
	// Experimental.
	Algorithm *float64 `field:"required" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_delegation_signer_record#flags AwsRoute53DomainsDelegationSignerRecord#flags}.
	// Experimental.
	Flags *float64 `field:"required" json:"flags" yaml:"flags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_delegation_signer_record#public_key AwsRoute53DomainsDelegationSignerRecord#public_key}.
	// Experimental.
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}

