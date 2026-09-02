package awssecretsmanager


// Experimental.
type DataTfSecrets_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/secretsmanager_secrets#name DataTfSecrets#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/secretsmanager_secrets#values DataTfSecrets#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

