package awsappflow


// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfilePropertiesProperty struct {
	// amplitude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#amplitude AwsAppflowConnectorProfile#amplitude}
	// Experimental.
	Amplitude *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty `field:"optional" json:"amplitude" yaml:"amplitude"`
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#custom_connector AwsAppflowConnectorProfile#custom_connector}
	// Experimental.
	CustomConnector *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty `field:"optional" json:"customConnector" yaml:"customConnector"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#datadog AwsAppflowConnectorProfile#datadog}
	// Experimental.
	Datadog *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#dynatrace AwsAppflowConnectorProfile#dynatrace}
	// Experimental.
	Dynatrace *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// google_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#google_analytics AwsAppflowConnectorProfile#google_analytics}
	// Experimental.
	GoogleAnalytics *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// honeycode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#honeycode AwsAppflowConnectorProfile#honeycode}
	// Experimental.
	Honeycode *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty `field:"optional" json:"honeycode" yaml:"honeycode"`
	// infor_nexus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#infor_nexus AwsAppflowConnectorProfile#infor_nexus}
	// Experimental.
	InforNexus *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#marketo AwsAppflowConnectorProfile#marketo}
	// Experimental.
	Marketo *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty `field:"optional" json:"marketo" yaml:"marketo"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#redshift AwsAppflowConnectorProfile#redshift}
	// Experimental.
	Redshift *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty `field:"optional" json:"redshift" yaml:"redshift"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#salesforce AwsAppflowConnectorProfile#salesforce}
	// Experimental.
	Salesforce *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#sapo_data AwsAppflowConnectorProfile#sapo_data}
	// Experimental.
	SapoData *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty `field:"optional" json:"sapoData" yaml:"sapoData"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#service_now AwsAppflowConnectorProfile#service_now}
	// Experimental.
	ServiceNow *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// singular block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#singular AwsAppflowConnectorProfile#singular}
	// Experimental.
	Singular *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty `field:"optional" json:"singular" yaml:"singular"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#slack AwsAppflowConnectorProfile#slack}
	// Experimental.
	Slack *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty `field:"optional" json:"slack" yaml:"slack"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#snowflake AwsAppflowConnectorProfile#snowflake}
	// Experimental.
	Snowflake *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty `field:"optional" json:"snowflake" yaml:"snowflake"`
	// trendmicro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#trendmicro AwsAppflowConnectorProfile#trendmicro}
	// Experimental.
	Trendmicro *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// veeva block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#veeva AwsAppflowConnectorProfile#veeva}
	// Experimental.
	Veeva *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty `field:"optional" json:"veeva" yaml:"veeva"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#zendesk AwsAppflowConnectorProfile#zendesk}
	// Experimental.
	Zendesk *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty `field:"optional" json:"zendesk" yaml:"zendesk"`
}

