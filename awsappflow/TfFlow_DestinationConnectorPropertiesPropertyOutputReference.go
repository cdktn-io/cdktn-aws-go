package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlow_DestinationConnectorPropertiesPropertyOutputReference interface {
	cdktn.ComplexObject
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
	CustomConnector() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty
	// Experimental.
	CustomerProfiles() TfFlow_CustomerProfilesPropertyOutputReference
	// Experimental.
	CustomerProfilesInput() *TfFlow_CustomerProfilesProperty
	// Experimental.
	EventBridge() TfFlow_EventBridgePropertyOutputReference
	// Experimental.
	EventBridgeInput() *TfFlow_EventBridgeProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Honeycode() TfFlow_HoneycodePropertyOutputReference
	// Experimental.
	HoneycodeInput() *TfFlow_HoneycodeProperty
	// Experimental.
	InternalValue() *TfFlow_DestinationConnectorPropertiesProperty
	// Experimental.
	SetInternalValue(val *TfFlow_DestinationConnectorPropertiesProperty)
	// Experimental.
	LookoutMetrics() TfFlow_LookoutMetricsPropertyOutputReference
	// Experimental.
	LookoutMetricsInput() *TfFlow_LookoutMetricsProperty
	// Experimental.
	Marketo() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty
	// Experimental.
	Redshift() TfFlow_RedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *TfFlow_RedshiftProperty
	// Experimental.
	S3() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3PropertyOutputReference
	// Experimental.
	S3Input() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property
	// Experimental.
	Salesforce() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty
	// Experimental.
	SapoData() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty
	// Experimental.
	Snowflake() TfFlow_SnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *TfFlow_SnowflakeProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Upsolver() TfFlow_UpsolverPropertyOutputReference
	// Experimental.
	UpsolverInput() *TfFlow_UpsolverProperty
	// Experimental.
	Zendesk() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty
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
	PutCustomConnector(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty)
	// Experimental.
	PutCustomerProfiles(value *TfFlow_CustomerProfilesProperty)
	// Experimental.
	PutEventBridge(value *TfFlow_EventBridgeProperty)
	// Experimental.
	PutHoneycode(value *TfFlow_HoneycodeProperty)
	// Experimental.
	PutLookoutMetrics(value *TfFlow_LookoutMetricsProperty)
	// Experimental.
	PutMarketo(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty)
	// Experimental.
	PutRedshift(value *TfFlow_RedshiftProperty)
	// Experimental.
	PutS3(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property)
	// Experimental.
	PutSalesforce(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty)
	// Experimental.
	PutSapoData(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty)
	// Experimental.
	PutSnowflake(value *TfFlow_SnowflakeProperty)
	// Experimental.
	PutUpsolver(value *TfFlow_UpsolverProperty)
	// Experimental.
	PutZendesk(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty)
	// Experimental.
	ResetCustomConnector()
	// Experimental.
	ResetCustomerProfiles()
	// Experimental.
	ResetEventBridge()
	// Experimental.
	ResetHoneycode()
	// Experimental.
	ResetLookoutMetrics()
	// Experimental.
	ResetMarketo()
	// Experimental.
	ResetRedshift()
	// Experimental.
	ResetS3()
	// Experimental.
	ResetSalesforce()
	// Experimental.
	ResetSapoData()
	// Experimental.
	ResetSnowflake()
	// Experimental.
	ResetUpsolver()
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

// The jsii proxy struct for TfFlow_DestinationConnectorPropertiesPropertyOutputReference
type jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomConnector() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorPropertyOutputReference {
	var returns TfFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomConnectorInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomerProfiles() TfFlow_CustomerProfilesPropertyOutputReference {
	var returns TfFlow_CustomerProfilesPropertyOutputReference
	_jsii_.Get(
		j,
		"customerProfiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomerProfilesInput() *TfFlow_CustomerProfilesProperty {
	var returns *TfFlow_CustomerProfilesProperty
	_jsii_.Get(
		j,
		"customerProfilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) EventBridge() TfFlow_EventBridgePropertyOutputReference {
	var returns TfFlow_EventBridgePropertyOutputReference
	_jsii_.Get(
		j,
		"eventBridge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) EventBridgeInput() *TfFlow_EventBridgeProperty {
	var returns *TfFlow_EventBridgeProperty
	_jsii_.Get(
		j,
		"eventBridgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) Honeycode() TfFlow_HoneycodePropertyOutputReference {
	var returns TfFlow_HoneycodePropertyOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) HoneycodeInput() *TfFlow_HoneycodeProperty {
	var returns *TfFlow_HoneycodeProperty
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) InternalValue() *TfFlow_DestinationConnectorPropertiesProperty {
	var returns *TfFlow_DestinationConnectorPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) LookoutMetrics() TfFlow_LookoutMetricsPropertyOutputReference {
	var returns TfFlow_LookoutMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"lookoutMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) LookoutMetricsInput() *TfFlow_LookoutMetricsProperty {
	var returns *TfFlow_LookoutMetricsProperty
	_jsii_.Get(
		j,
		"lookoutMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) Marketo() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoPropertyOutputReference {
	var returns TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) MarketoInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) Redshift() TfFlow_RedshiftPropertyOutputReference {
	var returns TfFlow_RedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) RedshiftInput() *TfFlow_RedshiftProperty {
	var returns *TfFlow_RedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) S3() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3PropertyOutputReference {
	var returns TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) S3Input() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) Salesforce() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference {
	var returns TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) SalesforceInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) SapoData() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference {
	var returns TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) SapoDataInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) Snowflake() TfFlow_SnowflakePropertyOutputReference {
	var returns TfFlow_SnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) SnowflakeInput() *TfFlow_SnowflakeProperty {
	var returns *TfFlow_SnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) Upsolver() TfFlow_UpsolverPropertyOutputReference {
	var returns TfFlow_UpsolverPropertyOutputReference
	_jsii_.Get(
		j,
		"upsolver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) UpsolverInput() *TfFlow_UpsolverProperty {
	var returns *TfFlow_UpsolverProperty
	_jsii_.Get(
		j,
		"upsolverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) Zendesk() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskPropertyOutputReference {
	var returns TfFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ZendeskInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlow_DestinationConnectorPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlow_DestinationConnectorPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlow_DestinationConnectorPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.DestinationConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlow_DestinationConnectorPropertiesPropertyOutputReference_Override(t TfFlow_DestinationConnectorPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.DestinationConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference)SetInternalValue(val *TfFlow_DestinationConnectorPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutCustomConnector(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty) {
	if err := t.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutCustomerProfiles(value *TfFlow_CustomerProfilesProperty) {
	if err := t.validatePutCustomerProfilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomerProfiles",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutEventBridge(value *TfFlow_EventBridgeProperty) {
	if err := t.validatePutEventBridgeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEventBridge",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutHoneycode(value *TfFlow_HoneycodeProperty) {
	if err := t.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutLookoutMetrics(value *TfFlow_LookoutMetricsProperty) {
	if err := t.validatePutLookoutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLookoutMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutMarketo(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty) {
	if err := t.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMarketo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutRedshift(value *TfFlow_RedshiftProperty) {
	if err := t.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedshift",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutS3(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property) {
	if err := t.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutSalesforce(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty) {
	if err := t.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutSapoData(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty) {
	if err := t.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSapoData",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutSnowflake(value *TfFlow_SnowflakeProperty) {
	if err := t.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutUpsolver(value *TfFlow_UpsolverProperty) {
	if err := t.validatePutUpsolverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUpsolver",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) PutZendesk(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty) {
	if err := t.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putZendesk",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetCustomerProfiles() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomerProfiles",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetEventBridge() {
	_jsii_.InvokeVoid(
		t,
		"resetEventBridge",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		t,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetLookoutMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetLookoutMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		t,
		"resetMarketo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		t,
		"resetRedshift",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		t,
		"resetS3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		t,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		t,
		"resetSapoData",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		t,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetUpsolver() {
	_jsii_.InvokeVoid(
		t,
		"resetUpsolver",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		t,
		"resetZendesk",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlow_DestinationConnectorPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

