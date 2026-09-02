package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfilePropertiesProperty struct {
	// amplitude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#amplitude TfConnectorProfile#amplitude}
	// Experimental.
	Amplitude *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty `field:"optional" json:"amplitude" yaml:"amplitude"`
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#custom_connector TfConnectorProfile#custom_connector}
	// Experimental.
	CustomConnector *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty `field:"optional" json:"customConnector" yaml:"customConnector"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#datadog TfConnectorProfile#datadog}
	// Experimental.
	Datadog *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#dynatrace TfConnectorProfile#dynatrace}
	// Experimental.
	Dynatrace *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// google_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#google_analytics TfConnectorProfile#google_analytics}
	// Experimental.
	GoogleAnalytics *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// honeycode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#honeycode TfConnectorProfile#honeycode}
	// Experimental.
	Honeycode *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty `field:"optional" json:"honeycode" yaml:"honeycode"`
	// infor_nexus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#infor_nexus TfConnectorProfile#infor_nexus}
	// Experimental.
	InforNexus *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#marketo TfConnectorProfile#marketo}
	// Experimental.
	Marketo *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty `field:"optional" json:"marketo" yaml:"marketo"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#redshift TfConnectorProfile#redshift}
	// Experimental.
	Redshift *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty `field:"optional" json:"redshift" yaml:"redshift"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#salesforce TfConnectorProfile#salesforce}
	// Experimental.
	Salesforce *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#sapo_data TfConnectorProfile#sapo_data}
	// Experimental.
	SapoData *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty `field:"optional" json:"sapoData" yaml:"sapoData"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#service_now TfConnectorProfile#service_now}
	// Experimental.
	ServiceNow *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// singular block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#singular TfConnectorProfile#singular}
	// Experimental.
	Singular *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty `field:"optional" json:"singular" yaml:"singular"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#slack TfConnectorProfile#slack}
	// Experimental.
	Slack *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty `field:"optional" json:"slack" yaml:"slack"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#snowflake TfConnectorProfile#snowflake}
	// Experimental.
	Snowflake *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty `field:"optional" json:"snowflake" yaml:"snowflake"`
	// trendmicro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#trendmicro TfConnectorProfile#trendmicro}
	// Experimental.
	Trendmicro *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// veeva block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#veeva TfConnectorProfile#veeva}
	// Experimental.
	Veeva *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty `field:"optional" json:"veeva" yaml:"veeva"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#zendesk TfConnectorProfile#zendesk}
	// Experimental.
	Zendesk *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty `field:"optional" json:"zendesk" yaml:"zendesk"`
}

