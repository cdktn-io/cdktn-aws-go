package awsarcregionswitch


// Experimental.
type AwsArcregionswitchPlan_WorkflowStepParallelConfigStepProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#execution_block_type AwsArcregionswitchPlan#execution_block_type}.
	// Experimental.
	ExecutionBlockType *string `field:"required" json:"executionBlockType" yaml:"executionBlockType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#name AwsArcregionswitchPlan#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// arc_routing_control_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#arc_routing_control_config AwsArcregionswitchPlan#arc_routing_control_config}
	// Experimental.
	ArcRoutingControlConfig interface{} `field:"optional" json:"arcRoutingControlConfig" yaml:"arcRoutingControlConfig"`
	// aurora_provisioned_scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#aurora_provisioned_scaling_config AwsArcregionswitchPlan#aurora_provisioned_scaling_config}
	// Experimental.
	AuroraProvisionedScalingConfig interface{} `field:"optional" json:"auroraProvisionedScalingConfig" yaml:"auroraProvisionedScalingConfig"`
	// aurora_serverless_scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#aurora_serverless_scaling_config AwsArcregionswitchPlan#aurora_serverless_scaling_config}
	// Experimental.
	AuroraServerlessScalingConfig interface{} `field:"optional" json:"auroraServerlessScalingConfig" yaml:"auroraServerlessScalingConfig"`
	// custom_action_lambda_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#custom_action_lambda_config AwsArcregionswitchPlan#custom_action_lambda_config}
	// Experimental.
	CustomActionLambdaConfig interface{} `field:"optional" json:"customActionLambdaConfig" yaml:"customActionLambdaConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#description AwsArcregionswitchPlan#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// document_db_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#document_db_config AwsArcregionswitchPlan#document_db_config}
	// Experimental.
	DocumentDbConfig interface{} `field:"optional" json:"documentDbConfig" yaml:"documentDbConfig"`
	// ec2_asg_capacity_increase_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#ec2_asg_capacity_increase_config AwsArcregionswitchPlan#ec2_asg_capacity_increase_config}
	// Experimental.
	Ec2AsgCapacityIncreaseConfig interface{} `field:"optional" json:"ec2AsgCapacityIncreaseConfig" yaml:"ec2AsgCapacityIncreaseConfig"`
	// ecs_capacity_increase_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#ecs_capacity_increase_config AwsArcregionswitchPlan#ecs_capacity_increase_config}
	// Experimental.
	EcsCapacityIncreaseConfig interface{} `field:"optional" json:"ecsCapacityIncreaseConfig" yaml:"ecsCapacityIncreaseConfig"`
	// eks_resource_scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#eks_resource_scaling_config AwsArcregionswitchPlan#eks_resource_scaling_config}
	// Experimental.
	EksResourceScalingConfig interface{} `field:"optional" json:"eksResourceScalingConfig" yaml:"eksResourceScalingConfig"`
	// execution_approval_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#execution_approval_config AwsArcregionswitchPlan#execution_approval_config}
	// Experimental.
	ExecutionApprovalConfig interface{} `field:"optional" json:"executionApprovalConfig" yaml:"executionApprovalConfig"`
	// global_aurora_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#global_aurora_config AwsArcregionswitchPlan#global_aurora_config}
	// Experimental.
	GlobalAuroraConfig interface{} `field:"optional" json:"globalAuroraConfig" yaml:"globalAuroraConfig"`
	// lambda_event_source_mapping_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#lambda_event_source_mapping_config AwsArcregionswitchPlan#lambda_event_source_mapping_config}
	// Experimental.
	LambdaEventSourceMappingConfig interface{} `field:"optional" json:"lambdaEventSourceMappingConfig" yaml:"lambdaEventSourceMappingConfig"`
	// neptune_global_database_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#neptune_global_database_config AwsArcregionswitchPlan#neptune_global_database_config}
	// Experimental.
	NeptuneGlobalDatabaseConfig interface{} `field:"optional" json:"neptuneGlobalDatabaseConfig" yaml:"neptuneGlobalDatabaseConfig"`
	// rds_create_cross_region_read_replica_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#rds_create_cross_region_read_replica_config AwsArcregionswitchPlan#rds_create_cross_region_read_replica_config}
	// Experimental.
	RdsCreateCrossRegionReadReplicaConfig interface{} `field:"optional" json:"rdsCreateCrossRegionReadReplicaConfig" yaml:"rdsCreateCrossRegionReadReplicaConfig"`
	// rds_promote_read_replica_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#rds_promote_read_replica_config AwsArcregionswitchPlan#rds_promote_read_replica_config}
	// Experimental.
	RdsPromoteReadReplicaConfig interface{} `field:"optional" json:"rdsPromoteReadReplicaConfig" yaml:"rdsPromoteReadReplicaConfig"`
	// region_switch_plan_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#region_switch_plan_config AwsArcregionswitchPlan#region_switch_plan_config}
	// Experimental.
	RegionSwitchPlanConfig interface{} `field:"optional" json:"regionSwitchPlanConfig" yaml:"regionSwitchPlanConfig"`
	// route53_health_check_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#route53_health_check_config AwsArcregionswitchPlan#route53_health_check_config}
	// Experimental.
	Route53HealthCheckConfig interface{} `field:"optional" json:"route53HealthCheckConfig" yaml:"route53HealthCheckConfig"`
}

