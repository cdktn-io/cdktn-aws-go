package awssesv2


// Experimental.
type TfConfigurationSet_VdmOptionsProperty struct {
	// dashboard_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#dashboard_options TfConfigurationSet#dashboard_options}
	// Experimental.
	DashboardOptions *TfConfigurationSet_DashboardOptionsProperty `field:"optional" json:"dashboardOptions" yaml:"dashboardOptions"`
	// guardian_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#guardian_options TfConfigurationSet#guardian_options}
	// Experimental.
	GuardianOptions *TfConfigurationSet_GuardianOptionsProperty `field:"optional" json:"guardianOptions" yaml:"guardianOptions"`
}

