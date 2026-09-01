package awsecs


// Experimental.
type AwsEcsService_TlsProperty struct {
	// issuer_cert_authority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#issuer_cert_authority AwsEcsService#issuer_cert_authority}
	// Experimental.
	IssuerCertAuthority *AwsEcsService_IssuerCertAuthorityProperty `field:"required" json:"issuerCertAuthority" yaml:"issuerCertAuthority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#kms_key AwsEcsService#kms_key}.
	// Experimental.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#role_arn AwsEcsService#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

