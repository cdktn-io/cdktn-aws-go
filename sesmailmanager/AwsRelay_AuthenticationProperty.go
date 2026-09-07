package sesmailmanager


// Experimental.
type AwsRelay_AuthenticationProperty struct {
	// no_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_relay#no_authentication AwsRelay#no_authentication}
	// Experimental.
	NoAuthentication interface{} `field:"optional" json:"noAuthentication" yaml:"noAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_relay#secret_arn AwsRelay#secret_arn}.
	// Experimental.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

