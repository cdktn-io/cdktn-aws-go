package fsx


// Experimental.
type AwsFileCache_NfsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#version AwsFileCache#version}.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#dns_ips AwsFileCache#dns_ips}.
	// Experimental.
	DnsIps *[]*string `field:"optional" json:"dnsIps" yaml:"dnsIps"`
}

