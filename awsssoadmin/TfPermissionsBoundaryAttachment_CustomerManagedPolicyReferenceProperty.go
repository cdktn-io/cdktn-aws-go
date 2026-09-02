package awsssoadmin


// Experimental.
type TfPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_permissions_boundary_attachment#name TfPermissionsBoundaryAttachment#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_permissions_boundary_attachment#path TfPermissionsBoundaryAttachment#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

