package awsarcregionswitch


// Experimental.
type TfPlan_WorkflowStepEcsCapacityIncreaseConfigServiceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#cluster_arn TfPlan#cluster_arn}.
	// Experimental.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#service_arn TfPlan#service_arn}.
	// Experimental.
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#cross_account_role TfPlan#cross_account_role}.
	// Experimental.
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#external_id TfPlan#external_id}.
	// Experimental.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

