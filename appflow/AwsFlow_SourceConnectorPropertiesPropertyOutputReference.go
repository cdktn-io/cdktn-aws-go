package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFlow_SourceConnectorPropertiesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Amplitude() AwsFlow_AmplitudePropertyOutputReference
	// Experimental.
	AmplitudeInput() *AwsFlow_AmplitudeProperty
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
	CustomConnector() AwsFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty
	// Experimental.
	Datadog() AwsFlow_DatadogPropertyOutputReference
	// Experimental.
	DatadogInput() *AwsFlow_DatadogProperty
	// Experimental.
	Dynatrace() AwsFlow_DynatracePropertyOutputReference
	// Experimental.
	DynatraceInput() *AwsFlow_DynatraceProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GoogleAnalytics() AwsFlow_GoogleAnalyticsPropertyOutputReference
	// Experimental.
	GoogleAnalyticsInput() *AwsFlow_GoogleAnalyticsProperty
	// Experimental.
	InforNexus() AwsFlow_InforNexusPropertyOutputReference
	// Experimental.
	InforNexusInput() *AwsFlow_InforNexusProperty
	// Experimental.
	InternalValue() *AwsFlow_SourceConnectorPropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsFlow_SourceConnectorPropertiesProperty)
	// Experimental.
	Marketo() AwsFlow_SourceFlowConfigSourceConnectorPropertiesMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty
	// Experimental.
	S3() AwsFlow_SourceFlowConfigSourceConnectorPropertiesS3PropertyOutputReference
	// Experimental.
	S3Input() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesS3Property
	// Experimental.
	Salesforce() AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty
	// Experimental.
	SapoData() AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty
	// Experimental.
	ServiceNow() AwsFlow_ServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *AwsFlow_ServiceNowProperty
	// Experimental.
	Singular() AwsFlow_SingularPropertyOutputReference
	// Experimental.
	SingularInput() *AwsFlow_SingularProperty
	// Experimental.
	Slack() AwsFlow_SlackPropertyOutputReference
	// Experimental.
	SlackInput() *AwsFlow_SlackProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trendmicro() AwsFlow_TrendmicroPropertyOutputReference
	// Experimental.
	TrendmicroInput() *AwsFlow_TrendmicroProperty
	// Experimental.
	Veeva() AwsFlow_VeevaPropertyOutputReference
	// Experimental.
	VeevaInput() *AwsFlow_VeevaProperty
	// Experimental.
	Zendesk() AwsFlow_SourceFlowConfigSourceConnectorPropertiesZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty
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
	PutAmplitude(value *AwsFlow_AmplitudeProperty)
	// Experimental.
	PutCustomConnector(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty)
	// Experimental.
	PutDatadog(value *AwsFlow_DatadogProperty)
	// Experimental.
	PutDynatrace(value *AwsFlow_DynatraceProperty)
	// Experimental.
	PutGoogleAnalytics(value *AwsFlow_GoogleAnalyticsProperty)
	// Experimental.
	PutInforNexus(value *AwsFlow_InforNexusProperty)
	// Experimental.
	PutMarketo(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty)
	// Experimental.
	PutS3(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesS3Property)
	// Experimental.
	PutSalesforce(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty)
	// Experimental.
	PutSapoData(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty)
	// Experimental.
	PutServiceNow(value *AwsFlow_ServiceNowProperty)
	// Experimental.
	PutSingular(value *AwsFlow_SingularProperty)
	// Experimental.
	PutSlack(value *AwsFlow_SlackProperty)
	// Experimental.
	PutTrendmicro(value *AwsFlow_TrendmicroProperty)
	// Experimental.
	PutVeeva(value *AwsFlow_VeevaProperty)
	// Experimental.
	PutZendesk(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty)
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
	ResetInforNexus()
	// Experimental.
	ResetMarketo()
	// Experimental.
	ResetS3()
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

// The jsii proxy struct for AwsFlow_SourceConnectorPropertiesPropertyOutputReference
type jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Amplitude() AwsFlow_AmplitudePropertyOutputReference {
	var returns AwsFlow_AmplitudePropertyOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) AmplitudeInput() *AwsFlow_AmplitudeProperty {
	var returns *AwsFlow_AmplitudeProperty
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) CustomConnector() AwsFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorPropertyOutputReference {
	var returns AwsFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) CustomConnectorInput() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty {
	var returns *AwsFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Datadog() AwsFlow_DatadogPropertyOutputReference {
	var returns AwsFlow_DatadogPropertyOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) DatadogInput() *AwsFlow_DatadogProperty {
	var returns *AwsFlow_DatadogProperty
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Dynatrace() AwsFlow_DynatracePropertyOutputReference {
	var returns AwsFlow_DynatracePropertyOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) DynatraceInput() *AwsFlow_DynatraceProperty {
	var returns *AwsFlow_DynatraceProperty
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) GoogleAnalytics() AwsFlow_GoogleAnalyticsPropertyOutputReference {
	var returns AwsFlow_GoogleAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) GoogleAnalyticsInput() *AwsFlow_GoogleAnalyticsProperty {
	var returns *AwsFlow_GoogleAnalyticsProperty
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) InforNexus() AwsFlow_InforNexusPropertyOutputReference {
	var returns AwsFlow_InforNexusPropertyOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) InforNexusInput() *AwsFlow_InforNexusProperty {
	var returns *AwsFlow_InforNexusProperty
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) InternalValue() *AwsFlow_SourceConnectorPropertiesProperty {
	var returns *AwsFlow_SourceConnectorPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Marketo() AwsFlow_SourceFlowConfigSourceConnectorPropertiesMarketoPropertyOutputReference {
	var returns AwsFlow_SourceFlowConfigSourceConnectorPropertiesMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) MarketoInput() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty {
	var returns *AwsFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) S3() AwsFlow_SourceFlowConfigSourceConnectorPropertiesS3PropertyOutputReference {
	var returns AwsFlow_SourceFlowConfigSourceConnectorPropertiesS3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) S3Input() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesS3Property {
	var returns *AwsFlow_SourceFlowConfigSourceConnectorPropertiesS3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Salesforce() AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference {
	var returns AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) SalesforceInput() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty {
	var returns *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) SapoData() AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference {
	var returns AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) SapoDataInput() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty {
	var returns *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ServiceNow() AwsFlow_ServiceNowPropertyOutputReference {
	var returns AwsFlow_ServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ServiceNowInput() *AwsFlow_ServiceNowProperty {
	var returns *AwsFlow_ServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Singular() AwsFlow_SingularPropertyOutputReference {
	var returns AwsFlow_SingularPropertyOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) SingularInput() *AwsFlow_SingularProperty {
	var returns *AwsFlow_SingularProperty
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Slack() AwsFlow_SlackPropertyOutputReference {
	var returns AwsFlow_SlackPropertyOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) SlackInput() *AwsFlow_SlackProperty {
	var returns *AwsFlow_SlackProperty
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Trendmicro() AwsFlow_TrendmicroPropertyOutputReference {
	var returns AwsFlow_TrendmicroPropertyOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) TrendmicroInput() *AwsFlow_TrendmicroProperty {
	var returns *AwsFlow_TrendmicroProperty
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Veeva() AwsFlow_VeevaPropertyOutputReference {
	var returns AwsFlow_VeevaPropertyOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) VeevaInput() *AwsFlow_VeevaProperty {
	var returns *AwsFlow_VeevaProperty
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Zendesk() AwsFlow_SourceFlowConfigSourceConnectorPropertiesZendeskPropertyOutputReference {
	var returns AwsFlow_SourceFlowConfigSourceConnectorPropertiesZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ZendeskInput() *AwsFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty {
	var returns *AwsFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFlow_SourceConnectorPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFlow_SourceConnectorPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFlow_SourceConnectorPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.SourceConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFlow_SourceConnectorPropertiesPropertyOutputReference_Override(a AwsFlow_SourceConnectorPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.SourceConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference)SetInternalValue(val *AwsFlow_SourceConnectorPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutAmplitude(value *AwsFlow_AmplitudeProperty) {
	if err := a.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutCustomConnector(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutDatadog(value *AwsFlow_DatadogProperty) {
	if err := a.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutDynatrace(value *AwsFlow_DynatraceProperty) {
	if err := a.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutGoogleAnalytics(value *AwsFlow_GoogleAnalyticsProperty) {
	if err := a.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutInforNexus(value *AwsFlow_InforNexusProperty) {
	if err := a.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutMarketo(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutS3(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesS3Property) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutSalesforce(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutSapoData(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutServiceNow(value *AwsFlow_ServiceNowProperty) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutSingular(value *AwsFlow_SingularProperty) {
	if err := a.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSingular",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutSlack(value *AwsFlow_SlackProperty) {
	if err := a.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlack",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutTrendmicro(value *AwsFlow_TrendmicroProperty) {
	if err := a.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutVeeva(value *AwsFlow_VeevaProperty) {
	if err := a.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVeeva",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) PutZendesk(value *AwsFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFlow_SourceConnectorPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

