package awsec2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEipConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#address TfEip#address}.
	// Experimental.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#associate_with_private_ip TfEip#associate_with_private_ip}.
	// Experimental.
	AssociateWithPrivateIp *string `field:"optional" json:"associateWithPrivateIp" yaml:"associateWithPrivateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#customer_owned_ipv4_pool TfEip#customer_owned_ipv4_pool}.
	// Experimental.
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#domain TfEip#domain}.
	// Experimental.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#id TfEip#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#instance TfEip#instance}.
	// Experimental.
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#ipam_pool_id TfEip#ipam_pool_id}.
	// Experimental.
	IpamPoolId *string `field:"optional" json:"ipamPoolId" yaml:"ipamPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#network_border_group TfEip#network_border_group}.
	// Experimental.
	NetworkBorderGroup *string `field:"optional" json:"networkBorderGroup" yaml:"networkBorderGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#network_interface TfEip#network_interface}.
	// Experimental.
	NetworkInterface *string `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#public_ipv4_pool TfEip#public_ipv4_pool}.
	// Experimental.
	PublicIpv4Pool *string `field:"optional" json:"publicIpv4Pool" yaml:"publicIpv4Pool"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#region TfEip#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#tags TfEip#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#tags_all TfEip#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eip#timeouts TfEip#timeouts}
	// Experimental.
	Timeouts *TfEip_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

