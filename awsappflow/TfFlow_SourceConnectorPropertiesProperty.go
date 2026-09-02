package awsappflow


// Experimental.
type TfFlow_SourceConnectorPropertiesProperty struct {
	// amplitude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#amplitude TfFlow#amplitude}
	// Experimental.
	Amplitude *TfFlow_AmplitudeProperty `field:"optional" json:"amplitude" yaml:"amplitude"`
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#custom_connector TfFlow#custom_connector}
	// Experimental.
	CustomConnector *TfFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty `field:"optional" json:"customConnector" yaml:"customConnector"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#datadog TfFlow#datadog}
	// Experimental.
	Datadog *TfFlow_DatadogProperty `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#dynatrace TfFlow#dynatrace}
	// Experimental.
	Dynatrace *TfFlow_DynatraceProperty `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// google_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#google_analytics TfFlow#google_analytics}
	// Experimental.
	GoogleAnalytics *TfFlow_GoogleAnalyticsProperty `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// infor_nexus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#infor_nexus TfFlow#infor_nexus}
	// Experimental.
	InforNexus *TfFlow_InforNexusProperty `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#marketo TfFlow#marketo}
	// Experimental.
	Marketo *TfFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty `field:"optional" json:"marketo" yaml:"marketo"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#s3 TfFlow#s3}
	// Experimental.
	S3 *TfFlow_SourceFlowConfigSourceConnectorPropertiesS3Property `field:"optional" json:"s3" yaml:"s3"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#salesforce TfFlow#salesforce}
	// Experimental.
	Salesforce *TfFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#sapo_data TfFlow#sapo_data}
	// Experimental.
	SapoData *TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty `field:"optional" json:"sapoData" yaml:"sapoData"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#service_now TfFlow#service_now}
	// Experimental.
	ServiceNow *TfFlow_ServiceNowProperty `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// singular block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#singular TfFlow#singular}
	// Experimental.
	Singular *TfFlow_SingularProperty `field:"optional" json:"singular" yaml:"singular"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#slack TfFlow#slack}
	// Experimental.
	Slack *TfFlow_SlackProperty `field:"optional" json:"slack" yaml:"slack"`
	// trendmicro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#trendmicro TfFlow#trendmicro}
	// Experimental.
	Trendmicro *TfFlow_TrendmicroProperty `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// veeva block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#veeva TfFlow#veeva}
	// Experimental.
	Veeva *TfFlow_VeevaProperty `field:"optional" json:"veeva" yaml:"veeva"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#zendesk TfFlow#zendesk}
	// Experimental.
	Zendesk *TfFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty `field:"optional" json:"zendesk" yaml:"zendesk"`
}

