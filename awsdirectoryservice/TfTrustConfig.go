package awsdirectoryservice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTrustConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_trust#directory_id TfTrust#directory_id}.
	// Experimental.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_trust#remote_domain_name TfTrust#remote_domain_name}.
	// Experimental.
	RemoteDomainName *string `field:"required" json:"remoteDomainName" yaml:"remoteDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_trust#trust_direction TfTrust#trust_direction}.
	// Experimental.
	TrustDirection *string `field:"required" json:"trustDirection" yaml:"trustDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_trust#trust_password TfTrust#trust_password}.
	// Experimental.
	TrustPassword *string `field:"required" json:"trustPassword" yaml:"trustPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_trust#conditional_forwarder_ip_addrs TfTrust#conditional_forwarder_ip_addrs}.
	// Experimental.
	ConditionalForwarderIpAddrs *[]*string `field:"optional" json:"conditionalForwarderIpAddrs" yaml:"conditionalForwarderIpAddrs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_trust#delete_associated_conditional_forwarder TfTrust#delete_associated_conditional_forwarder}.
	// Experimental.
	DeleteAssociatedConditionalForwarder interface{} `field:"optional" json:"deleteAssociatedConditionalForwarder" yaml:"deleteAssociatedConditionalForwarder"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_trust#region TfTrust#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_trust#selective_auth TfTrust#selective_auth}.
	// Experimental.
	SelectiveAuth *string `field:"optional" json:"selectiveAuth" yaml:"selectiveAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_trust#trust_type TfTrust#trust_type}.
	// Experimental.
	TrustType *string `field:"optional" json:"trustType" yaml:"trustType"`
}

