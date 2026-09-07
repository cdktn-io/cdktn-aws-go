package networkfirewall


// Experimental.
type AwsTlsInspectionConfiguration_TlsInspectionConfigurationProperty struct {
	// server_certificate_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_tls_inspection_configuration#server_certificate_configuration AwsTlsInspectionConfiguration#server_certificate_configuration}
	// Experimental.
	ServerCertificateConfiguration interface{} `field:"optional" json:"serverCertificateConfiguration" yaml:"serverCertificateConfiguration"`
}

