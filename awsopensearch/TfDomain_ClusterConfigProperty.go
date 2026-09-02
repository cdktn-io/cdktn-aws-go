package awsopensearch


// Experimental.
type TfDomain_ClusterConfigProperty struct {
	// cold_storage_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#cold_storage_options TfDomain#cold_storage_options}
	// Experimental.
	ColdStorageOptions *TfDomain_ColdStorageOptionsProperty `field:"optional" json:"coldStorageOptions" yaml:"coldStorageOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#dedicated_master_count TfDomain#dedicated_master_count}.
	// Experimental.
	DedicatedMasterCount *float64 `field:"optional" json:"dedicatedMasterCount" yaml:"dedicatedMasterCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#dedicated_master_enabled TfDomain#dedicated_master_enabled}.
	// Experimental.
	DedicatedMasterEnabled interface{} `field:"optional" json:"dedicatedMasterEnabled" yaml:"dedicatedMasterEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#dedicated_master_type TfDomain#dedicated_master_type}.
	// Experimental.
	DedicatedMasterType *string `field:"optional" json:"dedicatedMasterType" yaml:"dedicatedMasterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#instance_count TfDomain#instance_count}.
	// Experimental.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#instance_type TfDomain#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#multi_az_with_standby_enabled TfDomain#multi_az_with_standby_enabled}.
	// Experimental.
	MultiAzWithStandbyEnabled interface{} `field:"optional" json:"multiAzWithStandbyEnabled" yaml:"multiAzWithStandbyEnabled"`
	// node_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#node_options TfDomain#node_options}
	// Experimental.
	NodeOptions interface{} `field:"optional" json:"nodeOptions" yaml:"nodeOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#warm_count TfDomain#warm_count}.
	// Experimental.
	WarmCount *float64 `field:"optional" json:"warmCount" yaml:"warmCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#warm_enabled TfDomain#warm_enabled}.
	// Experimental.
	WarmEnabled interface{} `field:"optional" json:"warmEnabled" yaml:"warmEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#warm_type TfDomain#warm_type}.
	// Experimental.
	WarmType *string `field:"optional" json:"warmType" yaml:"warmType"`
	// zone_awareness_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#zone_awareness_config TfDomain#zone_awareness_config}
	// Experimental.
	ZoneAwarenessConfig *TfDomain_ZoneAwarenessConfigProperty `field:"optional" json:"zoneAwarenessConfig" yaml:"zoneAwarenessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#zone_awareness_enabled TfDomain#zone_awareness_enabled}.
	// Experimental.
	ZoneAwarenessEnabled interface{} `field:"optional" json:"zoneAwarenessEnabled" yaml:"zoneAwarenessEnabled"`
}

