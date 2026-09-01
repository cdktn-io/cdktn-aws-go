package awsarcregionswitch


// Experimental.
type AwsArcregionswitchPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#capacity_monitoring_approach AwsArcregionswitchPlan#capacity_monitoring_approach}.
	// Experimental.
	CapacityMonitoringApproach *string `field:"required" json:"capacityMonitoringApproach" yaml:"capacityMonitoringApproach"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#target_percent AwsArcregionswitchPlan#target_percent}.
	// Experimental.
	TargetPercent *float64 `field:"required" json:"targetPercent" yaml:"targetPercent"`
	// eks_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#eks_clusters AwsArcregionswitchPlan#eks_clusters}
	// Experimental.
	EksClusters interface{} `field:"optional" json:"eksClusters" yaml:"eksClusters"`
	// kubernetes_resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#kubernetes_resource_type AwsArcregionswitchPlan#kubernetes_resource_type}
	// Experimental.
	KubernetesResourceType interface{} `field:"optional" json:"kubernetesResourceType" yaml:"kubernetesResourceType"`
	// scaling_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#scaling_resources AwsArcregionswitchPlan#scaling_resources}
	// Experimental.
	ScalingResources interface{} `field:"optional" json:"scalingResources" yaml:"scalingResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes AwsArcregionswitchPlan#timeout_minutes}.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#ungraceful AwsArcregionswitchPlan#ungraceful}
	// Experimental.
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

