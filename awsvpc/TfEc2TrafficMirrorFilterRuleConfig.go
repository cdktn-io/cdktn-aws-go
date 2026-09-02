package awsvpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEc2TrafficMirrorFilterRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#destination_cidr_block TfEc2TrafficMirrorFilterRule#destination_cidr_block}.
	// Experimental.
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#rule_action TfEc2TrafficMirrorFilterRule#rule_action}.
	// Experimental.
	RuleAction *string `field:"required" json:"ruleAction" yaml:"ruleAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#rule_number TfEc2TrafficMirrorFilterRule#rule_number}.
	// Experimental.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#source_cidr_block TfEc2TrafficMirrorFilterRule#source_cidr_block}.
	// Experimental.
	SourceCidrBlock *string `field:"required" json:"sourceCidrBlock" yaml:"sourceCidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#traffic_direction TfEc2TrafficMirrorFilterRule#traffic_direction}.
	// Experimental.
	TrafficDirection *string `field:"required" json:"trafficDirection" yaml:"trafficDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#traffic_mirror_filter_id TfEc2TrafficMirrorFilterRule#traffic_mirror_filter_id}.
	// Experimental.
	TrafficMirrorFilterId *string `field:"required" json:"trafficMirrorFilterId" yaml:"trafficMirrorFilterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#description TfEc2TrafficMirrorFilterRule#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// destination_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#destination_port_range TfEc2TrafficMirrorFilterRule#destination_port_range}
	// Experimental.
	DestinationPortRange *TfEc2TrafficMirrorFilterRule_DestinationPortRangeProperty `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#id TfEc2TrafficMirrorFilterRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#protocol TfEc2TrafficMirrorFilterRule#protocol}.
	// Experimental.
	Protocol *float64 `field:"optional" json:"protocol" yaml:"protocol"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#region TfEc2TrafficMirrorFilterRule#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// source_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#source_port_range TfEc2TrafficMirrorFilterRule#source_port_range}
	// Experimental.
	SourcePortRange *TfEc2TrafficMirrorFilterRule_SourcePortRangeProperty `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
}

