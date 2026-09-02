package awsssoadmin


// Experimental.
type TfPermissionsBoundaryAttachment_PermissionsBoundaryProperty struct {
	// customer_managed_policy_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_permissions_boundary_attachment#customer_managed_policy_reference TfPermissionsBoundaryAttachment#customer_managed_policy_reference}
	// Experimental.
	CustomerManagedPolicyReference *TfPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty `field:"optional" json:"customerManagedPolicyReference" yaml:"customerManagedPolicyReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_permissions_boundary_attachment#managed_policy_arn TfPermissionsBoundaryAttachment#managed_policy_arn}.
	// Experimental.
	ManagedPolicyArn *string `field:"optional" json:"managedPolicyArn" yaml:"managedPolicyArn"`
}

