package vpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEncryptionControlConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#mode AwsEncryptionControl#mode}.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#vpc_id AwsEncryptionControl#vpc_id}.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#egress_only_internet_gateway_exclusion AwsEncryptionControl#egress_only_internet_gateway_exclusion}.
	// Experimental.
	EgressOnlyInternetGatewayExclusion *string `field:"optional" json:"egressOnlyInternetGatewayExclusion" yaml:"egressOnlyInternetGatewayExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#elastic_file_system_exclusion AwsEncryptionControl#elastic_file_system_exclusion}.
	// Experimental.
	ElasticFileSystemExclusion *string `field:"optional" json:"elasticFileSystemExclusion" yaml:"elasticFileSystemExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#internet_gateway_exclusion AwsEncryptionControl#internet_gateway_exclusion}.
	// Experimental.
	InternetGatewayExclusion *string `field:"optional" json:"internetGatewayExclusion" yaml:"internetGatewayExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#lambda_exclusion AwsEncryptionControl#lambda_exclusion}.
	// Experimental.
	LambdaExclusion *string `field:"optional" json:"lambdaExclusion" yaml:"lambdaExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#nat_gateway_exclusion AwsEncryptionControl#nat_gateway_exclusion}.
	// Experimental.
	NatGatewayExclusion *string `field:"optional" json:"natGatewayExclusion" yaml:"natGatewayExclusion"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#region AwsEncryptionControl#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#tags AwsEncryptionControl#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#timeouts AwsEncryptionControl#timeouts}
	// Experimental.
	Timeouts *AwsEncryptionControl_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#virtual_private_gateway_exclusion AwsEncryptionControl#virtual_private_gateway_exclusion}.
	// Experimental.
	VirtualPrivateGatewayExclusion *string `field:"optional" json:"virtualPrivateGatewayExclusion" yaml:"virtualPrivateGatewayExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#vpc_lattice_exclusion AwsEncryptionControl#vpc_lattice_exclusion}.
	// Experimental.
	VpcLatticeExclusion *string `field:"optional" json:"vpcLatticeExclusion" yaml:"vpcLatticeExclusion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_encryption_control#vpc_peering_exclusion AwsEncryptionControl#vpc_peering_exclusion}.
	// Experimental.
	VpcPeeringExclusion *string `field:"optional" json:"vpcPeeringExclusion" yaml:"vpcPeeringExclusion"`
}

