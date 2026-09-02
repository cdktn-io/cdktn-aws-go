package awsmsk


// Experimental.
type TfCluster_CloudwatchLogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#enabled TfCluster#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#log_group TfCluster#log_group}.
	// Experimental.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

