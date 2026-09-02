package awsopensearch

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOutboundConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#connection_alias TfOutboundConnection#connection_alias}.
	// Experimental.
	ConnectionAlias *string `field:"required" json:"connectionAlias" yaml:"connectionAlias"`
	// local_domain_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#local_domain_info TfOutboundConnection#local_domain_info}
	// Experimental.
	LocalDomainInfo *TfOutboundConnection_LocalDomainInfoProperty `field:"required" json:"localDomainInfo" yaml:"localDomainInfo"`
	// remote_domain_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#remote_domain_info TfOutboundConnection#remote_domain_info}
	// Experimental.
	RemoteDomainInfo *TfOutboundConnection_RemoteDomainInfoProperty `field:"required" json:"remoteDomainInfo" yaml:"remoteDomainInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#accept_connection TfOutboundConnection#accept_connection}.
	// Experimental.
	AcceptConnection interface{} `field:"optional" json:"acceptConnection" yaml:"acceptConnection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#connection_mode TfOutboundConnection#connection_mode}.
	// Experimental.
	ConnectionMode *string `field:"optional" json:"connectionMode" yaml:"connectionMode"`
	// connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#connection_properties TfOutboundConnection#connection_properties}
	// Experimental.
	ConnectionProperties *TfOutboundConnection_ConnectionPropertiesProperty `field:"optional" json:"connectionProperties" yaml:"connectionProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#id TfOutboundConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#region TfOutboundConnection#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#timeouts TfOutboundConnection#timeouts}
	// Experimental.
	Timeouts *TfOutboundConnection_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

