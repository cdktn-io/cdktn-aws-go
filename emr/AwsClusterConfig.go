package emr

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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#name AwsCluster#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#release_label AwsCluster#release_label}.
	// Experimental.
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#service_role AwsCluster#service_role}.
	// Experimental.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#additional_info AwsCluster#additional_info}.
	// Experimental.
	AdditionalInfo *string `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#applications AwsCluster#applications}.
	// Experimental.
	Applications *[]*string `field:"optional" json:"applications" yaml:"applications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#autoscaling_role AwsCluster#autoscaling_role}.
	// Experimental.
	AutoscalingRole *string `field:"optional" json:"autoscalingRole" yaml:"autoscalingRole"`
	// auto_termination_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#auto_termination_policy AwsCluster#auto_termination_policy}
	// Experimental.
	AutoTerminationPolicy *AwsCluster_AutoTerminationPolicyProperty `field:"optional" json:"autoTerminationPolicy" yaml:"autoTerminationPolicy"`
	// bootstrap_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#bootstrap_action AwsCluster#bootstrap_action}
	// Experimental.
	BootstrapAction interface{} `field:"optional" json:"bootstrapAction" yaml:"bootstrapAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#configurations AwsCluster#configurations}.
	// Experimental.
	Configurations *string `field:"optional" json:"configurations" yaml:"configurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#configurations_json AwsCluster#configurations_json}.
	// Experimental.
	ConfigurationsJson *string `field:"optional" json:"configurationsJson" yaml:"configurationsJson"`
	// core_instance_fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#core_instance_fleet AwsCluster#core_instance_fleet}
	// Experimental.
	CoreInstanceFleet *AwsCluster_CoreInstanceFleetProperty `field:"optional" json:"coreInstanceFleet" yaml:"coreInstanceFleet"`
	// core_instance_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#core_instance_group AwsCluster#core_instance_group}
	// Experimental.
	CoreInstanceGroup *AwsCluster_CoreInstanceGroupProperty `field:"optional" json:"coreInstanceGroup" yaml:"coreInstanceGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#custom_ami_id AwsCluster#custom_ami_id}.
	// Experimental.
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#ebs_root_volume_size AwsCluster#ebs_root_volume_size}.
	// Experimental.
	EbsRootVolumeSize *float64 `field:"optional" json:"ebsRootVolumeSize" yaml:"ebsRootVolumeSize"`
	// ec2_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#ec2_attributes AwsCluster#ec2_attributes}
	// Experimental.
	Ec2Attributes *AwsCluster_Ec2AttributesProperty `field:"optional" json:"ec2Attributes" yaml:"ec2Attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#id AwsCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#keep_job_flow_alive_when_no_steps AwsCluster#keep_job_flow_alive_when_no_steps}.
	// Experimental.
	KeepJobFlowAliveWhenNoSteps interface{} `field:"optional" json:"keepJobFlowAliveWhenNoSteps" yaml:"keepJobFlowAliveWhenNoSteps"`
	// kerberos_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#kerberos_attributes AwsCluster#kerberos_attributes}
	// Experimental.
	KerberosAttributes *AwsCluster_KerberosAttributesProperty `field:"optional" json:"kerberosAttributes" yaml:"kerberosAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#list_steps_states AwsCluster#list_steps_states}.
	// Experimental.
	ListStepsStates *[]*string `field:"optional" json:"listStepsStates" yaml:"listStepsStates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#log_encryption_kms_key_id AwsCluster#log_encryption_kms_key_id}.
	// Experimental.
	LogEncryptionKmsKeyId *string `field:"optional" json:"logEncryptionKmsKeyId" yaml:"logEncryptionKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#log_uri AwsCluster#log_uri}.
	// Experimental.
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// master_instance_fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#master_instance_fleet AwsCluster#master_instance_fleet}
	// Experimental.
	MasterInstanceFleet *AwsCluster_MasterInstanceFleetProperty `field:"optional" json:"masterInstanceFleet" yaml:"masterInstanceFleet"`
	// master_instance_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#master_instance_group AwsCluster#master_instance_group}
	// Experimental.
	MasterInstanceGroup *AwsCluster_MasterInstanceGroupProperty `field:"optional" json:"masterInstanceGroup" yaml:"masterInstanceGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#os_release_label AwsCluster#os_release_label}.
	// Experimental.
	OsReleaseLabel *string `field:"optional" json:"osReleaseLabel" yaml:"osReleaseLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#placement_group_config AwsCluster#placement_group_config}.
	// Experimental.
	PlacementGroupConfig interface{} `field:"optional" json:"placementGroupConfig" yaml:"placementGroupConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#region AwsCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#scale_down_behavior AwsCluster#scale_down_behavior}.
	// Experimental.
	ScaleDownBehavior *string `field:"optional" json:"scaleDownBehavior" yaml:"scaleDownBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#security_configuration AwsCluster#security_configuration}.
	// Experimental.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#step AwsCluster#step}.
	// Experimental.
	Step interface{} `field:"optional" json:"step" yaml:"step"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#step_concurrency_level AwsCluster#step_concurrency_level}.
	// Experimental.
	StepConcurrencyLevel *float64 `field:"optional" json:"stepConcurrencyLevel" yaml:"stepConcurrencyLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#tags AwsCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#tags_all AwsCluster#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#termination_protection AwsCluster#termination_protection}.
	// Experimental.
	TerminationProtection interface{} `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#unhealthy_node_replacement AwsCluster#unhealthy_node_replacement}.
	// Experimental.
	UnhealthyNodeReplacement interface{} `field:"optional" json:"unhealthyNodeReplacement" yaml:"unhealthyNodeReplacement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#visible_to_all_users AwsCluster#visible_to_all_users}.
	// Experimental.
	VisibleToAllUsers interface{} `field:"optional" json:"visibleToAllUsers" yaml:"visibleToAllUsers"`
}

