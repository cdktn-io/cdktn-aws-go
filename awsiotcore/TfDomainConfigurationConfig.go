package awsiotcore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomainConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#name TfDomainConfiguration#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#application_protocol TfDomainConfiguration#application_protocol}.
	// Experimental.
	ApplicationProtocol *string `field:"optional" json:"applicationProtocol" yaml:"applicationProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#authentication_type TfDomainConfiguration#authentication_type}.
	// Experimental.
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#authorizer_config TfDomainConfiguration#authorizer_config}
	// Experimental.
	AuthorizerConfig *TfDomainConfiguration_AuthorizerConfigProperty `field:"optional" json:"authorizerConfig" yaml:"authorizerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#domain_name TfDomainConfiguration#domain_name}.
	// Experimental.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#id TfDomainConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#region TfDomainConfiguration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#server_certificate_arns TfDomainConfiguration#server_certificate_arns}.
	// Experimental.
	ServerCertificateArns *[]*string `field:"optional" json:"serverCertificateArns" yaml:"serverCertificateArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#service_type TfDomainConfiguration#service_type}.
	// Experimental.
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#status TfDomainConfiguration#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#tags TfDomainConfiguration#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#tags_all TfDomainConfiguration#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#tls_config TfDomainConfiguration#tls_config}
	// Experimental.
	TlsConfig *TfDomainConfiguration_TlsConfigProperty `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_domain_configuration#validation_certificate_arn TfDomainConfiguration#validation_certificate_arn}.
	// Experimental.
	ValidationCertificateArn *string `field:"optional" json:"validationCertificateArn" yaml:"validationCertificateArn"`
}

