package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileCredentialsProperty struct {
	// amplitude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#amplitude TfConnectorProfile#amplitude}
	// Experimental.
	Amplitude *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty `field:"optional" json:"amplitude" yaml:"amplitude"`
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#custom_connector TfConnectorProfile#custom_connector}
	// Experimental.
	CustomConnector *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty `field:"optional" json:"customConnector" yaml:"customConnector"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#datadog TfConnectorProfile#datadog}
	// Experimental.
	Datadog *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#dynatrace TfConnectorProfile#dynatrace}
	// Experimental.
	Dynatrace *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// google_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#google_analytics TfConnectorProfile#google_analytics}
	// Experimental.
	GoogleAnalytics *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// honeycode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#honeycode TfConnectorProfile#honeycode}
	// Experimental.
	Honeycode *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty `field:"optional" json:"honeycode" yaml:"honeycode"`
	// infor_nexus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#infor_nexus TfConnectorProfile#infor_nexus}
	// Experimental.
	InforNexus *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#marketo TfConnectorProfile#marketo}
	// Experimental.
	Marketo *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty `field:"optional" json:"marketo" yaml:"marketo"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#redshift TfConnectorProfile#redshift}
	// Experimental.
	Redshift *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty `field:"optional" json:"redshift" yaml:"redshift"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#salesforce TfConnectorProfile#salesforce}
	// Experimental.
	Salesforce *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#sapo_data TfConnectorProfile#sapo_data}
	// Experimental.
	SapoData *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty `field:"optional" json:"sapoData" yaml:"sapoData"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#service_now TfConnectorProfile#service_now}
	// Experimental.
	ServiceNow *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// singular block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#singular TfConnectorProfile#singular}
	// Experimental.
	Singular *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty `field:"optional" json:"singular" yaml:"singular"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#slack TfConnectorProfile#slack}
	// Experimental.
	Slack *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty `field:"optional" json:"slack" yaml:"slack"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#snowflake TfConnectorProfile#snowflake}
	// Experimental.
	Snowflake *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty `field:"optional" json:"snowflake" yaml:"snowflake"`
	// trendmicro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#trendmicro TfConnectorProfile#trendmicro}
	// Experimental.
	Trendmicro *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// veeva block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#veeva TfConnectorProfile#veeva}
	// Experimental.
	Veeva *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty `field:"optional" json:"veeva" yaml:"veeva"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#zendesk TfConnectorProfile#zendesk}
	// Experimental.
	Zendesk *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty `field:"optional" json:"zendesk" yaml:"zendesk"`
}

