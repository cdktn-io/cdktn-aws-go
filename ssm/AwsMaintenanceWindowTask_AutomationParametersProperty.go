package ssm


// Experimental.
type AwsMaintenanceWindowTask_AutomationParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#document_version AwsMaintenanceWindowTask#document_version}.
	// Experimental.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_maintenance_window_task#parameter AwsMaintenanceWindowTask#parameter}
	// Experimental.
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

