package awsverifiedpermissions


// Experimental.
type AwsVerifiedpermissionsIdentitySource_AccessTokenOnlyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#audiences AwsVerifiedpermissionsIdentitySource#audiences}.
	// Experimental.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#principal_id_claim AwsVerifiedpermissionsIdentitySource#principal_id_claim}.
	// Experimental.
	PrincipalIdClaim *string `field:"optional" json:"principalIdClaim" yaml:"principalIdClaim"`
}

