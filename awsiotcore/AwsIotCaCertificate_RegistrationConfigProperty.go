package awsiotcore


// Experimental.
type AwsIotCaCertificate_RegistrationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_ca_certificate#role_arn AwsIotCaCertificate#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_ca_certificate#template_body AwsIotCaCertificate#template_body}.
	// Experimental.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_ca_certificate#template_name AwsIotCaCertificate#template_name}.
	// Experimental.
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}

