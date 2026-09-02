package awsappflow


// Experimental.
type TfFlow_DestinationConnectorPropertiesProperty struct {
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#custom_connector TfFlow#custom_connector}
	// Experimental.
	CustomConnector *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty `field:"optional" json:"customConnector" yaml:"customConnector"`
	// customer_profiles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#customer_profiles TfFlow#customer_profiles}
	// Experimental.
	CustomerProfiles *TfFlow_CustomerProfilesProperty `field:"optional" json:"customerProfiles" yaml:"customerProfiles"`
	// event_bridge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#event_bridge TfFlow#event_bridge}
	// Experimental.
	EventBridge *TfFlow_EventBridgeProperty `field:"optional" json:"eventBridge" yaml:"eventBridge"`
	// honeycode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#honeycode TfFlow#honeycode}
	// Experimental.
	Honeycode *TfFlow_HoneycodeProperty `field:"optional" json:"honeycode" yaml:"honeycode"`
	// lookout_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#lookout_metrics TfFlow#lookout_metrics}
	// Experimental.
	LookoutMetrics *TfFlow_LookoutMetricsProperty `field:"optional" json:"lookoutMetrics" yaml:"lookoutMetrics"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#marketo TfFlow#marketo}
	// Experimental.
	Marketo *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty `field:"optional" json:"marketo" yaml:"marketo"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#redshift TfFlow#redshift}
	// Experimental.
	Redshift *TfFlow_RedshiftProperty `field:"optional" json:"redshift" yaml:"redshift"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#s3 TfFlow#s3}
	// Experimental.
	S3 *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property `field:"optional" json:"s3" yaml:"s3"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#salesforce TfFlow#salesforce}
	// Experimental.
	Salesforce *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#sapo_data TfFlow#sapo_data}
	// Experimental.
	SapoData *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty `field:"optional" json:"sapoData" yaml:"sapoData"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#snowflake TfFlow#snowflake}
	// Experimental.
	Snowflake *TfFlow_SnowflakeProperty `field:"optional" json:"snowflake" yaml:"snowflake"`
	// upsolver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#upsolver TfFlow#upsolver}
	// Experimental.
	Upsolver *TfFlow_UpsolverProperty `field:"optional" json:"upsolver" yaml:"upsolver"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#zendesk TfFlow#zendesk}
	// Experimental.
	Zendesk *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty `field:"optional" json:"zendesk" yaml:"zendesk"`
}

