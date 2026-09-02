package awsapigateway


// Experimental.
type TfRestApi_EndpointConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_rest_api#types TfRestApi#types}.
	// Experimental.
	Types *[]*string `field:"required" json:"types" yaml:"types"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_rest_api#ip_address_type TfRestApi#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_rest_api#vpc_endpoint_ids TfRestApi#vpc_endpoint_ids}.
	// Experimental.
	VpcEndpointIds *[]*string `field:"optional" json:"vpcEndpointIds" yaml:"vpcEndpointIds"`
}

