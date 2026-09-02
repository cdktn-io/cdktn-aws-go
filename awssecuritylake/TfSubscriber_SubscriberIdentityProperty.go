package awssecuritylake


// Experimental.
type TfSubscriber_SubscriberIdentityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#external_id TfSubscriber#external_id}.
	// Experimental.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#principal TfSubscriber#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

