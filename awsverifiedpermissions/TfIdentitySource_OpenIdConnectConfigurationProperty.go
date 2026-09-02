package awsverifiedpermissions


// Experimental.
type TfIdentitySource_OpenIdConnectConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#issuer TfIdentitySource#issuer}.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#entity_id_prefix TfIdentitySource#entity_id_prefix}.
	// Experimental.
	EntityIdPrefix *string `field:"optional" json:"entityIdPrefix" yaml:"entityIdPrefix"`
	// group_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#group_configuration TfIdentitySource#group_configuration}
	// Experimental.
	GroupConfiguration interface{} `field:"optional" json:"groupConfiguration" yaml:"groupConfiguration"`
	// token_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#token_selection TfIdentitySource#token_selection}
	// Experimental.
	TokenSelection interface{} `field:"optional" json:"tokenSelection" yaml:"tokenSelection"`
}

