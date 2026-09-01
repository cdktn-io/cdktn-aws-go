package awsappflow


// Experimental.
type AwsAppflowFlow_SourceConnectorPropertiesProperty struct {
	// amplitude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#amplitude AwsAppflowFlow#amplitude}
	// Experimental.
	Amplitude *AwsAppflowFlow_AmplitudeProperty `field:"optional" json:"amplitude" yaml:"amplitude"`
	// custom_connector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#custom_connector AwsAppflowFlow#custom_connector}
	// Experimental.
	CustomConnector *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty `field:"optional" json:"customConnector" yaml:"customConnector"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#datadog AwsAppflowFlow#datadog}
	// Experimental.
	Datadog *AwsAppflowFlow_DatadogProperty `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#dynatrace AwsAppflowFlow#dynatrace}
	// Experimental.
	Dynatrace *AwsAppflowFlow_DynatraceProperty `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// google_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#google_analytics AwsAppflowFlow#google_analytics}
	// Experimental.
	GoogleAnalytics *AwsAppflowFlow_GoogleAnalyticsProperty `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// infor_nexus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#infor_nexus AwsAppflowFlow#infor_nexus}
	// Experimental.
	InforNexus *AwsAppflowFlow_InforNexusProperty `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// marketo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#marketo AwsAppflowFlow#marketo}
	// Experimental.
	Marketo *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty `field:"optional" json:"marketo" yaml:"marketo"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#s3 AwsAppflowFlow#s3}
	// Experimental.
	S3 *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesS3Property `field:"optional" json:"s3" yaml:"s3"`
	// salesforce block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#salesforce AwsAppflowFlow#salesforce}
	// Experimental.
	Salesforce *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty `field:"optional" json:"salesforce" yaml:"salesforce"`
	// sapo_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#sapo_data AwsAppflowFlow#sapo_data}
	// Experimental.
	SapoData *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty `field:"optional" json:"sapoData" yaml:"sapoData"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#service_now AwsAppflowFlow#service_now}
	// Experimental.
	ServiceNow *AwsAppflowFlow_ServiceNowProperty `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// singular block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#singular AwsAppflowFlow#singular}
	// Experimental.
	Singular *AwsAppflowFlow_SingularProperty `field:"optional" json:"singular" yaml:"singular"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#slack AwsAppflowFlow#slack}
	// Experimental.
	Slack *AwsAppflowFlow_SlackProperty `field:"optional" json:"slack" yaml:"slack"`
	// trendmicro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#trendmicro AwsAppflowFlow#trendmicro}
	// Experimental.
	Trendmicro *AwsAppflowFlow_TrendmicroProperty `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// veeva block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#veeva AwsAppflowFlow#veeva}
	// Experimental.
	Veeva *AwsAppflowFlow_VeevaProperty `field:"optional" json:"veeva" yaml:"veeva"`
	// zendesk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#zendesk AwsAppflowFlow#zendesk}
	// Experimental.
	Zendesk *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty `field:"optional" json:"zendesk" yaml:"zendesk"`
}

