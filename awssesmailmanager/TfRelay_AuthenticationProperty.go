package awssesmailmanager


// Experimental.
type TfRelay_AuthenticationProperty struct {
	// no_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_relay#no_authentication TfRelay#no_authentication}
	// Experimental.
	NoAuthentication interface{} `field:"optional" json:"noAuthentication" yaml:"noAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_relay#secret_arn TfRelay#secret_arn}.
	// Experimental.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

