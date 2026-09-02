package awsverifiedpermissions


// Experimental.
type TfIdentitySource_IdentityTokenOnlyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#client_ids TfIdentitySource#client_ids}.
	// Experimental.
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#principal_id_claim TfIdentitySource#principal_id_claim}.
	// Experimental.
	PrincipalIdClaim *string `field:"optional" json:"principalIdClaim" yaml:"principalIdClaim"`
}

