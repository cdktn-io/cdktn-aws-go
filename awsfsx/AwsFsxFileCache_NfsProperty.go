package awsfsx


// Experimental.
type AwsFsxFileCache_NfsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#version AwsFsxFileCache#version}.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#dns_ips AwsFsxFileCache#dns_ips}.
	// Experimental.
	DnsIps *[]*string `field:"optional" json:"dnsIps" yaml:"dnsIps"`
}

