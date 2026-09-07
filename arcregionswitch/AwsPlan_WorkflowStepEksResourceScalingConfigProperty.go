package arcregionswitch


// Experimental.
type AwsPlan_WorkflowStepEksResourceScalingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#capacity_monitoring_approach AwsPlan#capacity_monitoring_approach}.
	// Experimental.
	CapacityMonitoringApproach *string `field:"required" json:"capacityMonitoringApproach" yaml:"capacityMonitoringApproach"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#target_percent AwsPlan#target_percent}.
	// Experimental.
	TargetPercent *float64 `field:"required" json:"targetPercent" yaml:"targetPercent"`
	// eks_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#eks_clusters AwsPlan#eks_clusters}
	// Experimental.
	EksClusters interface{} `field:"optional" json:"eksClusters" yaml:"eksClusters"`
	// kubernetes_resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#kubernetes_resource_type AwsPlan#kubernetes_resource_type}
	// Experimental.
	KubernetesResourceType interface{} `field:"optional" json:"kubernetesResourceType" yaml:"kubernetesResourceType"`
	// scaling_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#scaling_resources AwsPlan#scaling_resources}
	// Experimental.
	ScalingResources interface{} `field:"optional" json:"scalingResources" yaml:"scalingResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes AwsPlan#timeout_minutes}.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#ungraceful AwsPlan#ungraceful}
	// Experimental.
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

