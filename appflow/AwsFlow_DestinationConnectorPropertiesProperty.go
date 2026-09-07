package appflow


// Experimental.
type AwsFlow_DestinationConnectorPropertiesProperty struct {
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#custom_connector AwsFlow#custom_connector}
	// Experimental.
	CustomConnector *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty `field:"optional" json:"customConnector" yaml:"customConnector"`
	// customer_profiles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#customer_profiles AwsFlow#customer_profiles}
	// Experimental.
	CustomerProfiles *AwsFlow_CustomerProfilesProperty `field:"optional" json:"customerProfiles" yaml:"customerProfiles"`
	// event_bridge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#event_bridge AwsFlow#event_bridge}
	// Experimental.
	EventBridge *AwsFlow_EventBridgeProperty `field:"optional" json:"eventBridge" yaml:"eventBridge"`
	// honeycode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#honeycode AwsFlow#honeycode}
	// Experimental.
	Honeycode *AwsFlow_HoneycodeProperty `field:"optional" json:"honeycode" yaml:"honeycode"`
	// lookout_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#lookout_metrics AwsFlow#lookout_metrics}
	// Experimental.
	LookoutMetrics *AwsFlow_LookoutMetricsProperty `field:"optional" json:"lookoutMetrics" yaml:"lookoutMetrics"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#marketo AwsFlow#marketo}
	// Experimental.
	Marketo *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty `field:"optional" json:"marketo" yaml:"marketo"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#redshift AwsFlow#redshift}
	// Experimental.
	Redshift *AwsFlow_RedshiftProperty `field:"optional" json:"redshift" yaml:"redshift"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#s3 AwsFlow#s3}
	// Experimental.
	S3 *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property `field:"optional" json:"s3" yaml:"s3"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#salesforce AwsFlow#salesforce}
	// Experimental.
	Salesforce *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#sapo_data AwsFlow#sapo_data}
	// Experimental.
	SapoData *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty `field:"optional" json:"sapoData" yaml:"sapoData"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#snowflake AwsFlow#snowflake}
	// Experimental.
	Snowflake *AwsFlow_SnowflakeProperty `field:"optional" json:"snowflake" yaml:"snowflake"`
	// upsolver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#upsolver AwsFlow#upsolver}
	// Experimental.
	Upsolver *AwsFlow_UpsolverProperty `field:"optional" json:"upsolver" yaml:"upsolver"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#zendesk AwsFlow#zendesk}
	// Experimental.
	Zendesk *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty `field:"optional" json:"zendesk" yaml:"zendesk"`
}

