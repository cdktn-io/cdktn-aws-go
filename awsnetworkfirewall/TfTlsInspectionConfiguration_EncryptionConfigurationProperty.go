package awsnetworkfirewall


// Experimental.
type TfTlsInspectionConfiguration_EncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#key_id TfTlsInspectionConfiguration#key_id}.
	// Experimental.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#type TfTlsInspectionConfiguration#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

