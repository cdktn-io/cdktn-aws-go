package awstimestreamforinfluxdb


// Experimental.
type TfDbCluster_MaintenanceScheduleProperty struct {
	// The preferred maintenance window in the format ddd:HH:MM-ddd:HH:MM.
	//
	// Day must be one of Mon, Tue, Wed, Thu, Fri, Sat, or Sun. Provide an empty
	// 								string to let the system choose a window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreaminfluxdb_db_cluster#preferred_maintenance_window TfDbCluster#preferred_maintenance_window}
	// Experimental.
	PreferredMaintenanceWindow *string `field:"required" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The IANA timezone identifier for the maintenance window. For 								example, America/New_York or UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreaminfluxdb_db_cluster#timezone TfDbCluster#timezone}
	// Experimental.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

