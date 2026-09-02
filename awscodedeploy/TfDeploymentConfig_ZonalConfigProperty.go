package awscodedeploy


// Experimental.
type TfDeploymentConfig_ZonalConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_config#first_zone_monitor_duration_in_seconds TfDeploymentConfig#first_zone_monitor_duration_in_seconds}.
	// Experimental.
	FirstZoneMonitorDurationInSeconds *float64 `field:"optional" json:"firstZoneMonitorDurationInSeconds" yaml:"firstZoneMonitorDurationInSeconds"`
	// minimum_healthy_hosts_per_zone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_config#minimum_healthy_hosts_per_zone TfDeploymentConfig#minimum_healthy_hosts_per_zone}
	// Experimental.
	MinimumHealthyHostsPerZone *TfDeploymentConfig_MinimumHealthyHostsPerZoneProperty `field:"optional" json:"minimumHealthyHostsPerZone" yaml:"minimumHealthyHostsPerZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_config#monitor_duration_in_seconds TfDeploymentConfig#monitor_duration_in_seconds}.
	// Experimental.
	MonitorDurationInSeconds *float64 `field:"optional" json:"monitorDurationInSeconds" yaml:"monitorDurationInSeconds"`
}

