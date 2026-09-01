package awseks

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEksClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#name AwsEksCluster#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#role_arn AwsEksCluster#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#vpc_config AwsEksCluster#vpc_config}
	// Experimental.
	VpcConfig *AwsEksCluster_VpcConfigProperty `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#access_config AwsEksCluster#access_config}
	// Experimental.
	AccessConfig *AwsEksCluster_AccessConfigProperty `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#bootstrap_self_managed_addons AwsEksCluster#bootstrap_self_managed_addons}.
	// Experimental.
	BootstrapSelfManagedAddons interface{} `field:"optional" json:"bootstrapSelfManagedAddons" yaml:"bootstrapSelfManagedAddons"`
	// compute_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#compute_config AwsEksCluster#compute_config}
	// Experimental.
	ComputeConfig *AwsEksCluster_ComputeConfigProperty `field:"optional" json:"computeConfig" yaml:"computeConfig"`
	// control_plane_scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#control_plane_scaling_config AwsEksCluster#control_plane_scaling_config}
	// Experimental.
	ControlPlaneScalingConfig *AwsEksCluster_ControlPlaneScalingConfigProperty `field:"optional" json:"controlPlaneScalingConfig" yaml:"controlPlaneScalingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#deletion_protection AwsEksCluster#deletion_protection}.
	// Experimental.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#enabled_cluster_log_types AwsEksCluster#enabled_cluster_log_types}.
	// Experimental.
	EnabledClusterLogTypes *[]*string `field:"optional" json:"enabledClusterLogTypes" yaml:"enabledClusterLogTypes"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#encryption_config AwsEksCluster#encryption_config}
	// Experimental.
	EncryptionConfig *AwsEksCluster_EncryptionConfigProperty `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#force_update_version AwsEksCluster#force_update_version}.
	// Experimental.
	ForceUpdateVersion interface{} `field:"optional" json:"forceUpdateVersion" yaml:"forceUpdateVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#id AwsEksCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kube_api_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kube_api_server_config AwsEksCluster#kube_api_server_config}
	// Experimental.
	KubeApiServerConfig *AwsEksCluster_KubeApiServerConfigProperty `field:"optional" json:"kubeApiServerConfig" yaml:"kubeApiServerConfig"`
	// kube_controller_manager_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kube_controller_manager_config AwsEksCluster#kube_controller_manager_config}
	// Experimental.
	KubeControllerManagerConfig *AwsEksCluster_KubeControllerManagerConfigProperty `field:"optional" json:"kubeControllerManagerConfig" yaml:"kubeControllerManagerConfig"`
	// kubernetes_network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kubernetes_network_config AwsEksCluster#kubernetes_network_config}
	// Experimental.
	KubernetesNetworkConfig *AwsEksCluster_KubernetesNetworkConfigProperty `field:"optional" json:"kubernetesNetworkConfig" yaml:"kubernetesNetworkConfig"`
	// kube_scheduler_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kube_scheduler_config AwsEksCluster#kube_scheduler_config}
	// Experimental.
	KubeSchedulerConfig *AwsEksCluster_KubeSchedulerConfigProperty `field:"optional" json:"kubeSchedulerConfig" yaml:"kubeSchedulerConfig"`
	// outpost_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#outpost_config AwsEksCluster#outpost_config}
	// Experimental.
	OutpostConfig *AwsEksCluster_OutpostConfigProperty `field:"optional" json:"outpostConfig" yaml:"outpostConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#region AwsEksCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// remote_network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#remote_network_config AwsEksCluster#remote_network_config}
	// Experimental.
	RemoteNetworkConfig *AwsEksCluster_RemoteNetworkConfigProperty `field:"optional" json:"remoteNetworkConfig" yaml:"remoteNetworkConfig"`
	// storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#storage_config AwsEksCluster#storage_config}
	// Experimental.
	StorageConfig *AwsEksCluster_StorageConfigProperty `field:"optional" json:"storageConfig" yaml:"storageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#tags AwsEksCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#tags_all AwsEksCluster#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#timeouts AwsEksCluster#timeouts}
	// Experimental.
	Timeouts *AwsEksCluster_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#upgrade_policy AwsEksCluster#upgrade_policy}
	// Experimental.
	UpgradePolicy *AwsEksCluster_UpgradePolicyProperty `field:"optional" json:"upgradePolicy" yaml:"upgradePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#version AwsEksCluster#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// zonal_shift_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#zonal_shift_config AwsEksCluster#zonal_shift_config}
	// Experimental.
	ZonalShiftConfig *AwsEksCluster_ZonalShiftConfigProperty `field:"optional" json:"zonalShiftConfig" yaml:"zonalShiftConfig"`
}

