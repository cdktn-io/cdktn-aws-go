package verifiedpermissions


// Experimental.
type AwsIdentitySource_OpenIdConnectConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#issuer AwsIdentitySource#issuer}.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#entity_id_prefix AwsIdentitySource#entity_id_prefix}.
	// Experimental.
	EntityIdPrefix *string `field:"optional" json:"entityIdPrefix" yaml:"entityIdPrefix"`
	// group_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#group_configuration AwsIdentitySource#group_configuration}
	// Experimental.
	GroupConfiguration interface{} `field:"optional" json:"groupConfiguration" yaml:"groupConfiguration"`
	// token_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#token_selection AwsIdentitySource#token_selection}
	// Experimental.
	TokenSelection interface{} `field:"optional" json:"tokenSelection" yaml:"tokenSelection"`
}

