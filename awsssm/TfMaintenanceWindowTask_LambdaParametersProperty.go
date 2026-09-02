package awsssm


// Experimental.
type TfMaintenanceWindowTask_LambdaParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#client_context TfMaintenanceWindowTask#client_context}.
	// Experimental.
	ClientContext *string `field:"optional" json:"clientContext" yaml:"clientContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#payload TfMaintenanceWindowTask#payload}.
	// Experimental.
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#qualifier TfMaintenanceWindowTask#qualifier}.
	// Experimental.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

