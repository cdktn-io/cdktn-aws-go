package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AliasTarget() AwsRoute53RecordsExclusive_AliasTargetPropertyList
	// Experimental.
	AliasTargetInput() interface{}
	// Experimental.
	CidrRoutingConfig() AwsRoute53RecordsExclusive_CidrRoutingConfigPropertyList
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
	Geolocation() AwsRoute53RecordsExclusive_GeolocationPropertyList
	// Experimental.
	GeolocationInput() interface{}
	// Experimental.
	GeoproximityLocation() AwsRoute53RecordsExclusive_GeoproximityLocationPropertyList
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
	ResourceRecords() AwsRoute53RecordsExclusive_ResourceRecordsPropertyList
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

// The jsii proxy struct for AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference
type jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) AliasTarget() AwsRoute53RecordsExclusive_AliasTargetPropertyList {
	var returns AwsRoute53RecordsExclusive_AliasTargetPropertyList
	_jsii_.Get(
		j,
		"aliasTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) AliasTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aliasTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) CidrRoutingConfig() AwsRoute53RecordsExclusive_CidrRoutingConfigPropertyList {
	var returns AwsRoute53RecordsExclusive_CidrRoutingConfigPropertyList
	_jsii_.Get(
		j,
		"cidrRoutingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) CidrRoutingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cidrRoutingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) Failover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) FailoverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) Geolocation() AwsRoute53RecordsExclusive_GeolocationPropertyList {
	var returns AwsRoute53RecordsExclusive_GeolocationPropertyList
	_jsii_.Get(
		j,
		"geolocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GeolocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geolocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GeoproximityLocation() AwsRoute53RecordsExclusive_GeoproximityLocationPropertyList {
	var returns AwsRoute53RecordsExclusive_GeoproximityLocationPropertyList
	_jsii_.Get(
		j,
		"geoproximityLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GeoproximityLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoproximityLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) HealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) HealthCheckIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) MultiValueAnswer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValueAnswer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) MultiValueAnswerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValueAnswerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResourceRecords() AwsRoute53RecordsExclusive_ResourceRecordsPropertyList {
	var returns AwsRoute53RecordsExclusive_ResourceRecordsPropertyList
	_jsii_.Get(
		j,
		"resourceRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResourceRecordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) SetIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) SetIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) TrafficPolicyInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficPolicyInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) TrafficPolicyInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficPolicyInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53.AwsRoute53RecordsExclusive.ResourceRecordSetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference_Override(a AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.AwsRoute53RecordsExclusive.ResourceRecordSetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetFailover(val *string) {
	if err := j.validateSetFailoverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failover",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetHealthCheckId(val *string) {
	if err := j.validateSetHealthCheckIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetMultiValueAnswer(val interface{}) {
	if err := j.validateSetMultiValueAnswerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiValueAnswer",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetSetIdentifier(val *string) {
	if err := j.validateSetSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetTrafficPolicyInstanceId(val *string) {
	if err := j.validateSetTrafficPolicyInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficPolicyInstanceId",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) PutAliasTarget(value interface{}) {
	if err := a.validatePutAliasTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAliasTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) PutCidrRoutingConfig(value interface{}) {
	if err := a.validatePutCidrRoutingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCidrRoutingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) PutGeolocation(value interface{}) {
	if err := a.validatePutGeolocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGeolocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) PutGeoproximityLocation(value interface{}) {
	if err := a.validatePutGeoproximityLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGeoproximityLocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) PutResourceRecords(value interface{}) {
	if err := a.validatePutResourceRecordsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceRecords",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetAliasTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetAliasTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetCidrRoutingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCidrRoutingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetFailover() {
	_jsii_.InvokeVoid(
		a,
		"resetFailover",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetGeolocation() {
	_jsii_.InvokeVoid(
		a,
		"resetGeolocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetGeoproximityLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetGeoproximityLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetHealthCheckId() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetMultiValueAnswer() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiValueAnswer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetResourceRecords() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceRecords",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetSetIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetSetIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetTrafficPolicyInstanceId() {
	_jsii_.InvokeVoid(
		a,
		"resetTrafficPolicyInstanceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		a,
		"resetWeight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRoute53RecordsExclusive_ResourceRecordSetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

