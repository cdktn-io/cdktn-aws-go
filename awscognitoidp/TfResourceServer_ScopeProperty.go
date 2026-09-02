package awscognitoidp


// Experimental.
type TfResourceServer_ScopeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_resource_server#scope_description TfResourceServer#scope_description}.
	// Experimental.
	ScopeDescription *string `field:"required" json:"scopeDescription" yaml:"scopeDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_resource_server#scope_name TfResourceServer#scope_name}.
	// Experimental.
	ScopeName *string `field:"required" json:"scopeName" yaml:"scopeName"`
}

