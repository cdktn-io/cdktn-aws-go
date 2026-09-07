package elasticache


// Experimental.
type DataAwsUser_AuthenticationModeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/elasticache_user#password_count DataAwsUser#password_count}.
	// Experimental.
	PasswordCount *float64 `field:"optional" json:"passwordCount" yaml:"passwordCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/elasticache_user#type DataAwsUser#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

