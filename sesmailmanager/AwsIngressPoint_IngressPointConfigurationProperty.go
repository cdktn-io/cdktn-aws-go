package sesmailmanager


// Experimental.
type AwsIngressPoint_IngressPointConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#secret_arn AwsIngressPoint#secret_arn}.
	// Experimental.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#smtp_password_wo AwsIngressPoint#smtp_password_wo}.
	// Experimental.
	SmtpPasswordWo *string `field:"optional" json:"smtpPasswordWo" yaml:"smtpPasswordWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#smtp_password_wo_version AwsIngressPoint#smtp_password_wo_version}.
	// Experimental.
	SmtpPasswordWoVersion *float64 `field:"optional" json:"smtpPasswordWoVersion" yaml:"smtpPasswordWoVersion"`
	// tls_auth_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#tls_auth_configuration AwsIngressPoint#tls_auth_configuration}
	// Experimental.
	TlsAuthConfiguration interface{} `field:"optional" json:"tlsAuthConfiguration" yaml:"tlsAuthConfiguration"`
}

