package awsvpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPeeringConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#peer_vpc_id TfPeeringConnection#peer_vpc_id}.
	// Experimental.
	PeerVpcId *string `field:"required" json:"peerVpcId" yaml:"peerVpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#vpc_id TfPeeringConnection#vpc_id}.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// accepter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#accepter TfPeeringConnection#accepter}
	// Experimental.
	Accepter *TfPeeringConnection_AccepterProperty `field:"optional" json:"accepter" yaml:"accepter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#auto_accept TfPeeringConnection#auto_accept}.
	// Experimental.
	AutoAccept interface{} `field:"optional" json:"autoAccept" yaml:"autoAccept"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#id TfPeeringConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#peer_owner_id TfPeeringConnection#peer_owner_id}.
	// Experimental.
	PeerOwnerId *string `field:"optional" json:"peerOwnerId" yaml:"peerOwnerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#peer_region TfPeeringConnection#peer_region}.
	// Experimental.
	PeerRegion *string `field:"optional" json:"peerRegion" yaml:"peerRegion"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#region TfPeeringConnection#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// requester block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#requester TfPeeringConnection#requester}
	// Experimental.
	Requester *TfPeeringConnection_RequesterProperty `field:"optional" json:"requester" yaml:"requester"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#tags TfPeeringConnection#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#tags_all TfPeeringConnection#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#timeouts TfPeeringConnection#timeouts}
	// Experimental.
	Timeouts *TfPeeringConnection_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

