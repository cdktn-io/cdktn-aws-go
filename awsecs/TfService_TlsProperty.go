package awsecs


// Experimental.
type TfService_TlsProperty struct {
	// issuer_cert_authority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#issuer_cert_authority TfService#issuer_cert_authority}
	// Experimental.
	IssuerCertAuthority *TfService_IssuerCertAuthorityProperty `field:"required" json:"issuerCertAuthority" yaml:"issuerCertAuthority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#kms_key TfService#kms_key}.
	// Experimental.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#role_arn TfService#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

