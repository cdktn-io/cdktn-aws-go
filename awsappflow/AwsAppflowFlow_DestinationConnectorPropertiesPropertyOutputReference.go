package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference interface {
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
	CustomConnector() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty
	// Experimental.
	CustomerProfiles() AwsAppflowFlow_CustomerProfilesPropertyOutputReference
	// Experimental.
	CustomerProfilesInput() *AwsAppflowFlow_CustomerProfilesProperty
	// Experimental.
	EventBridge() AwsAppflowFlow_EventBridgePropertyOutputReference
	// Experimental.
	EventBridgeInput() *AwsAppflowFlow_EventBridgeProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Honeycode() AwsAppflowFlow_HoneycodePropertyOutputReference
	// Experimental.
	HoneycodeInput() *AwsAppflowFlow_HoneycodeProperty
	// Experimental.
	InternalValue() *AwsAppflowFlow_DestinationConnectorPropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsAppflowFlow_DestinationConnectorPropertiesProperty)
	// Experimental.
	LookoutMetrics() AwsAppflowFlow_LookoutMetricsPropertyOutputReference
	// Experimental.
	LookoutMetricsInput() *AwsAppflowFlow_LookoutMetricsProperty
	// Experimental.
	Marketo() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty
	// Experimental.
	Redshift() AwsAppflowFlow_RedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *AwsAppflowFlow_RedshiftProperty
	// Experimental.
	S3() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesS3PropertyOutputReference
	// Experimental.
	S3Input() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property
	// Experimental.
	Salesforce() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty
	// Experimental.
	SapoData() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty
	// Experimental.
	Snowflake() AwsAppflowFlow_SnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *AwsAppflowFlow_SnowflakeProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Upsolver() AwsAppflowFlow_UpsolverPropertyOutputReference
	// Experimental.
	UpsolverInput() *AwsAppflowFlow_UpsolverProperty
	// Experimental.
	Zendesk() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty
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
	PutCustomConnector(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty)
	// Experimental.
	PutCustomerProfiles(value *AwsAppflowFlow_CustomerProfilesProperty)
	// Experimental.
	PutEventBridge(value *AwsAppflowFlow_EventBridgeProperty)
	// Experimental.
	PutHoneycode(value *AwsAppflowFlow_HoneycodeProperty)
	// Experimental.
	PutLookoutMetrics(value *AwsAppflowFlow_LookoutMetricsProperty)
	// Experimental.
	PutMarketo(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty)
	// Experimental.
	PutRedshift(value *AwsAppflowFlow_RedshiftProperty)
	// Experimental.
	PutS3(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property)
	// Experimental.
	PutSalesforce(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty)
	// Experimental.
	PutSapoData(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty)
	// Experimental.
	PutSnowflake(value *AwsAppflowFlow_SnowflakeProperty)
	// Experimental.
	PutUpsolver(value *AwsAppflowFlow_UpsolverProperty)
	// Experimental.
	PutZendesk(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty)
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

// The jsii proxy struct for AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference
type jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomConnector() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorPropertyOutputReference {
	var returns AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomConnectorInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomerProfiles() AwsAppflowFlow_CustomerProfilesPropertyOutputReference {
	var returns AwsAppflowFlow_CustomerProfilesPropertyOutputReference
	_jsii_.Get(
		j,
		"customerProfiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomerProfilesInput() *AwsAppflowFlow_CustomerProfilesProperty {
	var returns *AwsAppflowFlow_CustomerProfilesProperty
	_jsii_.Get(
		j,
		"customerProfilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) EventBridge() AwsAppflowFlow_EventBridgePropertyOutputReference {
	var returns AwsAppflowFlow_EventBridgePropertyOutputReference
	_jsii_.Get(
		j,
		"eventBridge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) EventBridgeInput() *AwsAppflowFlow_EventBridgeProperty {
	var returns *AwsAppflowFlow_EventBridgeProperty
	_jsii_.Get(
		j,
		"eventBridgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) Honeycode() AwsAppflowFlow_HoneycodePropertyOutputReference {
	var returns AwsAppflowFlow_HoneycodePropertyOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) HoneycodeInput() *AwsAppflowFlow_HoneycodeProperty {
	var returns *AwsAppflowFlow_HoneycodeProperty
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) InternalValue() *AwsAppflowFlow_DestinationConnectorPropertiesProperty {
	var returns *AwsAppflowFlow_DestinationConnectorPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) LookoutMetrics() AwsAppflowFlow_LookoutMetricsPropertyOutputReference {
	var returns AwsAppflowFlow_LookoutMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"lookoutMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) LookoutMetricsInput() *AwsAppflowFlow_LookoutMetricsProperty {
	var returns *AwsAppflowFlow_LookoutMetricsProperty
	_jsii_.Get(
		j,
		"lookoutMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) Marketo() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoPropertyOutputReference {
	var returns AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) MarketoInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) Redshift() AwsAppflowFlow_RedshiftPropertyOutputReference {
	var returns AwsAppflowFlow_RedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) RedshiftInput() *AwsAppflowFlow_RedshiftProperty {
	var returns *AwsAppflowFlow_RedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) S3() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesS3PropertyOutputReference {
	var returns AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesS3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) S3Input() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) Salesforce() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference {
	var returns AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) SalesforceInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) SapoData() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference {
	var returns AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) SapoDataInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) Snowflake() AwsAppflowFlow_SnowflakePropertyOutputReference {
	var returns AwsAppflowFlow_SnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) SnowflakeInput() *AwsAppflowFlow_SnowflakeProperty {
	var returns *AwsAppflowFlow_SnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) Upsolver() AwsAppflowFlow_UpsolverPropertyOutputReference {
	var returns AwsAppflowFlow_UpsolverPropertyOutputReference
	_jsii_.Get(
		j,
		"upsolver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) UpsolverInput() *AwsAppflowFlow_UpsolverProperty {
	var returns *AwsAppflowFlow_UpsolverProperty
	_jsii_.Get(
		j,
		"upsolverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) Zendesk() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskPropertyOutputReference {
	var returns AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ZendeskInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.DestinationConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference_Override(a AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.DestinationConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference)SetInternalValue(val *AwsAppflowFlow_DestinationConnectorPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutCustomConnector(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutCustomerProfiles(value *AwsAppflowFlow_CustomerProfilesProperty) {
	if err := a.validatePutCustomerProfilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomerProfiles",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutEventBridge(value *AwsAppflowFlow_EventBridgeProperty) {
	if err := a.validatePutEventBridgeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventBridge",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutHoneycode(value *AwsAppflowFlow_HoneycodeProperty) {
	if err := a.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutLookoutMetrics(value *AwsAppflowFlow_LookoutMetricsProperty) {
	if err := a.validatePutLookoutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLookoutMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutMarketo(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutRedshift(value *AwsAppflowFlow_RedshiftProperty) {
	if err := a.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshift",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutS3(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutSalesforce(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutSapoData(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutSnowflake(value *AwsAppflowFlow_SnowflakeProperty) {
	if err := a.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutUpsolver(value *AwsAppflowFlow_UpsolverProperty) {
	if err := a.validatePutUpsolverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpsolver",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) PutZendesk(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetCustomerProfiles() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerProfiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetEventBridge() {
	_jsii_.InvokeVoid(
		a,
		"resetEventBridge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		a,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetLookoutMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetLookoutMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetUpsolver() {
	_jsii_.InvokeVoid(
		a,
		"resetUpsolver",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationConnectorPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

