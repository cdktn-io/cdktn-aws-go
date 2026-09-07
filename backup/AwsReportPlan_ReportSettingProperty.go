package backup


// Experimental.
type AwsReportPlan_ReportSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_report_plan#report_template AwsReportPlan#report_template}.
	// Experimental.
	ReportTemplate *string `field:"required" json:"reportTemplate" yaml:"reportTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_report_plan#accounts AwsReportPlan#accounts}.
	// Experimental.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_report_plan#framework_arns AwsReportPlan#framework_arns}.
	// Experimental.
	FrameworkArns *[]*string `field:"optional" json:"frameworkArns" yaml:"frameworkArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_report_plan#number_of_frameworks AwsReportPlan#number_of_frameworks}.
	// Experimental.
	NumberOfFrameworks *float64 `field:"optional" json:"numberOfFrameworks" yaml:"numberOfFrameworks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_report_plan#organization_units AwsReportPlan#organization_units}.
	// Experimental.
	OrganizationUnits *[]*string `field:"optional" json:"organizationUnits" yaml:"organizationUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_report_plan#regions AwsReportPlan#regions}.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

