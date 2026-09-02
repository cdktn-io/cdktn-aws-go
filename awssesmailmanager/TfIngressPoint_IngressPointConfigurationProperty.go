package awssesmailmanager


// Experimental.
type TfIngressPoint_IngressPointConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#secret_arn TfIngressPoint#secret_arn}.
	// Experimental.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#smtp_password_wo TfIngressPoint#smtp_password_wo}.
	// Experimental.
	SmtpPasswordWo *string `field:"optional" json:"smtpPasswordWo" yaml:"smtpPasswordWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#smtp_password_wo_version TfIngressPoint#smtp_password_wo_version}.
	// Experimental.
	SmtpPasswordWoVersion *float64 `field:"optional" json:"smtpPasswordWoVersion" yaml:"smtpPasswordWoVersion"`
	// tls_auth_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#tls_auth_configuration TfIngressPoint#tls_auth_configuration}
	// Experimental.
	TlsAuthConfiguration interface{} `field:"optional" json:"tlsAuthConfiguration" yaml:"tlsAuthConfiguration"`
}

