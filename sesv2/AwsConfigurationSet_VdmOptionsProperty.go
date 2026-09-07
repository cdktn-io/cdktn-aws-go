package sesv2


// Experimental.
type AwsConfigurationSet_VdmOptionsProperty struct {
	// dashboard_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#dashboard_options AwsConfigurationSet#dashboard_options}
	// Experimental.
	DashboardOptions *AwsConfigurationSet_DashboardOptionsProperty `field:"optional" json:"dashboardOptions" yaml:"dashboardOptions"`
	// guardian_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#guardian_options AwsConfigurationSet#guardian_options}
	// Experimental.
	GuardianOptions *AwsConfigurationSet_GuardianOptionsProperty `field:"optional" json:"guardianOptions" yaml:"guardianOptions"`
}

