package elasticsearch


// Experimental.
type AwsDomain_ClusterConfigProperty struct {
	// cold_storage_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#cold_storage_options AwsDomain#cold_storage_options}
	// Experimental.
	ColdStorageOptions *AwsDomain_ColdStorageOptionsProperty `field:"optional" json:"coldStorageOptions" yaml:"coldStorageOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#dedicated_master_count AwsDomain#dedicated_master_count}.
	// Experimental.
	DedicatedMasterCount *float64 `field:"optional" json:"dedicatedMasterCount" yaml:"dedicatedMasterCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#dedicated_master_enabled AwsDomain#dedicated_master_enabled}.
	// Experimental.
	DedicatedMasterEnabled interface{} `field:"optional" json:"dedicatedMasterEnabled" yaml:"dedicatedMasterEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#dedicated_master_type AwsDomain#dedicated_master_type}.
	// Experimental.
	DedicatedMasterType *string `field:"optional" json:"dedicatedMasterType" yaml:"dedicatedMasterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#instance_count AwsDomain#instance_count}.
	// Experimental.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#instance_type AwsDomain#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#warm_count AwsDomain#warm_count}.
	// Experimental.
	WarmCount *float64 `field:"optional" json:"warmCount" yaml:"warmCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#warm_enabled AwsDomain#warm_enabled}.
	// Experimental.
	WarmEnabled interface{} `field:"optional" json:"warmEnabled" yaml:"warmEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#warm_type AwsDomain#warm_type}.
	// Experimental.
	WarmType *string `field:"optional" json:"warmType" yaml:"warmType"`
	// zone_awareness_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#zone_awareness_config AwsDomain#zone_awareness_config}
	// Experimental.
	ZoneAwarenessConfig *AwsDomain_ZoneAwarenessConfigProperty `field:"optional" json:"zoneAwarenessConfig" yaml:"zoneAwarenessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#zone_awareness_enabled AwsDomain#zone_awareness_enabled}.
	// Experimental.
	ZoneAwarenessEnabled interface{} `field:"optional" json:"zoneAwarenessEnabled" yaml:"zoneAwarenessEnabled"`
}

