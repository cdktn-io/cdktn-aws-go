package finspace


// Experimental.
type AwsKxEnvironment_CustomDnsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#custom_dns_server_ip AwsKxEnvironment#custom_dns_server_ip}.
	// Experimental.
	CustomDnsServerIp *string `field:"required" json:"customDnsServerIp" yaml:"customDnsServerIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#custom_dns_server_name AwsKxEnvironment#custom_dns_server_name}.
	// Experimental.
	CustomDnsServerName *string `field:"required" json:"customDnsServerName" yaml:"customDnsServerName"`
}

