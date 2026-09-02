package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Amplitude() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudePropertyOutputReference
	// Experimental.
	AmplitudeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty
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
	CustomConnector() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty
	// Experimental.
	Datadog() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogPropertyOutputReference
	// Experimental.
	DatadogInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty
	// Experimental.
	Dynatrace() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatracePropertyOutputReference
	// Experimental.
	DynatraceInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GoogleAnalytics() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsPropertyOutputReference
	// Experimental.
	GoogleAnalyticsInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty
	// Experimental.
	Honeycode() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodePropertyOutputReference
	// Experimental.
	HoneycodeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty
	// Experimental.
	InforNexus() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusPropertyOutputReference
	// Experimental.
	InforNexusInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty
	// Experimental.
	InternalValue() *TfConnectorProfile_ConnectorProfilePropertiesProperty
	// Experimental.
	SetInternalValue(val *TfConnectorProfile_ConnectorProfilePropertiesProperty)
	// Experimental.
	Marketo() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty
	// Experimental.
	Redshift() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty
	// Experimental.
	Salesforce() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty
	// Experimental.
	SapoData() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty
	// Experimental.
	ServiceNow() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty
	// Experimental.
	Singular() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularPropertyOutputReference
	// Experimental.
	SingularInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty
	// Experimental.
	Slack() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackPropertyOutputReference
	// Experimental.
	SlackInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty
	// Experimental.
	Snowflake() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trendmicro() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroPropertyOutputReference
	// Experimental.
	TrendmicroInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty
	// Experimental.
	Veeva() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference
	// Experimental.
	VeevaInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty
	// Experimental.
	Zendesk() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty
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
	PutAmplitude(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty)
	// Experimental.
	PutCustomConnector(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty)
	// Experimental.
	PutDatadog(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty)
	// Experimental.
	PutDynatrace(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty)
	// Experimental.
	PutGoogleAnalytics(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty)
	// Experimental.
	PutHoneycode(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty)
	// Experimental.
	PutInforNexus(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty)
	// Experimental.
	PutMarketo(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty)
	// Experimental.
	PutRedshift(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty)
	// Experimental.
	PutSalesforce(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty)
	// Experimental.
	PutSapoData(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty)
	// Experimental.
	PutServiceNow(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty)
	// Experimental.
	PutSingular(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty)
	// Experimental.
	PutSlack(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty)
	// Experimental.
	PutSnowflake(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty)
	// Experimental.
	PutTrendmicro(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty)
	// Experimental.
	PutVeeva(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty)
	// Experimental.
	PutZendesk(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty)
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

// The jsii proxy struct for TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference
type jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Amplitude() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudePropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudePropertyOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) AmplitudeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) CustomConnector() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) CustomConnectorInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Datadog() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogPropertyOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) DatadogInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Dynatrace() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatracePropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatracePropertyOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) DynatraceInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GoogleAnalytics() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GoogleAnalyticsInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Honeycode() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodePropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodePropertyOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) HoneycodeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InforNexus() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusPropertyOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InforNexusInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InternalValue() *TfConnectorProfile_ConnectorProfilePropertiesProperty {
	var returns *TfConnectorProfile_ConnectorProfilePropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Marketo() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) MarketoInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Redshift() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) RedshiftInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Salesforce() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SalesforceInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SapoData() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SapoDataInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ServiceNow() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ServiceNowInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Singular() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularPropertyOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SingularInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Slack() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackPropertyOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SlackInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Snowflake() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakePropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) SnowflakeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Trendmicro() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroPropertyOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) TrendmicroInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Veeva() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaPropertyOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) VeevaInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Zendesk() TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ZendeskInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfilePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference_Override(t TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfilePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetInternalValue(val *TfConnectorProfile_ConnectorProfilePropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutAmplitude(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesAmplitudeProperty) {
	if err := t.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutCustomConnector(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty) {
	if err := t.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutDatadog(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDatadogProperty) {
	if err := t.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDatadog",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutDynatrace(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesDynatraceProperty) {
	if err := t.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutGoogleAnalytics(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesGoogleAnalyticsProperty) {
	if err := t.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutHoneycode(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesHoneycodeProperty) {
	if err := t.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutInforNexus(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesInforNexusProperty) {
	if err := t.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutMarketo(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesMarketoProperty) {
	if err := t.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMarketo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutRedshift(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesRedshiftProperty) {
	if err := t.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedshift",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSalesforce(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSalesforceProperty) {
	if err := t.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSapoData(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty) {
	if err := t.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSapoData",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutServiceNow(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesServiceNowProperty) {
	if err := t.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSingular(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSingularProperty) {
	if err := t.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingular",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSlack(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSlackProperty) {
	if err := t.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSlack",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutSnowflake(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSnowflakeProperty) {
	if err := t.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutTrendmicro(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesTrendmicroProperty) {
	if err := t.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutVeeva(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesVeevaProperty) {
	if err := t.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVeeva",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) PutZendesk(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesZendeskProperty) {
	if err := t.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putZendesk",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		t,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		t,
		"resetDatadog",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		t,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		t,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		t,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		t,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		t,
		"resetMarketo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		t,
		"resetRedshift",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		t,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		t,
		"resetSapoData",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		t,
		"resetSingular",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		t,
		"resetSlack",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		t,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		t,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		t,
		"resetVeeva",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		t,
		"resetZendesk",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfilePropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

