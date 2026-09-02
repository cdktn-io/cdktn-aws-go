package awselasticache


// Experimental.
type DataTfUser_AuthenticationModeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/elasticache_user#password_count DataTfUser#password_count}.
	// Experimental.
	PasswordCount *float64 `field:"optional" json:"passwordCount" yaml:"passwordCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/elasticache_user#type DataTfUser#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

