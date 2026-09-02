package awscloudfront


// Experimental.
type TfMultitenantDistribution_CustomOriginConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#http_port TfMultitenantDistribution#http_port}.
	// Experimental.
	HttpPort *float64 `field:"required" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#https_port TfMultitenantDistribution#https_port}.
	// Experimental.
	HttpsPort *float64 `field:"required" json:"httpsPort" yaml:"httpsPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_protocol_policy TfMultitenantDistribution#origin_protocol_policy}.
	// Experimental.
	OriginProtocolPolicy *string `field:"required" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_ssl_protocols TfMultitenantDistribution#origin_ssl_protocols}.
	// Experimental.
	OriginSslProtocols *[]*string `field:"required" json:"originSslProtocols" yaml:"originSslProtocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#ip_address_type TfMultitenantDistribution#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_keepalive_timeout TfMultitenantDistribution#origin_keepalive_timeout}.
	// Experimental.
	OriginKeepaliveTimeout *float64 `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// origin_mtls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_mtls_config TfMultitenantDistribution#origin_mtls_config}
	// Experimental.
	OriginMtlsConfig interface{} `field:"optional" json:"originMtlsConfig" yaml:"originMtlsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_read_timeout TfMultitenantDistribution#origin_read_timeout}.
	// Experimental.
	OriginReadTimeout *float64 `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
}

