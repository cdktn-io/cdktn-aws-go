package awsec2


// Experimental.
type TfSpotInstanceRequest_PrivateDnsNameOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#enable_resource_name_dns_aaaa_record TfSpotInstanceRequest#enable_resource_name_dns_aaaa_record}.
	// Experimental.
	EnableResourceNameDnsAaaaRecord interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecord" yaml:"enableResourceNameDnsAaaaRecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#enable_resource_name_dns_a_record TfSpotInstanceRequest#enable_resource_name_dns_a_record}.
	// Experimental.
	EnableResourceNameDnsARecord interface{} `field:"optional" json:"enableResourceNameDnsARecord" yaml:"enableResourceNameDnsARecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#hostname_type TfSpotInstanceRequest#hostname_type}.
	// Experimental.
	HostnameType *string `field:"optional" json:"hostnameType" yaml:"hostnameType"`
}

