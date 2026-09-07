package eks

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#name AwsCluster#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#role_arn AwsCluster#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#vpc_config AwsCluster#vpc_config}
	// Experimental.
	VpcConfig *AwsCluster_VpcConfigProperty `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#access_config AwsCluster#access_config}
	// Experimental.
	AccessConfig *AwsCluster_AccessConfigProperty `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#bootstrap_self_managed_addons AwsCluster#bootstrap_self_managed_addons}.
	// Experimental.
	BootstrapSelfManagedAddons interface{} `field:"optional" json:"bootstrapSelfManagedAddons" yaml:"bootstrapSelfManagedAddons"`
	// compute_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#compute_config AwsCluster#compute_config}
	// Experimental.
	ComputeConfig *AwsCluster_ComputeConfigProperty `field:"optional" json:"computeConfig" yaml:"computeConfig"`
	// control_plane_scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#control_plane_scaling_config AwsCluster#control_plane_scaling_config}
	// Experimental.
	ControlPlaneScalingConfig *AwsCluster_ControlPlaneScalingConfigProperty `field:"optional" json:"controlPlaneScalingConfig" yaml:"controlPlaneScalingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#deletion_protection AwsCluster#deletion_protection}.
	// Experimental.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#enabled_cluster_log_types AwsCluster#enabled_cluster_log_types}.
	// Experimental.
	EnabledClusterLogTypes *[]*string `field:"optional" json:"enabledClusterLogTypes" yaml:"enabledClusterLogTypes"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#encryption_config AwsCluster#encryption_config}
	// Experimental.
	EncryptionConfig *AwsCluster_EncryptionConfigProperty `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#force_update_version AwsCluster#force_update_version}.
	// Experimental.
	ForceUpdateVersion interface{} `field:"optional" json:"forceUpdateVersion" yaml:"forceUpdateVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#id AwsCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kube_api_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kube_api_server_config AwsCluster#kube_api_server_config}
	// Experimental.
	KubeApiServerConfig *AwsCluster_KubeApiServerConfigProperty `field:"optional" json:"kubeApiServerConfig" yaml:"kubeApiServerConfig"`
	// kube_controller_manager_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kube_controller_manager_config AwsCluster#kube_controller_manager_config}
	// Experimental.
	KubeControllerManagerConfig *AwsCluster_KubeControllerManagerConfigProperty `field:"optional" json:"kubeControllerManagerConfig" yaml:"kubeControllerManagerConfig"`
	// kubernetes_network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kubernetes_network_config AwsCluster#kubernetes_network_config}
	// Experimental.
	KubernetesNetworkConfig *AwsCluster_KubernetesNetworkConfigProperty `field:"optional" json:"kubernetesNetworkConfig" yaml:"kubernetesNetworkConfig"`
	// kube_scheduler_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kube_scheduler_config AwsCluster#kube_scheduler_config}
	// Experimental.
	KubeSchedulerConfig *AwsCluster_KubeSchedulerConfigProperty `field:"optional" json:"kubeSchedulerConfig" yaml:"kubeSchedulerConfig"`
	// outpost_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#outpost_config AwsCluster#outpost_config}
	// Experimental.
	OutpostConfig *AwsCluster_OutpostConfigProperty `field:"optional" json:"outpostConfig" yaml:"outpostConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#region AwsCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// remote_network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#remote_network_config AwsCluster#remote_network_config}
	// Experimental.
	RemoteNetworkConfig *AwsCluster_RemoteNetworkConfigProperty `field:"optional" json:"remoteNetworkConfig" yaml:"remoteNetworkConfig"`
	// storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#storage_config AwsCluster#storage_config}
	// Experimental.
	StorageConfig *AwsCluster_StorageConfigProperty `field:"optional" json:"storageConfig" yaml:"storageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#tags AwsCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#tags_all AwsCluster#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#timeouts AwsCluster#timeouts}
	// Experimental.
	Timeouts *AwsCluster_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#upgrade_policy AwsCluster#upgrade_policy}
	// Experimental.
	UpgradePolicy *AwsCluster_UpgradePolicyProperty `field:"optional" json:"upgradePolicy" yaml:"upgradePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#version AwsCluster#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// zonal_shift_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#zonal_shift_config AwsCluster#zonal_shift_config}
	// Experimental.
	ZonalShiftConfig *AwsCluster_ZonalShiftConfigProperty `field:"optional" json:"zonalShiftConfig" yaml:"zonalShiftConfig"`
}

