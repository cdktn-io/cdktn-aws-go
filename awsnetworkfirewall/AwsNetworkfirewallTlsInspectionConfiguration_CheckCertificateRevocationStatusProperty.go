package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallTlsInspectionConfiguration_CheckCertificateRevocationStatusProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#revoked_status_action AwsNetworkfirewallTlsInspectionConfiguration#revoked_status_action}.
	// Experimental.
	RevokedStatusAction *string `field:"optional" json:"revokedStatusAction" yaml:"revokedStatusAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#unknown_status_action AwsNetworkfirewallTlsInspectionConfiguration#unknown_status_action}.
	// Experimental.
	UnknownStatusAction *string `field:"optional" json:"unknownStatusAction" yaml:"unknownStatusAction"`
}

