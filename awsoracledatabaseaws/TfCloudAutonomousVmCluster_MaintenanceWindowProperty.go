package awsoracledatabaseaws


// Experimental.
type TfCloudAutonomousVmCluster_MaintenanceWindowProperty struct {
	// The preference for the maintenance window scheduling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#preference TfCloudAutonomousVmCluster#preference}
	// Experimental.
	Preference *string `field:"required" json:"preference" yaml:"preference"`
	// The days of the week when maintenance can be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#days_of_week TfCloudAutonomousVmCluster#days_of_week}
	// Experimental.
	DaysOfWeek interface{} `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// The hours of the day when maintenance can be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#hours_of_day TfCloudAutonomousVmCluster#hours_of_day}
	// Experimental.
	HoursOfDay *[]*float64 `field:"optional" json:"hoursOfDay" yaml:"hoursOfDay"`
	// The lead time in weeks before the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#lead_time_in_weeks TfCloudAutonomousVmCluster#lead_time_in_weeks}
	// Experimental.
	LeadTimeInWeeks *float64 `field:"optional" json:"leadTimeInWeeks" yaml:"leadTimeInWeeks"`
	// The months when maintenance can be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#months TfCloudAutonomousVmCluster#months}
	// Experimental.
	Months interface{} `field:"optional" json:"months" yaml:"months"`
	// Indicates whether to skip release updates during maintenance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#weeks_of_month TfCloudAutonomousVmCluster#weeks_of_month}
	// Experimental.
	WeeksOfMonth *[]*float64 `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}

