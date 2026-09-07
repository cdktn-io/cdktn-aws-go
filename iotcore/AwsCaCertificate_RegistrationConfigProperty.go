package iotcore


// Experimental.
type AwsCaCertificate_RegistrationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_ca_certificate#role_arn AwsCaCertificate#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_ca_certificate#template_body AwsCaCertificate#template_body}.
	// Experimental.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_ca_certificate#template_name AwsCaCertificate#template_name}.
	// Experimental.
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}

