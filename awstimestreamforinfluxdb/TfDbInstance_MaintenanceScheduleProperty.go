package awstimestreamforinfluxdb


// Experimental.
type TfDbInstance_MaintenanceScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreaminfluxdb_db_instance#preferred_maintenance_window TfDbInstance#preferred_maintenance_window}.
	// Experimental.
	PreferredMaintenanceWindow *string `field:"required" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreaminfluxdb_db_instance#timezone TfDbInstance#timezone}.
	// Experimental.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

