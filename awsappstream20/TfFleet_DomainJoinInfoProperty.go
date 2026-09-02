package awsappstream20


// Experimental.
type TfFleet_DomainJoinInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#directory_name TfFleet#directory_name}.
	// Experimental.
	DirectoryName *string `field:"optional" json:"directoryName" yaml:"directoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#organizational_unit_distinguished_name TfFleet#organizational_unit_distinguished_name}.
	// Experimental.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

