package awsvpclattice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfServiceNetworkVpcAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#service_network_identifier TfServiceNetworkVpcAssociation#service_network_identifier}.
	// Experimental.
	ServiceNetworkIdentifier *string `field:"required" json:"serviceNetworkIdentifier" yaml:"serviceNetworkIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#vpc_identifier TfServiceNetworkVpcAssociation#vpc_identifier}.
	// Experimental.
	VpcIdentifier *string `field:"required" json:"vpcIdentifier" yaml:"vpcIdentifier"`
	// dns_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#dns_options TfServiceNetworkVpcAssociation#dns_options}
	// Experimental.
	DnsOptions *TfServiceNetworkVpcAssociation_DnsOptionsProperty `field:"optional" json:"dnsOptions" yaml:"dnsOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#id TfServiceNetworkVpcAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#private_dns_enabled TfServiceNetworkVpcAssociation#private_dns_enabled}.
	// Experimental.
	PrivateDnsEnabled interface{} `field:"optional" json:"privateDnsEnabled" yaml:"privateDnsEnabled"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#region TfServiceNetworkVpcAssociation#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#security_group_ids TfServiceNetworkVpcAssociation#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#tags TfServiceNetworkVpcAssociation#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#tags_all TfServiceNetworkVpcAssociation#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#timeouts TfServiceNetworkVpcAssociation#timeouts}
	// Experimental.
	Timeouts *TfServiceNetworkVpcAssociation_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

