package secretsmanager


// Experimental.
type AwsSecretRotation_ExternalSecretRotationMetadataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#key AwsSecretRotation#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#value AwsSecretRotation#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

