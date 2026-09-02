package awsstoragegateway


// Experimental.
type TfGateway_SmbActiveDirectorySettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#domain_name TfGateway#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#password TfGateway#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#username TfGateway#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#domain_controllers TfGateway#domain_controllers}.
	// Experimental.
	DomainControllers *[]*string `field:"optional" json:"domainControllers" yaml:"domainControllers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#organizational_unit TfGateway#organizational_unit}.
	// Experimental.
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#timeout_in_seconds TfGateway#timeout_in_seconds}.
	// Experimental.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

