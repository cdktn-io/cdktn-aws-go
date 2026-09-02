package awsdms


// Experimental.
type TfEndpoint_RedisSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#auth_type TfEndpoint#auth_type}.
	// Experimental.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#port TfEndpoint#port}.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#server_name TfEndpoint#server_name}.
	// Experimental.
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#auth_password TfEndpoint#auth_password}.
	// Experimental.
	AuthPassword *string `field:"optional" json:"authPassword" yaml:"authPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#auth_user_name TfEndpoint#auth_user_name}.
	// Experimental.
	AuthUserName *string `field:"optional" json:"authUserName" yaml:"authUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#ssl_ca_certificate_arn TfEndpoint#ssl_ca_certificate_arn}.
	// Experimental.
	SslCaCertificateArn *string `field:"optional" json:"sslCaCertificateArn" yaml:"sslCaCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#ssl_security_protocol TfEndpoint#ssl_security_protocol}.
	// Experimental.
	SslSecurityProtocol *string `field:"optional" json:"sslSecurityProtocol" yaml:"sslSecurityProtocol"`
}

