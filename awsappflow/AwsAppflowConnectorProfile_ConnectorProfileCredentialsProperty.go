package awsappflow


// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfileCredentialsProperty struct {
	// amplitude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#amplitude AwsAppflowConnectorProfile#amplitude}
	// Experimental.
	Amplitude *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty `field:"optional" json:"amplitude" yaml:"amplitude"`
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#custom_connector AwsAppflowConnectorProfile#custom_connector}
	// Experimental.
	CustomConnector *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty `field:"optional" json:"customConnector" yaml:"customConnector"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#datadog AwsAppflowConnectorProfile#datadog}
	// Experimental.
	Datadog *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#dynatrace AwsAppflowConnectorProfile#dynatrace}
	// Experimental.
	Dynatrace *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// google_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#google_analytics AwsAppflowConnectorProfile#google_analytics}
	// Experimental.
	GoogleAnalytics *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// honeycode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#honeycode AwsAppflowConnectorProfile#honeycode}
	// Experimental.
	Honeycode *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty `field:"optional" json:"honeycode" yaml:"honeycode"`
	// infor_nexus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#infor_nexus AwsAppflowConnectorProfile#infor_nexus}
	// Experimental.
	InforNexus *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#marketo AwsAppflowConnectorProfile#marketo}
	// Experimental.
	Marketo *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty `field:"optional" json:"marketo" yaml:"marketo"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#redshift AwsAppflowConnectorProfile#redshift}
	// Experimental.
	Redshift *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty `field:"optional" json:"redshift" yaml:"redshift"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#salesforce AwsAppflowConnectorProfile#salesforce}
	// Experimental.
	Salesforce *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#sapo_data AwsAppflowConnectorProfile#sapo_data}
	// Experimental.
	SapoData *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty `field:"optional" json:"sapoData" yaml:"sapoData"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#service_now AwsAppflowConnectorProfile#service_now}
	// Experimental.
	ServiceNow *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// singular block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#singular AwsAppflowConnectorProfile#singular}
	// Experimental.
	Singular *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty `field:"optional" json:"singular" yaml:"singular"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#slack AwsAppflowConnectorProfile#slack}
	// Experimental.
	Slack *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty `field:"optional" json:"slack" yaml:"slack"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#snowflake AwsAppflowConnectorProfile#snowflake}
	// Experimental.
	Snowflake *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty `field:"optional" json:"snowflake" yaml:"snowflake"`
	// trendmicro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#trendmicro AwsAppflowConnectorProfile#trendmicro}
	// Experimental.
	Trendmicro *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// veeva block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#veeva AwsAppflowConnectorProfile#veeva}
	// Experimental.
	Veeva *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty `field:"optional" json:"veeva" yaml:"veeva"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#zendesk AwsAppflowConnectorProfile#zendesk}
	// Experimental.
	Zendesk *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty `field:"optional" json:"zendesk" yaml:"zendesk"`
}

