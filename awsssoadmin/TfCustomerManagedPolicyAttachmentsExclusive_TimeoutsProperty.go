package awsssoadmin


// Experimental.
type TfCustomerManagedPolicyAttachmentsExclusive_TimeoutsProperty struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_customer_managed_policy_attachments_exclusive#create TfCustomerManagedPolicyAttachmentsExclusive#create}
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_customer_managed_policy_attachments_exclusive#update TfCustomerManagedPolicyAttachmentsExclusive#update}
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

