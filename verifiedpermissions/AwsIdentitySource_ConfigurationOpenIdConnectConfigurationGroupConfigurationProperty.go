package verifiedpermissions


// Experimental.
type AwsIdentitySource_ConfigurationOpenIdConnectConfigurationGroupConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#group_claim AwsIdentitySource#group_claim}.
	// Experimental.
	GroupClaim *string `field:"required" json:"groupClaim" yaml:"groupClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#group_entity_type AwsIdentitySource#group_entity_type}.
	// Experimental.
	GroupEntityType *string `field:"required" json:"groupEntityType" yaml:"groupEntityType"`
}

