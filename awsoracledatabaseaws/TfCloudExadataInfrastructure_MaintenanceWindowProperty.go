package awsoracledatabaseaws


// Experimental.
type TfCloudExadataInfrastructure_MaintenanceWindowProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#custom_action_timeout_in_mins TfCloudExadataInfrastructure#custom_action_timeout_in_mins}.
	// Experimental.
	CustomActionTimeoutInMins *float64 `field:"required" json:"customActionTimeoutInMins" yaml:"customActionTimeoutInMins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#is_custom_action_timeout_enabled TfCloudExadataInfrastructure#is_custom_action_timeout_enabled}.
	// Experimental.
	IsCustomActionTimeoutEnabled interface{} `field:"required" json:"isCustomActionTimeoutEnabled" yaml:"isCustomActionTimeoutEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#patching_mode TfCloudExadataInfrastructure#patching_mode}.
	// Experimental.
	PatchingMode *string `field:"required" json:"patchingMode" yaml:"patchingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#preference TfCloudExadataInfrastructure#preference}.
	// Experimental.
	Preference *string `field:"required" json:"preference" yaml:"preference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#days_of_week TfCloudExadataInfrastructure#days_of_week}.
	// Experimental.
	DaysOfWeek interface{} `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#hours_of_day TfCloudExadataInfrastructure#hours_of_day}.
	// Experimental.
	HoursOfDay *[]*float64 `field:"optional" json:"hoursOfDay" yaml:"hoursOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#lead_time_in_weeks TfCloudExadataInfrastructure#lead_time_in_weeks}.
	// Experimental.
	LeadTimeInWeeks *float64 `field:"optional" json:"leadTimeInWeeks" yaml:"leadTimeInWeeks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#months TfCloudExadataInfrastructure#months}.
	// Experimental.
	Months interface{} `field:"optional" json:"months" yaml:"months"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#weeks_of_month TfCloudExadataInfrastructure#weeks_of_month}.
	// Experimental.
	WeeksOfMonth *[]*float64 `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}

