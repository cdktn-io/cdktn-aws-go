package awscloudfront


// Experimental.
type TfVpcOrigin_VpcOriginEndpointConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_vpc_origin#arn TfVpcOrigin#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_vpc_origin#http_port TfVpcOrigin#http_port}.
	// Experimental.
	HttpPort *float64 `field:"required" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_vpc_origin#https_port TfVpcOrigin#https_port}.
	// Experimental.
	HttpsPort *float64 `field:"required" json:"httpsPort" yaml:"httpsPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_vpc_origin#name TfVpcOrigin#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_vpc_origin#origin_protocol_policy TfVpcOrigin#origin_protocol_policy}.
	// Experimental.
	OriginProtocolPolicy *string `field:"required" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// origin_ssl_protocols block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_vpc_origin#origin_ssl_protocols TfVpcOrigin#origin_ssl_protocols}
	// Experimental.
	OriginSslProtocols interface{} `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
}

