package awsemr

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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#name TfCluster#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#release_label TfCluster#release_label}.
	// Experimental.
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#service_role TfCluster#service_role}.
	// Experimental.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#additional_info TfCluster#additional_info}.
	// Experimental.
	AdditionalInfo *string `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#applications TfCluster#applications}.
	// Experimental.
	Applications *[]*string `field:"optional" json:"applications" yaml:"applications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#autoscaling_role TfCluster#autoscaling_role}.
	// Experimental.
	AutoscalingRole *string `field:"optional" json:"autoscalingRole" yaml:"autoscalingRole"`
	// auto_termination_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#auto_termination_policy TfCluster#auto_termination_policy}
	// Experimental.
	AutoTerminationPolicy *TfCluster_AutoTerminationPolicyProperty `field:"optional" json:"autoTerminationPolicy" yaml:"autoTerminationPolicy"`
	// bootstrap_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#bootstrap_action TfCluster#bootstrap_action}
	// Experimental.
	BootstrapAction interface{} `field:"optional" json:"bootstrapAction" yaml:"bootstrapAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#configurations TfCluster#configurations}.
	// Experimental.
	Configurations *string `field:"optional" json:"configurations" yaml:"configurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#configurations_json TfCluster#configurations_json}.
	// Experimental.
	ConfigurationsJson *string `field:"optional" json:"configurationsJson" yaml:"configurationsJson"`
	// core_instance_fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#core_instance_fleet TfCluster#core_instance_fleet}
	// Experimental.
	CoreInstanceFleet *TfCluster_CoreInstanceFleetProperty `field:"optional" json:"coreInstanceFleet" yaml:"coreInstanceFleet"`
	// core_instance_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#core_instance_group TfCluster#core_instance_group}
	// Experimental.
	CoreInstanceGroup *TfCluster_CoreInstanceGroupProperty `field:"optional" json:"coreInstanceGroup" yaml:"coreInstanceGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#custom_ami_id TfCluster#custom_ami_id}.
	// Experimental.
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#ebs_root_volume_size TfCluster#ebs_root_volume_size}.
	// Experimental.
	EbsRootVolumeSize *float64 `field:"optional" json:"ebsRootVolumeSize" yaml:"ebsRootVolumeSize"`
	// ec2_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#ec2_attributes TfCluster#ec2_attributes}
	// Experimental.
	Ec2Attributes *TfCluster_Ec2AttributesProperty `field:"optional" json:"ec2Attributes" yaml:"ec2Attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#id TfCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#keep_job_flow_alive_when_no_steps TfCluster#keep_job_flow_alive_when_no_steps}.
	// Experimental.
	KeepJobFlowAliveWhenNoSteps interface{} `field:"optional" json:"keepJobFlowAliveWhenNoSteps" yaml:"keepJobFlowAliveWhenNoSteps"`
	// kerberos_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#kerberos_attributes TfCluster#kerberos_attributes}
	// Experimental.
	KerberosAttributes *TfCluster_KerberosAttributesProperty `field:"optional" json:"kerberosAttributes" yaml:"kerberosAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#list_steps_states TfCluster#list_steps_states}.
	// Experimental.
	ListStepsStates *[]*string `field:"optional" json:"listStepsStates" yaml:"listStepsStates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#log_encryption_kms_key_id TfCluster#log_encryption_kms_key_id}.
	// Experimental.
	LogEncryptionKmsKeyId *string `field:"optional" json:"logEncryptionKmsKeyId" yaml:"logEncryptionKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#log_uri TfCluster#log_uri}.
	// Experimental.
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// master_instance_fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#master_instance_fleet TfCluster#master_instance_fleet}
	// Experimental.
	MasterInstanceFleet *TfCluster_MasterInstanceFleetProperty `field:"optional" json:"masterInstanceFleet" yaml:"masterInstanceFleet"`
	// master_instance_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#master_instance_group TfCluster#master_instance_group}
	// Experimental.
	MasterInstanceGroup *TfCluster_MasterInstanceGroupProperty `field:"optional" json:"masterInstanceGroup" yaml:"masterInstanceGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#os_release_label TfCluster#os_release_label}.
	// Experimental.
	OsReleaseLabel *string `field:"optional" json:"osReleaseLabel" yaml:"osReleaseLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#placement_group_config TfCluster#placement_group_config}.
	// Experimental.
	PlacementGroupConfig interface{} `field:"optional" json:"placementGroupConfig" yaml:"placementGroupConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#region TfCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#scale_down_behavior TfCluster#scale_down_behavior}.
	// Experimental.
	ScaleDownBehavior *string `field:"optional" json:"scaleDownBehavior" yaml:"scaleDownBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#security_configuration TfCluster#security_configuration}.
	// Experimental.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#step TfCluster#step}.
	// Experimental.
	Step interface{} `field:"optional" json:"step" yaml:"step"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#step_concurrency_level TfCluster#step_concurrency_level}.
	// Experimental.
	StepConcurrencyLevel *float64 `field:"optional" json:"stepConcurrencyLevel" yaml:"stepConcurrencyLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#tags TfCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#tags_all TfCluster#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#termination_protection TfCluster#termination_protection}.
	// Experimental.
	TerminationProtection interface{} `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#unhealthy_node_replacement TfCluster#unhealthy_node_replacement}.
	// Experimental.
	UnhealthyNodeReplacement interface{} `field:"optional" json:"unhealthyNodeReplacement" yaml:"unhealthyNodeReplacement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#visible_to_all_users TfCluster#visible_to_all_users}.
	// Experimental.
	VisibleToAllUsers interface{} `field:"optional" json:"visibleToAllUsers" yaml:"visibleToAllUsers"`
}

