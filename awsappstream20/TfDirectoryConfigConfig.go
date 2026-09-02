package awsappstream20

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDirectoryConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_directory_config#directory_name TfDirectoryConfig#directory_name}.
	// Experimental.
	DirectoryName *string `field:"required" json:"directoryName" yaml:"directoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_directory_config#organizational_unit_distinguished_names TfDirectoryConfig#organizational_unit_distinguished_names}.
	// Experimental.
	OrganizationalUnitDistinguishedNames *[]*string `field:"required" json:"organizationalUnitDistinguishedNames" yaml:"organizationalUnitDistinguishedNames"`
	// service_account_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_directory_config#service_account_credentials TfDirectoryConfig#service_account_credentials}
	// Experimental.
	ServiceAccountCredentials *TfDirectoryConfig_ServiceAccountCredentialsProperty `field:"required" json:"serviceAccountCredentials" yaml:"serviceAccountCredentials"`
	// certificate_based_auth_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_directory_config#certificate_based_auth_properties TfDirectoryConfig#certificate_based_auth_properties}
	// Experimental.
	CertificateBasedAuthProperties *TfDirectoryConfig_CertificateBasedAuthPropertiesProperty `field:"optional" json:"certificateBasedAuthProperties" yaml:"certificateBasedAuthProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_directory_config#id TfDirectoryConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_directory_config#region TfDirectoryConfig#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

