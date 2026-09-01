package awssecuritylake


// Experimental.
type AwsSecuritylakeSubscriber_SubscriberIdentityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#external_id AwsSecuritylakeSubscriber#external_id}.
	// Experimental.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#principal AwsSecuritylakeSubscriber#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

