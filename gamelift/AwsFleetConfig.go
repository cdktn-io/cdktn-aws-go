package gamelift

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFleetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#ec2_instance_type AwsFleet#ec2_instance_type}.
	// Experimental.
	Ec2InstanceType *string `field:"required" json:"ec2InstanceType" yaml:"ec2InstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#name AwsFleet#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#build_id AwsFleet#build_id}.
	// Experimental.
	BuildId *string `field:"optional" json:"buildId" yaml:"buildId"`
	// certificate_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#certificate_configuration AwsFleet#certificate_configuration}
	// Experimental.
	CertificateConfiguration *AwsFleet_CertificateConfigurationProperty `field:"optional" json:"certificateConfiguration" yaml:"certificateConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#description AwsFleet#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// ec2_inbound_permission block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#ec2_inbound_permission AwsFleet#ec2_inbound_permission}
	// Experimental.
	Ec2InboundPermission interface{} `field:"optional" json:"ec2InboundPermission" yaml:"ec2InboundPermission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#fleet_type AwsFleet#fleet_type}.
	// Experimental.
	FleetType *string `field:"optional" json:"fleetType" yaml:"fleetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#id AwsFleet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#instance_role_arn AwsFleet#instance_role_arn}.
	// Experimental.
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#metric_groups AwsFleet#metric_groups}.
	// Experimental.
	MetricGroups *[]*string `field:"optional" json:"metricGroups" yaml:"metricGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#new_game_session_protection_policy AwsFleet#new_game_session_protection_policy}.
	// Experimental.
	NewGameSessionProtectionPolicy *string `field:"optional" json:"newGameSessionProtectionPolicy" yaml:"newGameSessionProtectionPolicy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#region AwsFleet#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// resource_creation_limit_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#resource_creation_limit_policy AwsFleet#resource_creation_limit_policy}
	// Experimental.
	ResourceCreationLimitPolicy *AwsFleet_ResourceCreationLimitPolicyProperty `field:"optional" json:"resourceCreationLimitPolicy" yaml:"resourceCreationLimitPolicy"`
	// runtime_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#runtime_configuration AwsFleet#runtime_configuration}
	// Experimental.
	RuntimeConfiguration *AwsFleet_RuntimeConfigurationProperty `field:"optional" json:"runtimeConfiguration" yaml:"runtimeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#script_id AwsFleet#script_id}.
	// Experimental.
	ScriptId *string `field:"optional" json:"scriptId" yaml:"scriptId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#tags AwsFleet#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#tags_all AwsFleet#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#timeouts AwsFleet#timeouts}
	// Experimental.
	Timeouts *AwsFleet_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

