package awsappfabric


// Experimental.
type TfAppAuthorization_TenantProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_app_authorization#tenant_display_name TfAppAuthorization#tenant_display_name}.
	// Experimental.
	TenantDisplayName *string `field:"required" json:"tenantDisplayName" yaml:"tenantDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_app_authorization#tenant_identifier TfAppAuthorization#tenant_identifier}.
	// Experimental.
	TenantIdentifier *string `field:"required" json:"tenantIdentifier" yaml:"tenantIdentifier"`
}

