package awssecretsmanager


// Experimental.
type AwsSecretsmanagerSecretRotation_ExternalSecretRotationMetadataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#key AwsSecretsmanagerSecretRotation#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#value AwsSecretsmanagerSecretRotation#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

