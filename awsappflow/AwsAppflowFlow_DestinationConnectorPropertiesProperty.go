package awsappflow


// Experimental.
type AwsAppflowFlow_DestinationConnectorPropertiesProperty struct {
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#custom_connector AwsAppflowFlow#custom_connector}
	// Experimental.
	CustomConnector *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty `field:"optional" json:"customConnector" yaml:"customConnector"`
	// customer_profiles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#customer_profiles AwsAppflowFlow#customer_profiles}
	// Experimental.
	CustomerProfiles *AwsAppflowFlow_CustomerProfilesProperty `field:"optional" json:"customerProfiles" yaml:"customerProfiles"`
	// event_bridge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#event_bridge AwsAppflowFlow#event_bridge}
	// Experimental.
	EventBridge *AwsAppflowFlow_EventBridgeProperty `field:"optional" json:"eventBridge" yaml:"eventBridge"`
	// honeycode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#honeycode AwsAppflowFlow#honeycode}
	// Experimental.
	Honeycode *AwsAppflowFlow_HoneycodeProperty `field:"optional" json:"honeycode" yaml:"honeycode"`
	// lookout_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#lookout_metrics AwsAppflowFlow#lookout_metrics}
	// Experimental.
	LookoutMetrics *AwsAppflowFlow_LookoutMetricsProperty `field:"optional" json:"lookoutMetrics" yaml:"lookoutMetrics"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#marketo AwsAppflowFlow#marketo}
	// Experimental.
	Marketo *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty `field:"optional" json:"marketo" yaml:"marketo"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#redshift AwsAppflowFlow#redshift}
	// Experimental.
	Redshift *AwsAppflowFlow_RedshiftProperty `field:"optional" json:"redshift" yaml:"redshift"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#s3 AwsAppflowFlow#s3}
	// Experimental.
	S3 *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property `field:"optional" json:"s3" yaml:"s3"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#salesforce AwsAppflowFlow#salesforce}
	// Experimental.
	Salesforce *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#sapo_data AwsAppflowFlow#sapo_data}
	// Experimental.
	SapoData *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty `field:"optional" json:"sapoData" yaml:"sapoData"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#snowflake AwsAppflowFlow#snowflake}
	// Experimental.
	Snowflake *AwsAppflowFlow_SnowflakeProperty `field:"optional" json:"snowflake" yaml:"snowflake"`
	// upsolver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#upsolver AwsAppflowFlow#upsolver}
	// Experimental.
	Upsolver *AwsAppflowFlow_UpsolverProperty `field:"optional" json:"upsolver" yaml:"upsolver"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#zendesk AwsAppflowFlow#zendesk}
	// Experimental.
	Zendesk *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty `field:"optional" json:"zendesk" yaml:"zendesk"`
}

