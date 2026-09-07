package elasticsearch


// Experimental.
type AwsDomain_SnapshotOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#automated_snapshot_start_hour AwsDomain#automated_snapshot_start_hour}.
	// Experimental.
	AutomatedSnapshotStartHour *float64 `field:"required" json:"automatedSnapshotStartHour" yaml:"automatedSnapshotStartHour"`
}

