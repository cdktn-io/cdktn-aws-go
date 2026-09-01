package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Amplitude() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudePropertyOutputReference
	// Experimental.
	AmplitudeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomConnector() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty
	// Experimental.
	Datadog() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogPropertyOutputReference
	// Experimental.
	DatadogInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty
	// Experimental.
	Dynatrace() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatracePropertyOutputReference
	// Experimental.
	DynatraceInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GoogleAnalytics() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsPropertyOutputReference
	// Experimental.
	GoogleAnalyticsInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty
	// Experimental.
	Honeycode() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodePropertyOutputReference
	// Experimental.
	HoneycodeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty
	// Experimental.
	InforNexus() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusPropertyOutputReference
	// Experimental.
	InforNexusInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty
	// Experimental.
	InternalValue() *AwsAppflowConnectorProfile_ConnectorProfilePropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsAppflowConnectorProfile_ConnectorProfilePropertiesProperty)
	// Experimental.
	Marketo() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty
	// Experimental.
	Redshift() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty
	// Experimental.
	Salesforce() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty
	// Experimental.
	SapoData() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty
	// Experimental.
	ServiceNow() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty
	// Experimental.
	Singular() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularPropertyOutputReference
	// Experimental.
	SingularInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty
	// Experimental.
	Slack() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackPropertyOutputReference
	// Experimental.
	SlackInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty
	// Experimental.
	Snowflake() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trendmicro() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroPropertyOutputReference
	// Experimental.
	TrendmicroInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty
	// Experimental.
	Veeva() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference
	// Experimental.
	VeevaInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty
	// Experimental.
	Zendesk() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutAmplitude(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty)
	// Experimental.
	PutCustomConnector(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty)
	// Experimental.
	PutDatadog(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty)
	// Experimental.
	PutDynatrace(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty)
	// Experimental.
	PutGoogleAnalytics(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty)
	// Experimental.
	PutHoneycode(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty)
	// Experimental.
	PutInforNexus(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty)
	// Experimental.
	PutMarketo(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty)
	// Experimental.
	PutRedshift(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty)
	// Experimental.
	PutSalesforce(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty)
	// Experimental.
	PutSapoData(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty)
	// Experimental.
	PutServiceNow(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty)
	// Experimental.
	PutSingular(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty)
	// Experimental.
	PutSlack(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty)
	// Experimental.
	PutSnowflake(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty)
	// Experimental.
	PutTrendmicro(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty)
	// Experimental.
	PutVeeva(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty)
	// Experimental.
	PutZendesk(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty)
	// Experimental.
	ResetAmplitude()
	// Experimental.
	ResetCustomConnector()
	// Experimental.
	ResetDatadog()
	// Experimental.
	ResetDynatrace()
	// Experimental.
	ResetGoogleAnalytics()
	// Experimental.
	ResetHoneycode()
	// Experimental.
	ResetInforNexus()
	// Experimental.
	ResetMarketo()
	// Experimental.
	ResetRedshift()
	// Experimental.
	ResetSalesforce()
	// Experimental.
	ResetSapoData()
	// Experimental.
	ResetServiceNow()
	// Experimental.
	ResetSingular()
	// Experimental.
	ResetSlack()
	// Experimental.
	ResetSnowflake()
	// Experimental.
	ResetTrendmicro()
	// Experimental.
	ResetVeeva()
	// Experimental.
	ResetZendesk()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference
type jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Amplitude() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudePropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudePropertyOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) AmplitudeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) CustomConnector() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) CustomConnectorInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Datadog() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogPropertyOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) DatadogInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Dynatrace() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatracePropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatracePropertyOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) DynatraceInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GoogleAnalytics() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GoogleAnalyticsInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Honeycode() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodePropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodePropertyOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) HoneycodeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InforNexus() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusPropertyOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InforNexusInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InternalValue() *AwsAppflowConnectorProfile_ConnectorProfilePropertiesProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfilePropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Marketo() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) MarketoInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Redshift() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) RedshiftInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Salesforce() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SalesforceInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SapoData() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SapoDataInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ServiceNow() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ServiceNowInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Singular() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularPropertyOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SingularInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Slack() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackPropertyOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SlackInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Snowflake() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakePropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SnowflakeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Trendmicro() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroPropertyOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) TrendmicroInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Veeva() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) VeevaInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Zendesk() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ZendeskInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowConnectorProfile.ConnectorProfilePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference_Override(a AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowConnectorProfile.ConnectorProfilePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetInternalValue(val *AwsAppflowConnectorProfile_ConnectorProfilePropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutAmplitude(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty) {
	if err := a.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutCustomConnector(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutDatadog(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty) {
	if err := a.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutDynatrace(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty) {
	if err := a.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutGoogleAnalytics(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty) {
	if err := a.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutHoneycode(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty) {
	if err := a.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutInforNexus(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty) {
	if err := a.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutMarketo(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutRedshift(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty) {
	if err := a.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshift",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSalesforce(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSapoData(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutServiceNow(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSingular(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty) {
	if err := a.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSingular",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSlack(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty) {
	if err := a.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlack",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSnowflake(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty) {
	if err := a.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutTrendmicro(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty) {
	if err := a.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutVeeva(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty) {
	if err := a.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVeeva",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutZendesk(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		a,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

