package awsoracledatabaseaws


// Experimental.
type TfCloudVmCluster_DataCollectionOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#is_diagnostics_events_enabled TfCloudVmCluster#is_diagnostics_events_enabled}.
	// Experimental.
	IsDiagnosticsEventsEnabled interface{} `field:"required" json:"isDiagnosticsEventsEnabled" yaml:"isDiagnosticsEventsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#is_health_monitoring_enabled TfCloudVmCluster#is_health_monitoring_enabled}.
	// Experimental.
	IsHealthMonitoringEnabled interface{} `field:"required" json:"isHealthMonitoringEnabled" yaml:"isHealthMonitoringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#is_incident_logs_enabled TfCloudVmCluster#is_incident_logs_enabled}.
	// Experimental.
	IsIncidentLogsEnabled interface{} `field:"required" json:"isIncidentLogsEnabled" yaml:"isIncidentLogsEnabled"`
}

