package ssoadmin


// Experimental.
type AwsPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_permissions_boundary_attachment#name AwsPermissionsBoundaryAttachment#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_permissions_boundary_attachment#path AwsPermissionsBoundaryAttachment#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

