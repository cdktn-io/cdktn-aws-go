package awsvpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlowLogConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#deliver_cross_account_role TfFlowLog#deliver_cross_account_role}.
	// Experimental.
	DeliverCrossAccountRole *string `field:"optional" json:"deliverCrossAccountRole" yaml:"deliverCrossAccountRole"`
	// destination_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#destination_options TfFlowLog#destination_options}
	// Experimental.
	DestinationOptions *TfFlowLog_DestinationOptionsProperty `field:"optional" json:"destinationOptions" yaml:"destinationOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#eni_id TfFlowLog#eni_id}.
	// Experimental.
	EniId *string `field:"optional" json:"eniId" yaml:"eniId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#iam_role_arn TfFlowLog#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#id TfFlowLog#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#log_destination TfFlowLog#log_destination}.
	// Experimental.
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#log_destination_type TfFlowLog#log_destination_type}.
	// Experimental.
	LogDestinationType *string `field:"optional" json:"logDestinationType" yaml:"logDestinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#log_format TfFlowLog#log_format}.
	// Experimental.
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#max_aggregation_interval TfFlowLog#max_aggregation_interval}.
	// Experimental.
	MaxAggregationInterval *float64 `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#region TfFlowLog#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#regional_nat_gateway_id TfFlowLog#regional_nat_gateway_id}.
	// Experimental.
	RegionalNatGatewayId *string `field:"optional" json:"regionalNatGatewayId" yaml:"regionalNatGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#subnet_id TfFlowLog#subnet_id}.
	// Experimental.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// tag_field_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#tag_field_specification TfFlowLog#tag_field_specification}
	// Experimental.
	TagFieldSpecification interface{} `field:"optional" json:"tagFieldSpecification" yaml:"tagFieldSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#tags TfFlowLog#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#tags_all TfFlowLog#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#traffic_type TfFlowLog#traffic_type}.
	// Experimental.
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#transit_gateway_attachment_id TfFlowLog#transit_gateway_attachment_id}.
	// Experimental.
	TransitGatewayAttachmentId *string `field:"optional" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#transit_gateway_id TfFlowLog#transit_gateway_id}.
	// Experimental.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/flow_log#vpc_id TfFlowLog#vpc_id}.
	// Experimental.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

