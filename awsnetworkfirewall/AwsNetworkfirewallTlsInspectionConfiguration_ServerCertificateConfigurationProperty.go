package awsnetworkfirewall


// Experimental.
type AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#certificate_authority_arn AwsNetworkfirewallTlsInspectionConfiguration#certificate_authority_arn}.
	// Experimental.
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// check_certificate_revocation_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#check_certificate_revocation_status AwsNetworkfirewallTlsInspectionConfiguration#check_certificate_revocation_status}
	// Experimental.
	CheckCertificateRevocationStatus interface{} `field:"optional" json:"checkCertificateRevocationStatus" yaml:"checkCertificateRevocationStatus"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#scope AwsNetworkfirewallTlsInspectionConfiguration#scope}
	// Experimental.
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
	// server_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#server_certificate AwsNetworkfirewallTlsInspectionConfiguration#server_certificate}
	// Experimental.
	ServerCertificate interface{} `field:"optional" json:"serverCertificate" yaml:"serverCertificate"`
}

