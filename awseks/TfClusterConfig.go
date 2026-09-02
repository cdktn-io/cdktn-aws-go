package awseks

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#name TfCluster#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#role_arn TfCluster#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#vpc_config TfCluster#vpc_config}
	// Experimental.
	VpcConfig *TfCluster_VpcConfigProperty `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#access_config TfCluster#access_config}
	// Experimental.
	AccessConfig *TfCluster_AccessConfigProperty `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#bootstrap_self_managed_addons TfCluster#bootstrap_self_managed_addons}.
	// Experimental.
	BootstrapSelfManagedAddons interface{} `field:"optional" json:"bootstrapSelfManagedAddons" yaml:"bootstrapSelfManagedAddons"`
	// compute_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#compute_config TfCluster#compute_config}
	// Experimental.
	ComputeConfig *TfCluster_ComputeConfigProperty `field:"optional" json:"computeConfig" yaml:"computeConfig"`
	// control_plane_scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#control_plane_scaling_config TfCluster#control_plane_scaling_config}
	// Experimental.
	ControlPlaneScalingConfig *TfCluster_ControlPlaneScalingConfigProperty `field:"optional" json:"controlPlaneScalingConfig" yaml:"controlPlaneScalingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#deletion_protection TfCluster#deletion_protection}.
	// Experimental.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#enabled_cluster_log_types TfCluster#enabled_cluster_log_types}.
	// Experimental.
	EnabledClusterLogTypes *[]*string `field:"optional" json:"enabledClusterLogTypes" yaml:"enabledClusterLogTypes"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#encryption_config TfCluster#encryption_config}
	// Experimental.
	EncryptionConfig *TfCluster_EncryptionConfigProperty `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#force_update_version TfCluster#force_update_version}.
	// Experimental.
	ForceUpdateVersion interface{} `field:"optional" json:"forceUpdateVersion" yaml:"forceUpdateVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#id TfCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kube_api_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kube_api_server_config TfCluster#kube_api_server_config}
	// Experimental.
	KubeApiServerConfig *TfCluster_KubeApiServerConfigProperty `field:"optional" json:"kubeApiServerConfig" yaml:"kubeApiServerConfig"`
	// kube_controller_manager_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kube_controller_manager_config TfCluster#kube_controller_manager_config}
	// Experimental.
	KubeControllerManagerConfig *TfCluster_KubeControllerManagerConfigProperty `field:"optional" json:"kubeControllerManagerConfig" yaml:"kubeControllerManagerConfig"`
	// kubernetes_network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kubernetes_network_config TfCluster#kubernetes_network_config}
	// Experimental.
	KubernetesNetworkConfig *TfCluster_KubernetesNetworkConfigProperty `field:"optional" json:"kubernetesNetworkConfig" yaml:"kubernetesNetworkConfig"`
	// kube_scheduler_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#kube_scheduler_config TfCluster#kube_scheduler_config}
	// Experimental.
	KubeSchedulerConfig *TfCluster_KubeSchedulerConfigProperty `field:"optional" json:"kubeSchedulerConfig" yaml:"kubeSchedulerConfig"`
	// outpost_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#outpost_config TfCluster#outpost_config}
	// Experimental.
	OutpostConfig *TfCluster_OutpostConfigProperty `field:"optional" json:"outpostConfig" yaml:"outpostConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#region TfCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// remote_network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#remote_network_config TfCluster#remote_network_config}
	// Experimental.
	RemoteNetworkConfig *TfCluster_RemoteNetworkConfigProperty `field:"optional" json:"remoteNetworkConfig" yaml:"remoteNetworkConfig"`
	// storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#storage_config TfCluster#storage_config}
	// Experimental.
	StorageConfig *TfCluster_StorageConfigProperty `field:"optional" json:"storageConfig" yaml:"storageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#tags TfCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#tags_all TfCluster#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#timeouts TfCluster#timeouts}
	// Experimental.
	Timeouts *TfCluster_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#upgrade_policy TfCluster#upgrade_policy}
	// Experimental.
	UpgradePolicy *TfCluster_UpgradePolicyProperty `field:"optional" json:"upgradePolicy" yaml:"upgradePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#version TfCluster#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// zonal_shift_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#zonal_shift_config TfCluster#zonal_shift_config}
	// Experimental.
	ZonalShiftConfig *TfCluster_ZonalShiftConfigProperty `field:"optional" json:"zonalShiftConfig" yaml:"zonalShiftConfig"`
}

