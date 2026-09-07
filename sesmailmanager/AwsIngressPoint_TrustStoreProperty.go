package sesmailmanager


// Experimental.
type AwsIngressPoint_TrustStoreProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#ca_content AwsIngressPoint#ca_content}.
	// Experimental.
	CaContent *string `field:"required" json:"caContent" yaml:"caContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#crl_content AwsIngressPoint#crl_content}.
	// Experimental.
	CrlContent *string `field:"optional" json:"crlContent" yaml:"crlContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#kms_key_arn AwsIngressPoint#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

