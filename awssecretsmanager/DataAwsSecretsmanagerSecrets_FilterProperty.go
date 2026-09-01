package awssecretsmanager


// Experimental.
type DataAwsSecretsmanagerSecrets_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/secretsmanager_secrets#name DataAwsSecretsmanagerSecrets#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/secretsmanager_secrets#values DataAwsSecretsmanagerSecrets#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

