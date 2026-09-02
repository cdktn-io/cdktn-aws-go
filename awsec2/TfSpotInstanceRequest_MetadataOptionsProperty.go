package awsec2


// Experimental.
type TfSpotInstanceRequest_MetadataOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#http_endpoint TfSpotInstanceRequest#http_endpoint}.
	// Experimental.
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#http_protocol_ipv6 TfSpotInstanceRequest#http_protocol_ipv6}.
	// Experimental.
	HttpProtocolIpv6 *string `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#http_put_response_hop_limit TfSpotInstanceRequest#http_put_response_hop_limit}.
	// Experimental.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#http_tokens TfSpotInstanceRequest#http_tokens}.
	// Experimental.
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#instance_metadata_tags TfSpotInstanceRequest#instance_metadata_tags}.
	// Experimental.
	InstanceMetadataTags *string `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
}

