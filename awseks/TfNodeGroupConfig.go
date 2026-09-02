package awseks

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNodeGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#cluster_name TfNodeGroup#cluster_name}.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#node_role_arn TfNodeGroup#node_role_arn}.
	// Experimental.
	NodeRoleArn *string `field:"required" json:"nodeRoleArn" yaml:"nodeRoleArn"`
	// scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#scaling_config TfNodeGroup#scaling_config}
	// Experimental.
	ScalingConfig *TfNodeGroup_ScalingConfigProperty `field:"required" json:"scalingConfig" yaml:"scalingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#subnet_ids TfNodeGroup#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#ami_type TfNodeGroup#ami_type}.
	// Experimental.
	AmiType *string `field:"optional" json:"amiType" yaml:"amiType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#capacity_type TfNodeGroup#capacity_type}.
	// Experimental.
	CapacityType *string `field:"optional" json:"capacityType" yaml:"capacityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#disk_size TfNodeGroup#disk_size}.
	// Experimental.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#force_update_version TfNodeGroup#force_update_version}.
	// Experimental.
	ForceUpdateVersion interface{} `field:"optional" json:"forceUpdateVersion" yaml:"forceUpdateVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#id TfNodeGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#instance_types TfNodeGroup#instance_types}.
	// Experimental.
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#labels TfNodeGroup#labels}.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#launch_template TfNodeGroup#launch_template}
	// Experimental.
	LaunchTemplate *TfNodeGroup_LaunchTemplateProperty `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#node_group_name TfNodeGroup#node_group_name}.
	// Experimental.
	NodeGroupName *string `field:"optional" json:"nodeGroupName" yaml:"nodeGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#node_group_name_prefix TfNodeGroup#node_group_name_prefix}.
	// Experimental.
	NodeGroupNamePrefix *string `field:"optional" json:"nodeGroupNamePrefix" yaml:"nodeGroupNamePrefix"`
	// node_repair_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#node_repair_config TfNodeGroup#node_repair_config}
	// Experimental.
	NodeRepairConfig *TfNodeGroup_NodeRepairConfigProperty `field:"optional" json:"nodeRepairConfig" yaml:"nodeRepairConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#region TfNodeGroup#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#release_version TfNodeGroup#release_version}.
	// Experimental.
	ReleaseVersion *string `field:"optional" json:"releaseVersion" yaml:"releaseVersion"`
	// remote_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#remote_access TfNodeGroup#remote_access}
	// Experimental.
	RemoteAccess *TfNodeGroup_RemoteAccessProperty `field:"optional" json:"remoteAccess" yaml:"remoteAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#tags TfNodeGroup#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#tags_all TfNodeGroup#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// taint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#taint TfNodeGroup#taint}
	// Experimental.
	Taint interface{} `field:"optional" json:"taint" yaml:"taint"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#timeouts TfNodeGroup#timeouts}
	// Experimental.
	Timeouts *TfNodeGroup_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// update_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#update_config TfNodeGroup#update_config}
	// Experimental.
	UpdateConfig *TfNodeGroup_UpdateConfigProperty `field:"optional" json:"updateConfig" yaml:"updateConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#version TfNodeGroup#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// warm_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#warm_pool_config TfNodeGroup#warm_pool_config}
	// Experimental.
	WarmPoolConfig *TfNodeGroup_WarmPoolConfigProperty `field:"optional" json:"warmPoolConfig" yaml:"warmPoolConfig"`
}

