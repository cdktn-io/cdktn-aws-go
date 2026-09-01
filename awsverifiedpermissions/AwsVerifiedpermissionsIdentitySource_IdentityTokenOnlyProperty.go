package awsverifiedpermissions


// Experimental.
type AwsVerifiedpermissionsIdentitySource_IdentityTokenOnlyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#client_ids AwsVerifiedpermissionsIdentitySource#client_ids}.
	// Experimental.
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#principal_id_claim AwsVerifiedpermissionsIdentitySource#principal_id_claim}.
	// Experimental.
	PrincipalIdClaim *string `field:"optional" json:"principalIdClaim" yaml:"principalIdClaim"`
}

