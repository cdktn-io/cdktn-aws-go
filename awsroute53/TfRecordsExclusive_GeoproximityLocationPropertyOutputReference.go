package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRecordsExclusive_GeoproximityLocationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsRegion() *string
	// Experimental.
	SetAwsRegion(val *string)
	// Experimental.
	AwsRegionInput() *string
	// Experimental.
	Bias() *float64
	// Experimental.
	SetBias(val *float64)
	// Experimental.
	BiasInput() *float64
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
	// Experimental.
	Coordinates() TfRecordsExclusive_CoordinatesPropertyList
	// Experimental.
	CoordinatesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LocalZoneGroup() *string
	// Experimental.
	SetLocalZoneGroup(val *string)
	// Experimental.
	LocalZoneGroupInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutCoordinates(value interface{})
	// Experimental.
	ResetAwsRegion()
	// Experimental.
	ResetBias()
	// Experimental.
	ResetCoordinates()
	// Experimental.
	ResetLocalZoneGroup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRecordsExclusive_GeoproximityLocationPropertyOutputReference
type jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) Bias() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) BiasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"biasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) Coordinates() TfRecordsExclusive_CoordinatesPropertyList {
	var returns TfRecordsExclusive_CoordinatesPropertyList
	_jsii_.Get(
		j,
		"coordinates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) CoordinatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coordinatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) LocalZoneGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localZoneGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) LocalZoneGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localZoneGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRecordsExclusive_GeoproximityLocationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfRecordsExclusive_GeoproximityLocationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRecordsExclusive_GeoproximityLocationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53.TfRecordsExclusive.GeoproximityLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRecordsExclusive_GeoproximityLocationPropertyOutputReference_Override(t TfRecordsExclusive_GeoproximityLocationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.TfRecordsExclusive.GeoproximityLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference)SetAwsRegion(val *string) {
	if err := j.validateSetAwsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference)SetBias(val *float64) {
	if err := j.validateSetBiasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bias",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference)SetLocalZoneGroup(val *string) {
	if err := j.validateSetLocalZoneGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localZoneGroup",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) PutCoordinates(value interface{}) {
	if err := t.validatePutCoordinatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCoordinates",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) ResetAwsRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) ResetBias() {
	_jsii_.InvokeVoid(
		t,
		"resetBias",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) ResetCoordinates() {
	_jsii_.InvokeVoid(
		t,
		"resetCoordinates",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) ResetLocalZoneGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetLocalZoneGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRecordsExclusive_GeoproximityLocationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

