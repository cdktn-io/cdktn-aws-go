package awsapigateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomainNameConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#domain_name TfDomainName#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#certificate_arn TfDomainName#certificate_arn}.
	// Experimental.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#certificate_body TfDomainName#certificate_body}.
	// Experimental.
	CertificateBody *string `field:"optional" json:"certificateBody" yaml:"certificateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#certificate_chain TfDomainName#certificate_chain}.
	// Experimental.
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#certificate_name TfDomainName#certificate_name}.
	// Experimental.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#certificate_private_key TfDomainName#certificate_private_key}.
	// Experimental.
	CertificatePrivateKey *string `field:"optional" json:"certificatePrivateKey" yaml:"certificatePrivateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#endpoint_access_mode TfDomainName#endpoint_access_mode}.
	// Experimental.
	EndpointAccessMode *string `field:"optional" json:"endpointAccessMode" yaml:"endpointAccessMode"`
	// endpoint_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#endpoint_configuration TfDomainName#endpoint_configuration}
	// Experimental.
	EndpointConfiguration *TfDomainName_EndpointConfigurationProperty `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#id TfDomainName#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mutual_tls_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#mutual_tls_authentication TfDomainName#mutual_tls_authentication}
	// Experimental.
	MutualTlsAuthentication *TfDomainName_MutualTlsAuthenticationProperty `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#ownership_verification_certificate_arn TfDomainName#ownership_verification_certificate_arn}.
	// Experimental.
	OwnershipVerificationCertificateArn *string `field:"optional" json:"ownershipVerificationCertificateArn" yaml:"ownershipVerificationCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#policy TfDomainName#policy}.
	// Experimental.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#region TfDomainName#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#regional_certificate_arn TfDomainName#regional_certificate_arn}.
	// Experimental.
	RegionalCertificateArn *string `field:"optional" json:"regionalCertificateArn" yaml:"regionalCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#regional_certificate_name TfDomainName#regional_certificate_name}.
	// Experimental.
	RegionalCertificateName *string `field:"optional" json:"regionalCertificateName" yaml:"regionalCertificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#routing_mode TfDomainName#routing_mode}.
	// Experimental.
	RoutingMode *string `field:"optional" json:"routingMode" yaml:"routingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#security_policy TfDomainName#security_policy}.
	// Experimental.
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#tags TfDomainName#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#tags_all TfDomainName#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name#timeouts TfDomainName#timeouts}
	// Experimental.
	Timeouts *TfDomainName_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

