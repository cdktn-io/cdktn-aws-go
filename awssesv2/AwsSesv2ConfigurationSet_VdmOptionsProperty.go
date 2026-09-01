package awssesv2


// Experimental.
type AwsSesv2ConfigurationSet_VdmOptionsProperty struct {
	// dashboard_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#dashboard_options AwsSesv2ConfigurationSet#dashboard_options}
	// Experimental.
	DashboardOptions *AwsSesv2ConfigurationSet_DashboardOptionsProperty `field:"optional" json:"dashboardOptions" yaml:"dashboardOptions"`
	// guardian_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#guardian_options AwsSesv2ConfigurationSet#guardian_options}
	// Experimental.
	GuardianOptions *AwsSesv2ConfigurationSet_GuardianOptionsProperty `field:"optional" json:"guardianOptions" yaml:"guardianOptions"`
}

