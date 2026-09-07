package transitgateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMeteringPolicyEntryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#metered_account AwsMeteringPolicyEntry#metered_account}.
	// Experimental.
	MeteredAccount *string `field:"required" json:"meteredAccount" yaml:"meteredAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#policy_rule_number AwsMeteringPolicyEntry#policy_rule_number}.
	// Experimental.
	PolicyRuleNumber *float64 `field:"required" json:"policyRuleNumber" yaml:"policyRuleNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#transit_gateway_metering_policy_id AwsMeteringPolicyEntry#transit_gateway_metering_policy_id}.
	// Experimental.
	TransitGatewayMeteringPolicyId *string `field:"required" json:"transitGatewayMeteringPolicyId" yaml:"transitGatewayMeteringPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#destination_cidr_block AwsMeteringPolicyEntry#destination_cidr_block}.
	// Experimental.
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#destination_port_range AwsMeteringPolicyEntry#destination_port_range}.
	// Experimental.
	DestinationPortRange *string `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#destination_transit_gateway_attachment_id AwsMeteringPolicyEntry#destination_transit_gateway_attachment_id}.
	// Experimental.
	DestinationTransitGatewayAttachmentId *string `field:"optional" json:"destinationTransitGatewayAttachmentId" yaml:"destinationTransitGatewayAttachmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#destination_transit_gateway_attachment_type AwsMeteringPolicyEntry#destination_transit_gateway_attachment_type}.
	// Experimental.
	DestinationTransitGatewayAttachmentType *string `field:"optional" json:"destinationTransitGatewayAttachmentType" yaml:"destinationTransitGatewayAttachmentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#protocol AwsMeteringPolicyEntry#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#region AwsMeteringPolicyEntry#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#source_cidr_block AwsMeteringPolicyEntry#source_cidr_block}.
	// Experimental.
	SourceCidrBlock *string `field:"optional" json:"sourceCidrBlock" yaml:"sourceCidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#source_port_range AwsMeteringPolicyEntry#source_port_range}.
	// Experimental.
	SourcePortRange *string `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#source_transit_gateway_attachment_id AwsMeteringPolicyEntry#source_transit_gateway_attachment_id}.
	// Experimental.
	SourceTransitGatewayAttachmentId *string `field:"optional" json:"sourceTransitGatewayAttachmentId" yaml:"sourceTransitGatewayAttachmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#source_transit_gateway_attachment_type AwsMeteringPolicyEntry#source_transit_gateway_attachment_type}.
	// Experimental.
	SourceTransitGatewayAttachmentType *string `field:"optional" json:"sourceTransitGatewayAttachmentType" yaml:"sourceTransitGatewayAttachmentType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry#timeouts AwsMeteringPolicyEntry#timeouts}
	// Experimental.
	Timeouts *AwsMeteringPolicyEntry_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

