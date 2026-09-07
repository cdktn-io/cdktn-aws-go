package ssoadmin


// Experimental.
type AwsCustomerManagedPolicyAttachment_CustomerManagedPolicyReferenceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_customer_managed_policy_attachment#name AwsCustomerManagedPolicyAttachment#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_customer_managed_policy_attachment#path AwsCustomerManagedPolicyAttachment#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

