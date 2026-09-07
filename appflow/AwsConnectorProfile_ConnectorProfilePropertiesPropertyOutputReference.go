package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Amplitude() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudePropertyOutputReference
	// Experimental.
	AmplitudeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty
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
	CustomConnector() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty
	// Experimental.
	Datadog() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogPropertyOutputReference
	// Experimental.
	DatadogInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty
	// Experimental.
	Dynatrace() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatracePropertyOutputReference
	// Experimental.
	DynatraceInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GoogleAnalytics() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsPropertyOutputReference
	// Experimental.
	GoogleAnalyticsInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty
	// Experimental.
	Honeycode() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodePropertyOutputReference
	// Experimental.
	HoneycodeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty
	// Experimental.
	InforNexus() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusPropertyOutputReference
	// Experimental.
	InforNexusInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty
	// Experimental.
	InternalValue() *AwsConnectorProfile_ConnectorProfilePropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsConnectorProfile_ConnectorProfilePropertiesProperty)
	// Experimental.
	Marketo() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty
	// Experimental.
	Redshift() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty
	// Experimental.
	Salesforce() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty
	// Experimental.
	SapoData() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty
	// Experimental.
	ServiceNow() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty
	// Experimental.
	Singular() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularPropertyOutputReference
	// Experimental.
	SingularInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty
	// Experimental.
	Slack() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackPropertyOutputReference
	// Experimental.
	SlackInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty
	// Experimental.
	Snowflake() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trendmicro() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroPropertyOutputReference
	// Experimental.
	TrendmicroInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty
	// Experimental.
	Veeva() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference
	// Experimental.
	VeevaInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty
	// Experimental.
	Zendesk() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty
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
	PutAmplitude(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty)
	// Experimental.
	PutCustomConnector(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty)
	// Experimental.
	PutDatadog(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty)
	// Experimental.
	PutDynatrace(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty)
	// Experimental.
	PutGoogleAnalytics(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty)
	// Experimental.
	PutHoneycode(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty)
	// Experimental.
	PutInforNexus(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty)
	// Experimental.
	PutMarketo(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty)
	// Experimental.
	PutRedshift(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty)
	// Experimental.
	PutSalesforce(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty)
	// Experimental.
	PutSapoData(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty)
	// Experimental.
	PutServiceNow(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty)
	// Experimental.
	PutSingular(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty)
	// Experimental.
	PutSlack(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty)
	// Experimental.
	PutSnowflake(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty)
	// Experimental.
	PutTrendmicro(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty)
	// Experimental.
	PutVeeva(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty)
	// Experimental.
	PutZendesk(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty)
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

// The jsii proxy struct for AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference
type jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Amplitude() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudePropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudePropertyOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) AmplitudeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) CustomConnector() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) CustomConnectorInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Datadog() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogPropertyOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) DatadogInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Dynatrace() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatracePropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatracePropertyOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) DynatraceInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GoogleAnalytics() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GoogleAnalyticsInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Honeycode() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodePropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodePropertyOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) HoneycodeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InforNexus() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusPropertyOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InforNexusInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InternalValue() *AwsConnectorProfile_ConnectorProfilePropertiesProperty {
	var returns *AwsConnectorProfile_ConnectorProfilePropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Marketo() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) MarketoInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Redshift() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) RedshiftInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Salesforce() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SalesforceInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SapoData() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SapoDataInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ServiceNow() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ServiceNowInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Singular() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularPropertyOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SingularInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Slack() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackPropertyOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SlackInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Snowflake() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakePropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SnowflakeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Trendmicro() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroPropertyOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) TrendmicroInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Veeva() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) VeevaInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Zendesk() AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ZendeskInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfilePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference_Override(a AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfilePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetInternalValue(val *AwsConnectorProfile_ConnectorProfilePropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutAmplitude(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty) {
	if err := a.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutCustomConnector(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutDatadog(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty) {
	if err := a.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutDynatrace(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty) {
	if err := a.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutGoogleAnalytics(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty) {
	if err := a.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutHoneycode(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty) {
	if err := a.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutInforNexus(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty) {
	if err := a.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutMarketo(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutRedshift(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty) {
	if err := a.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshift",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSalesforce(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSapoData(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutServiceNow(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSingular(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty) {
	if err := a.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSingular",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSlack(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty) {
	if err := a.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlack",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSnowflake(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty) {
	if err := a.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutTrendmicro(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty) {
	if err := a.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutVeeva(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty) {
	if err := a.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVeeva",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutZendesk(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		a,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

