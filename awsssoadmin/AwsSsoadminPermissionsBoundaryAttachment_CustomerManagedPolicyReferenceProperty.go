package awsssoadmin


// Experimental.
type AwsSsoadminPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_permissions_boundary_attachment#name AwsSsoadminPermissionsBoundaryAttachment#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_permissions_boundary_attachment#path AwsSsoadminPermissionsBoundaryAttachment#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

