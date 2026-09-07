package vpclattice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsResourceConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#name AwsResourceConfiguration#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#allow_association_to_shareable_service_network AwsResourceConfiguration#allow_association_to_shareable_service_network}.
	// Experimental.
	AllowAssociationToShareableServiceNetwork interface{} `field:"optional" json:"allowAssociationToShareableServiceNetwork" yaml:"allowAssociationToShareableServiceNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#custom_domain_name AwsResourceConfiguration#custom_domain_name}.
	// Experimental.
	CustomDomainName *string `field:"optional" json:"customDomainName" yaml:"customDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#domain_verification_id AwsResourceConfiguration#domain_verification_id}.
	// Experimental.
	DomainVerificationId *string `field:"optional" json:"domainVerificationId" yaml:"domainVerificationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#port_ranges AwsResourceConfiguration#port_ranges}.
	// Experimental.
	PortRanges *[]*string `field:"optional" json:"portRanges" yaml:"portRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#protocol AwsResourceConfiguration#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#region AwsResourceConfiguration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// resource_configuration_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#resource_configuration_definition AwsResourceConfiguration#resource_configuration_definition}
	// Experimental.
	ResourceConfigurationDefinition interface{} `field:"optional" json:"resourceConfigurationDefinition" yaml:"resourceConfigurationDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#resource_configuration_group_id AwsResourceConfiguration#resource_configuration_group_id}.
	// Experimental.
	ResourceConfigurationGroupId *string `field:"optional" json:"resourceConfigurationGroupId" yaml:"resourceConfigurationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#resource_gateway_identifier AwsResourceConfiguration#resource_gateway_identifier}.
	// Experimental.
	ResourceGatewayIdentifier *string `field:"optional" json:"resourceGatewayIdentifier" yaml:"resourceGatewayIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#tags AwsResourceConfiguration#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#timeouts AwsResourceConfiguration#timeouts}
	// Experimental.
	Timeouts *AwsResourceConfiguration_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_resource_configuration#type AwsResourceConfiguration#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

