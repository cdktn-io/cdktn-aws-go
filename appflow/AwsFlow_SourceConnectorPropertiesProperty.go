package appflow


// Experimental.
type AwsFlow_SourceConnectorPropertiesProperty struct {
	// amplitude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#amplitude AwsFlow#amplitude}
	// Experimental.
	Amplitude *AwsFlow_AmplitudeProperty `field:"optional" json:"amplitude" yaml:"amplitude"`
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#custom_connector AwsFlow#custom_connector}
	// Experimental.
	CustomConnector *AwsFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty `field:"optional" json:"customConnector" yaml:"customConnector"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#datadog AwsFlow#datadog}
	// Experimental.
	Datadog *AwsFlow_DatadogProperty `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#dynatrace AwsFlow#dynatrace}
	// Experimental.
	Dynatrace *AwsFlow_DynatraceProperty `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// google_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#google_analytics AwsFlow#google_analytics}
	// Experimental.
	GoogleAnalytics *AwsFlow_GoogleAnalyticsProperty `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// infor_nexus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#infor_nexus AwsFlow#infor_nexus}
	// Experimental.
	InforNexus *AwsFlow_InforNexusProperty `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#marketo AwsFlow#marketo}
	// Experimental.
	Marketo *AwsFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty `field:"optional" json:"marketo" yaml:"marketo"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#s3 AwsFlow#s3}
	// Experimental.
	S3 *AwsFlow_SourceFlowConfigSourceConnectorPropertiesS3Property `field:"optional" json:"s3" yaml:"s3"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#salesforce AwsFlow#salesforce}
	// Experimental.
	Salesforce *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#sapo_data AwsFlow#sapo_data}
	// Experimental.
	SapoData *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty `field:"optional" json:"sapoData" yaml:"sapoData"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#service_now AwsFlow#service_now}
	// Experimental.
	ServiceNow *AwsFlow_ServiceNowProperty `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// singular block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#singular AwsFlow#singular}
	// Experimental.
	Singular *AwsFlow_SingularProperty `field:"optional" json:"singular" yaml:"singular"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#slack AwsFlow#slack}
	// Experimental.
	Slack *AwsFlow_SlackProperty `field:"optional" json:"slack" yaml:"slack"`
	// trendmicro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#trendmicro AwsFlow#trendmicro}
	// Experimental.
	Trendmicro *AwsFlow_TrendmicroProperty `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// veeva block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#veeva AwsFlow#veeva}
	// Experimental.
	Veeva *AwsFlow_VeevaProperty `field:"optional" json:"veeva" yaml:"veeva"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#zendesk AwsFlow#zendesk}
	// Experimental.
	Zendesk *AwsFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty `field:"optional" json:"zendesk" yaml:"zendesk"`
}

