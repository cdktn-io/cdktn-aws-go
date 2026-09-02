package awsappstream20

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFleetConfig struct {
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
	// compute_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#compute_capacity TfFleet#compute_capacity}
	// Experimental.
	ComputeCapacity *TfFleet_ComputeCapacityProperty `field:"required" json:"computeCapacity" yaml:"computeCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#instance_type TfFleet#instance_type}.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#name TfFleet#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#description TfFleet#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#disconnect_timeout_in_seconds TfFleet#disconnect_timeout_in_seconds}.
	// Experimental.
	DisconnectTimeoutInSeconds *float64 `field:"optional" json:"disconnectTimeoutInSeconds" yaml:"disconnectTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#display_name TfFleet#display_name}.
	// Experimental.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// domain_join_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#domain_join_info TfFleet#domain_join_info}
	// Experimental.
	DomainJoinInfo *TfFleet_DomainJoinInfoProperty `field:"optional" json:"domainJoinInfo" yaml:"domainJoinInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#enable_default_internet_access TfFleet#enable_default_internet_access}.
	// Experimental.
	EnableDefaultInternetAccess interface{} `field:"optional" json:"enableDefaultInternetAccess" yaml:"enableDefaultInternetAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#fleet_type TfFleet#fleet_type}.
	// Experimental.
	FleetType *string `field:"optional" json:"fleetType" yaml:"fleetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#iam_role_arn TfFleet#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#id TfFleet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#idle_disconnect_timeout_in_seconds TfFleet#idle_disconnect_timeout_in_seconds}.
	// Experimental.
	IdleDisconnectTimeoutInSeconds *float64 `field:"optional" json:"idleDisconnectTimeoutInSeconds" yaml:"idleDisconnectTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#image_arn TfFleet#image_arn}.
	// Experimental.
	ImageArn *string `field:"optional" json:"imageArn" yaml:"imageArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#image_name TfFleet#image_name}.
	// Experimental.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#max_sessions_per_instance TfFleet#max_sessions_per_instance}.
	// Experimental.
	MaxSessionsPerInstance *float64 `field:"optional" json:"maxSessionsPerInstance" yaml:"maxSessionsPerInstance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#max_user_duration_in_seconds TfFleet#max_user_duration_in_seconds}.
	// Experimental.
	MaxUserDurationInSeconds *float64 `field:"optional" json:"maxUserDurationInSeconds" yaml:"maxUserDurationInSeconds"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#region TfFleet#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#stream_view TfFleet#stream_view}.
	// Experimental.
	StreamView *string `field:"optional" json:"streamView" yaml:"streamView"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#tags TfFleet#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#tags_all TfFleet#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#vpc_config TfFleet#vpc_config}
	// Experimental.
	VpcConfig *TfFleet_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

