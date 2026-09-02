package awselasticache


// Experimental.
type TfUser_AuthenticationModeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_user#type TfUser#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_user#passwords TfUser#passwords}.
	// Experimental.
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
}

