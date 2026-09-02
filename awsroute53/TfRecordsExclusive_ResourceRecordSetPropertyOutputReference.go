package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRecordsExclusive_ResourceRecordSetPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AliasTarget() TfRecordsExclusive_AliasTargetPropertyList
	// Experimental.
	AliasTargetInput() interface{}
	// Experimental.
	CidrRoutingConfig() TfRecordsExclusive_CidrRoutingConfigPropertyList
	// Experimental.
	CidrRoutingConfigInput() interface{}
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
	Failover() *string
	// Experimental.
	SetFailover(val *string)
	// Experimental.
	FailoverInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	Geolocation() TfRecordsExclusive_GeolocationPropertyList
	// Experimental.
	GeolocationInput() interface{}
	// Experimental.
	GeoproximityLocation() TfRecordsExclusive_GeoproximityLocationPropertyList
	// Experimental.
	GeoproximityLocationInput() interface{}
	// Experimental.
	HealthCheckId() *string
	// Experimental.
	SetHealthCheckId(val *string)
	// Experimental.
	HealthCheckIdInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MultiValueAnswer() interface{}
	// Experimental.
	SetMultiValueAnswer(val interface{})
	// Experimental.
	MultiValueAnswerInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ResourceRecords() TfRecordsExclusive_ResourceRecordsPropertyList
	// Experimental.
	ResourceRecordsInput() interface{}
	// Experimental.
	SetIdentifier() *string
	// Experimental.
	SetSetIdentifier(val *string)
	// Experimental.
	SetIdentifierInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrafficPolicyInstanceId() *string
	// Experimental.
	SetTrafficPolicyInstanceId(val *string)
	// Experimental.
	TrafficPolicyInstanceIdInput() *string
	// Experimental.
	Ttl() *float64
	// Experimental.
	SetTtl(val *float64)
	// Experimental.
	TtlInput() *float64
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
	// Experimental.
	Weight() *float64
	// Experimental.
	SetWeight(val *float64)
	// Experimental.
	WeightInput() *float64
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
	PutAliasTarget(value interface{})
	// Experimental.
	PutCidrRoutingConfig(value interface{})
	// Experimental.
	PutGeolocation(value interface{})
	// Experimental.
	PutGeoproximityLocation(value interface{})
	// Experimental.
	PutResourceRecords(value interface{})
	// Experimental.
	ResetAliasTarget()
	// Experimental.
	ResetCidrRoutingConfig()
	// Experimental.
	ResetFailover()
	// Experimental.
	ResetGeolocation()
	// Experimental.
	ResetGeoproximityLocation()
	// Experimental.
	ResetHealthCheckId()
	// Experimental.
	ResetMultiValueAnswer()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetResourceRecords()
	// Experimental.
	ResetSetIdentifier()
	// Experimental.
	ResetTrafficPolicyInstanceId()
	// Experimental.
	ResetTtl()
	// Experimental.
	ResetType()
	// Experimental.
	ResetWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRecordsExclusive_ResourceRecordSetPropertyOutputReference
type jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) AliasTarget() TfRecordsExclusive_AliasTargetPropertyList {
	var returns TfRecordsExclusive_AliasTargetPropertyList
	_jsii_.Get(
		j,
		"aliasTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) AliasTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aliasTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) CidrRoutingConfig() TfRecordsExclusive_CidrRoutingConfigPropertyList {
	var returns TfRecordsExclusive_CidrRoutingConfigPropertyList
	_jsii_.Get(
		j,
		"cidrRoutingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) CidrRoutingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cidrRoutingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) Failover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) FailoverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) Geolocation() TfRecordsExclusive_GeolocationPropertyList {
	var returns TfRecordsExclusive_GeolocationPropertyList
	_jsii_.Get(
		j,
		"geolocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GeolocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geolocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GeoproximityLocation() TfRecordsExclusive_GeoproximityLocationPropertyList {
	var returns TfRecordsExclusive_GeoproximityLocationPropertyList
	_jsii_.Get(
		j,
		"geoproximityLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GeoproximityLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoproximityLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) HealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) HealthCheckIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) MultiValueAnswer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValueAnswer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) MultiValueAnswerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValueAnswerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResourceRecords() TfRecordsExclusive_ResourceRecordsPropertyList {
	var returns TfRecordsExclusive_ResourceRecordsPropertyList
	_jsii_.Get(
		j,
		"resourceRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResourceRecordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) SetIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) SetIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) TrafficPolicyInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficPolicyInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) TrafficPolicyInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficPolicyInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRecordsExclusive_ResourceRecordSetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfRecordsExclusive_ResourceRecordSetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRecordsExclusive_ResourceRecordSetPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53.TfRecordsExclusive.ResourceRecordSetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRecordsExclusive_ResourceRecordSetPropertyOutputReference_Override(t TfRecordsExclusive_ResourceRecordSetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.TfRecordsExclusive.ResourceRecordSetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetFailover(val *string) {
	if err := j.validateSetFailoverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failover",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetHealthCheckId(val *string) {
	if err := j.validateSetHealthCheckIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetMultiValueAnswer(val interface{}) {
	if err := j.validateSetMultiValueAnswerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiValueAnswer",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetSetIdentifier(val *string) {
	if err := j.validateSetSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetTrafficPolicyInstanceId(val *string) {
	if err := j.validateSetTrafficPolicyInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficPolicyInstanceId",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) PutAliasTarget(value interface{}) {
	if err := t.validatePutAliasTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAliasTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) PutCidrRoutingConfig(value interface{}) {
	if err := t.validatePutCidrRoutingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCidrRoutingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) PutGeolocation(value interface{}) {
	if err := t.validatePutGeolocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGeolocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) PutGeoproximityLocation(value interface{}) {
	if err := t.validatePutGeoproximityLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGeoproximityLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) PutResourceRecords(value interface{}) {
	if err := t.validatePutResourceRecordsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceRecords",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetAliasTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetAliasTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetCidrRoutingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCidrRoutingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetFailover() {
	_jsii_.InvokeVoid(
		t,
		"resetFailover",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetGeolocation() {
	_jsii_.InvokeVoid(
		t,
		"resetGeolocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetGeoproximityLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetGeoproximityLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetHealthCheckId() {
	_jsii_.InvokeVoid(
		t,
		"resetHealthCheckId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetMultiValueAnswer() {
	_jsii_.InvokeVoid(
		t,
		"resetMultiValueAnswer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetResourceRecords() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceRecords",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetSetIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetSetIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetTrafficPolicyInstanceId() {
	_jsii_.InvokeVoid(
		t,
		"resetTrafficPolicyInstanceId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		t,
		"resetType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		t,
		"resetWeight",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRecordsExclusive_ResourceRecordSetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

