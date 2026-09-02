package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlow_SourceConnectorPropertiesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Amplitude() TfFlow_AmplitudePropertyOutputReference
	// Experimental.
	AmplitudeInput() *TfFlow_AmplitudeProperty
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
	CustomConnector() TfFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *TfFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty
	// Experimental.
	Datadog() TfFlow_DatadogPropertyOutputReference
	// Experimental.
	DatadogInput() *TfFlow_DatadogProperty
	// Experimental.
	Dynatrace() TfFlow_DynatracePropertyOutputReference
	// Experimental.
	DynatraceInput() *TfFlow_DynatraceProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GoogleAnalytics() TfFlow_GoogleAnalyticsPropertyOutputReference
	// Experimental.
	GoogleAnalyticsInput() *TfFlow_GoogleAnalyticsProperty
	// Experimental.
	InforNexus() TfFlow_InforNexusPropertyOutputReference
	// Experimental.
	InforNexusInput() *TfFlow_InforNexusProperty
	// Experimental.
	InternalValue() *TfFlow_SourceConnectorPropertiesProperty
	// Experimental.
	SetInternalValue(val *TfFlow_SourceConnectorPropertiesProperty)
	// Experimental.
	Marketo() TfFlow_SourceFlowConfigSourceConnectorPropertiesMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *TfFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty
	// Experimental.
	S3() TfFlow_SourceFlowConfigSourceConnectorPropertiesS3PropertyOutputReference
	// Experimental.
	S3Input() *TfFlow_SourceFlowConfigSourceConnectorPropertiesS3Property
	// Experimental.
	Salesforce() TfFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *TfFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty
	// Experimental.
	SapoData() TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty
	// Experimental.
	ServiceNow() TfFlow_ServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *TfFlow_ServiceNowProperty
	// Experimental.
	Singular() TfFlow_SingularPropertyOutputReference
	// Experimental.
	SingularInput() *TfFlow_SingularProperty
	// Experimental.
	Slack() TfFlow_SlackPropertyOutputReference
	// Experimental.
	SlackInput() *TfFlow_SlackProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trendmicro() TfFlow_TrendmicroPropertyOutputReference
	// Experimental.
	TrendmicroInput() *TfFlow_TrendmicroProperty
	// Experimental.
	Veeva() TfFlow_VeevaPropertyOutputReference
	// Experimental.
	VeevaInput() *TfFlow_VeevaProperty
	// Experimental.
	Zendesk() TfFlow_SourceFlowConfigSourceConnectorPropertiesZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *TfFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty
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
	PutAmplitude(value *TfFlow_AmplitudeProperty)
	// Experimental.
	PutCustomConnector(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty)
	// Experimental.
	PutDatadog(value *TfFlow_DatadogProperty)
	// Experimental.
	PutDynatrace(value *TfFlow_DynatraceProperty)
	// Experimental.
	PutGoogleAnalytics(value *TfFlow_GoogleAnalyticsProperty)
	// Experimental.
	PutInforNexus(value *TfFlow_InforNexusProperty)
	// Experimental.
	PutMarketo(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty)
	// Experimental.
	PutS3(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesS3Property)
	// Experimental.
	PutSalesforce(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty)
	// Experimental.
	PutSapoData(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty)
	// Experimental.
	PutServiceNow(value *TfFlow_ServiceNowProperty)
	// Experimental.
	PutSingular(value *TfFlow_SingularProperty)
	// Experimental.
	PutSlack(value *TfFlow_SlackProperty)
	// Experimental.
	PutTrendmicro(value *TfFlow_TrendmicroProperty)
	// Experimental.
	PutVeeva(value *TfFlow_VeevaProperty)
	// Experimental.
	PutZendesk(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty)
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

// The jsii proxy struct for TfFlow_SourceConnectorPropertiesPropertyOutputReference
type jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Amplitude() TfFlow_AmplitudePropertyOutputReference {
	var returns TfFlow_AmplitudePropertyOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) AmplitudeInput() *TfFlow_AmplitudeProperty {
	var returns *TfFlow_AmplitudeProperty
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) CustomConnector() TfFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorPropertyOutputReference {
	var returns TfFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) CustomConnectorInput() *TfFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty {
	var returns *TfFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Datadog() TfFlow_DatadogPropertyOutputReference {
	var returns TfFlow_DatadogPropertyOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) DatadogInput() *TfFlow_DatadogProperty {
	var returns *TfFlow_DatadogProperty
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Dynatrace() TfFlow_DynatracePropertyOutputReference {
	var returns TfFlow_DynatracePropertyOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) DynatraceInput() *TfFlow_DynatraceProperty {
	var returns *TfFlow_DynatraceProperty
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) GoogleAnalytics() TfFlow_GoogleAnalyticsPropertyOutputReference {
	var returns TfFlow_GoogleAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) GoogleAnalyticsInput() *TfFlow_GoogleAnalyticsProperty {
	var returns *TfFlow_GoogleAnalyticsProperty
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) InforNexus() TfFlow_InforNexusPropertyOutputReference {
	var returns TfFlow_InforNexusPropertyOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) InforNexusInput() *TfFlow_InforNexusProperty {
	var returns *TfFlow_InforNexusProperty
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) InternalValue() *TfFlow_SourceConnectorPropertiesProperty {
	var returns *TfFlow_SourceConnectorPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Marketo() TfFlow_SourceFlowConfigSourceConnectorPropertiesMarketoPropertyOutputReference {
	var returns TfFlow_SourceFlowConfigSourceConnectorPropertiesMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) MarketoInput() *TfFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty {
	var returns *TfFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) S3() TfFlow_SourceFlowConfigSourceConnectorPropertiesS3PropertyOutputReference {
	var returns TfFlow_SourceFlowConfigSourceConnectorPropertiesS3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) S3Input() *TfFlow_SourceFlowConfigSourceConnectorPropertiesS3Property {
	var returns *TfFlow_SourceFlowConfigSourceConnectorPropertiesS3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Salesforce() TfFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference {
	var returns TfFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) SalesforceInput() *TfFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty {
	var returns *TfFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) SapoData() TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference {
	var returns TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) SapoDataInput() *TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty {
	var returns *TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ServiceNow() TfFlow_ServiceNowPropertyOutputReference {
	var returns TfFlow_ServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ServiceNowInput() *TfFlow_ServiceNowProperty {
	var returns *TfFlow_ServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Singular() TfFlow_SingularPropertyOutputReference {
	var returns TfFlow_SingularPropertyOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) SingularInput() *TfFlow_SingularProperty {
	var returns *TfFlow_SingularProperty
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Slack() TfFlow_SlackPropertyOutputReference {
	var returns TfFlow_SlackPropertyOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) SlackInput() *TfFlow_SlackProperty {
	var returns *TfFlow_SlackProperty
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Trendmicro() TfFlow_TrendmicroPropertyOutputReference {
	var returns TfFlow_TrendmicroPropertyOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) TrendmicroInput() *TfFlow_TrendmicroProperty {
	var returns *TfFlow_TrendmicroProperty
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Veeva() TfFlow_VeevaPropertyOutputReference {
	var returns TfFlow_VeevaPropertyOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) VeevaInput() *TfFlow_VeevaProperty {
	var returns *TfFlow_VeevaProperty
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Zendesk() TfFlow_SourceFlowConfigSourceConnectorPropertiesZendeskPropertyOutputReference {
	var returns TfFlow_SourceFlowConfigSourceConnectorPropertiesZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ZendeskInput() *TfFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty {
	var returns *TfFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlow_SourceConnectorPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlow_SourceConnectorPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlow_SourceConnectorPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.SourceConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlow_SourceConnectorPropertiesPropertyOutputReference_Override(t TfFlow_SourceConnectorPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.SourceConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference)SetInternalValue(val *TfFlow_SourceConnectorPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutAmplitude(value *TfFlow_AmplitudeProperty) {
	if err := t.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutCustomConnector(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty) {
	if err := t.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutDatadog(value *TfFlow_DatadogProperty) {
	if err := t.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDatadog",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutDynatrace(value *TfFlow_DynatraceProperty) {
	if err := t.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutGoogleAnalytics(value *TfFlow_GoogleAnalyticsProperty) {
	if err := t.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutInforNexus(value *TfFlow_InforNexusProperty) {
	if err := t.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutMarketo(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty) {
	if err := t.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMarketo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutS3(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesS3Property) {
	if err := t.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutSalesforce(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty) {
	if err := t.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutSapoData(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty) {
	if err := t.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSapoData",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutServiceNow(value *TfFlow_ServiceNowProperty) {
	if err := t.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutSingular(value *TfFlow_SingularProperty) {
	if err := t.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingular",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutSlack(value *TfFlow_SlackProperty) {
	if err := t.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSlack",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutTrendmicro(value *TfFlow_TrendmicroProperty) {
	if err := t.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutVeeva(value *TfFlow_VeevaProperty) {
	if err := t.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVeeva",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) PutZendesk(value *TfFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty) {
	if err := t.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putZendesk",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		t,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		t,
		"resetDatadog",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		t,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		t,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		t,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		t,
		"resetMarketo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		t,
		"resetS3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		t,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		t,
		"resetSapoData",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		t,
		"resetSingular",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		t,
		"resetSlack",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		t,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		t,
		"resetVeeva",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		t,
		"resetZendesk",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlow_SourceConnectorPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

