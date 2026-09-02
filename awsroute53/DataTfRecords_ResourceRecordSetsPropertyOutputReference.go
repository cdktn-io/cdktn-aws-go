package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfRecords_ResourceRecordSetsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AliasTarget() DataTfRecords_AliasTargetPropertyOutputReference
	// Experimental.
	CidrRoutingConfig() DataTfRecords_CidrRoutingConfigPropertyOutputReference
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
	Fqn() *string
	// Experimental.
	Geolocation() DataTfRecords_GeolocationPropertyOutputReference
	// Experimental.
	GeoproximityLocation() DataTfRecords_GeoproximityLocationPropertyOutputReference
	// Experimental.
	HealthCheckId() *string
	// Experimental.
	InternalValue() *DataTfRecords_ResourceRecordSetsProperty
	// Experimental.
	SetInternalValue(val *DataTfRecords_ResourceRecordSetsProperty)
	// Experimental.
	MultiValueAnswer() cdktn.IResolvable
	// Experimental.
	Name() *string
	// Experimental.
	Region() *string
	// Experimental.
	ResourceRecords() DataTfRecords_ResourceRecordsPropertyList
	// Experimental.
	SetIdentifier() *string
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
	Ttl() *float64
	// Experimental.
	Type() *string
	// Experimental.
	Weight() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfRecords_ResourceRecordSetsPropertyOutputReference
type jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) AliasTarget() DataTfRecords_AliasTargetPropertyOutputReference {
	var returns DataTfRecords_AliasTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"aliasTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) CidrRoutingConfig() DataTfRecords_CidrRoutingConfigPropertyOutputReference {
	var returns DataTfRecords_CidrRoutingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"cidrRoutingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) Failover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) Geolocation() DataTfRecords_GeolocationPropertyOutputReference {
	var returns DataTfRecords_GeolocationPropertyOutputReference
	_jsii_.Get(
		j,
		"geolocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) GeoproximityLocation() DataTfRecords_GeoproximityLocationPropertyOutputReference {
	var returns DataTfRecords_GeoproximityLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"geoproximityLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) HealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) InternalValue() *DataTfRecords_ResourceRecordSetsProperty {
	var returns *DataTfRecords_ResourceRecordSetsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) MultiValueAnswer() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"multiValueAnswer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) ResourceRecords() DataTfRecords_ResourceRecordsPropertyList {
	var returns DataTfRecords_ResourceRecordsPropertyList
	_jsii_.Get(
		j,
		"resourceRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) SetIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) TrafficPolicyInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficPolicyInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfRecords_ResourceRecordSetsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfRecords_ResourceRecordSetsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfRecords_ResourceRecordSetsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53.DataTfRecords.ResourceRecordSetsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfRecords_ResourceRecordSetsPropertyOutputReference_Override(d DataTfRecords_ResourceRecordSetsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.DataTfRecords.ResourceRecordSetsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference)SetInternalValue(val *DataTfRecords_ResourceRecordSetsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfRecords_ResourceRecordSetsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

