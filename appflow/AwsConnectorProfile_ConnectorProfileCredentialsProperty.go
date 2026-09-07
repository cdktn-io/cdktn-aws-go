package appflow


// Experimental.
type AwsConnectorProfile_ConnectorProfileCredentialsProperty struct {
	// amplitude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#amplitude AwsConnectorProfile#amplitude}
	// Experimental.
	Amplitude *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty `field:"optional" json:"amplitude" yaml:"amplitude"`
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#custom_connector AwsConnectorProfile#custom_connector}
	// Experimental.
	CustomConnector *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty `field:"optional" json:"customConnector" yaml:"customConnector"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#datadog AwsConnectorProfile#datadog}
	// Experimental.
	Datadog *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#dynatrace AwsConnectorProfile#dynatrace}
	// Experimental.
	Dynatrace *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// google_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#google_analytics AwsConnectorProfile#google_analytics}
	// Experimental.
	GoogleAnalytics *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// honeycode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#honeycode AwsConnectorProfile#honeycode}
	// Experimental.
	Honeycode *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty `field:"optional" json:"honeycode" yaml:"honeycode"`
	// infor_nexus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#infor_nexus AwsConnectorProfile#infor_nexus}
	// Experimental.
	InforNexus *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#marketo AwsConnectorProfile#marketo}
	// Experimental.
	Marketo *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty `field:"optional" json:"marketo" yaml:"marketo"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#redshift AwsConnectorProfile#redshift}
	// Experimental.
	Redshift *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty `field:"optional" json:"redshift" yaml:"redshift"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#salesforce AwsConnectorProfile#salesforce}
	// Experimental.
	Salesforce *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#sapo_data AwsConnectorProfile#sapo_data}
	// Experimental.
	SapoData *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty `field:"optional" json:"sapoData" yaml:"sapoData"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#service_now AwsConnectorProfile#service_now}
	// Experimental.
	ServiceNow *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// singular block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#singular AwsConnectorProfile#singular}
	// Experimental.
	Singular *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty `field:"optional" json:"singular" yaml:"singular"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#slack AwsConnectorProfile#slack}
	// Experimental.
	Slack *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty `field:"optional" json:"slack" yaml:"slack"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#snowflake AwsConnectorProfile#snowflake}
	// Experimental.
	Snowflake *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty `field:"optional" json:"snowflake" yaml:"snowflake"`
	// trendmicro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#trendmicro AwsConnectorProfile#trendmicro}
	// Experimental.
	Trendmicro *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// veeva block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#veeva AwsConnectorProfile#veeva}
	// Experimental.
	Veeva *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty `field:"optional" json:"veeva" yaml:"veeva"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#zendesk AwsConnectorProfile#zendesk}
	// Experimental.
	Zendesk *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty `field:"optional" json:"zendesk" yaml:"zendesk"`
}

