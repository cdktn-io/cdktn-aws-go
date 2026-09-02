package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeliveryStream_DeserializerPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	HiveJsonSerDe() TfDeliveryStream_HiveJsonSerDePropertyOutputReference
	// Experimental.
	HiveJsonSerDeInput() *TfDeliveryStream_HiveJsonSerDeProperty
	// Experimental.
	InternalValue() *TfDeliveryStream_DeserializerProperty
	// Experimental.
	SetInternalValue(val *TfDeliveryStream_DeserializerProperty)
	// Experimental.
	OpenXJsonSerDe() TfDeliveryStream_OpenXJsonSerDePropertyOutputReference
	// Experimental.
	OpenXJsonSerDeInput() *TfDeliveryStream_OpenXJsonSerDeProperty
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
	PutHiveJsonSerDe(value *TfDeliveryStream_HiveJsonSerDeProperty)
	// Experimental.
	PutOpenXJsonSerDe(value *TfDeliveryStream_OpenXJsonSerDeProperty)
	// Experimental.
	ResetHiveJsonSerDe()
	// Experimental.
	ResetOpenXJsonSerDe()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDeliveryStream_DeserializerPropertyOutputReference
type jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) HiveJsonSerDe() TfDeliveryStream_HiveJsonSerDePropertyOutputReference {
	var returns TfDeliveryStream_HiveJsonSerDePropertyOutputReference
	_jsii_.Get(
		j,
		"hiveJsonSerDe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) HiveJsonSerDeInput() *TfDeliveryStream_HiveJsonSerDeProperty {
	var returns *TfDeliveryStream_HiveJsonSerDeProperty
	_jsii_.Get(
		j,
		"hiveJsonSerDeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) InternalValue() *TfDeliveryStream_DeserializerProperty {
	var returns *TfDeliveryStream_DeserializerProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) OpenXJsonSerDe() TfDeliveryStream_OpenXJsonSerDePropertyOutputReference {
	var returns TfDeliveryStream_OpenXJsonSerDePropertyOutputReference
	_jsii_.Get(
		j,
		"openXJsonSerDe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) OpenXJsonSerDeInput() *TfDeliveryStream_OpenXJsonSerDeProperty {
	var returns *TfDeliveryStream_OpenXJsonSerDeProperty
	_jsii_.Get(
		j,
		"openXJsonSerDeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeliveryStream_DeserializerPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeliveryStream_DeserializerPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeliveryStream_DeserializerPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.DeserializerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeliveryStream_DeserializerPropertyOutputReference_Override(t TfDeliveryStream_DeserializerPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.DeserializerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference)SetInternalValue(val *TfDeliveryStream_DeserializerProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) PutHiveJsonSerDe(value *TfDeliveryStream_HiveJsonSerDeProperty) {
	if err := t.validatePutHiveJsonSerDeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHiveJsonSerDe",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) PutOpenXJsonSerDe(value *TfDeliveryStream_OpenXJsonSerDeProperty) {
	if err := t.validatePutOpenXJsonSerDeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOpenXJsonSerDe",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) ResetHiveJsonSerDe() {
	_jsii_.InvokeVoid(
		t,
		"resetHiveJsonSerDe",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) ResetOpenXJsonSerDe() {
	_jsii_.InvokeVoid(
		t,
		"resetOpenXJsonSerDe",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_DeserializerPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

