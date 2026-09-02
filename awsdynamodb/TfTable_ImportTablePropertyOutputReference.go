package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdynamodb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdynamodb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTable_ImportTablePropertyOutputReference interface {
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
	InputCompressionType() *string
	// Experimental.
	SetInputCompressionType(val *string)
	// Experimental.
	InputCompressionTypeInput() *string
	// Experimental.
	InputFormat() *string
	// Experimental.
	SetInputFormat(val *string)
	// Experimental.
	InputFormatInput() *string
	// Experimental.
	InputFormatOptions() TfTable_InputFormatOptionsPropertyOutputReference
	// Experimental.
	InputFormatOptionsInput() *TfTable_InputFormatOptionsProperty
	// Experimental.
	InternalValue() *TfTable_ImportTableProperty
	// Experimental.
	SetInternalValue(val *TfTable_ImportTableProperty)
	// Experimental.
	S3BucketSource() TfTable_S3BucketSourcePropertyOutputReference
	// Experimental.
	S3BucketSourceInput() *TfTable_S3BucketSourceProperty
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
	PutInputFormatOptions(value *TfTable_InputFormatOptionsProperty)
	// Experimental.
	PutS3BucketSource(value *TfTable_S3BucketSourceProperty)
	// Experimental.
	ResetInputCompressionType()
	// Experimental.
	ResetInputFormatOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTable_ImportTablePropertyOutputReference
type jsiiProxy_TfTable_ImportTablePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) InputCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) InputCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) InputFormatOptions() TfTable_InputFormatOptionsPropertyOutputReference {
	var returns TfTable_InputFormatOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"inputFormatOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) InputFormatOptionsInput() *TfTable_InputFormatOptionsProperty {
	var returns *TfTable_InputFormatOptionsProperty
	_jsii_.Get(
		j,
		"inputFormatOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) InternalValue() *TfTable_ImportTableProperty {
	var returns *TfTable_ImportTableProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) S3BucketSource() TfTable_S3BucketSourcePropertyOutputReference {
	var returns TfTable_S3BucketSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"s3BucketSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) S3BucketSourceInput() *TfTable_S3BucketSourceProperty {
	var returns *TfTable_S3BucketSourceProperty
	_jsii_.Get(
		j,
		"s3BucketSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTable_ImportTablePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTable_ImportTablePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTable_ImportTablePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTable_ImportTablePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dynamodb.TfTable.ImportTablePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTable_ImportTablePropertyOutputReference_Override(t TfTable_ImportTablePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dynamodb.TfTable.ImportTablePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference)SetInputCompressionType(val *string) {
	if err := j.validateSetInputCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputCompressionType",
		val,
	)
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference)SetInternalValue(val *TfTable_ImportTableProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTable_ImportTablePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) PutInputFormatOptions(value *TfTable_InputFormatOptionsProperty) {
	if err := t.validatePutInputFormatOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputFormatOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) PutS3BucketSource(value *TfTable_S3BucketSourceProperty) {
	if err := t.validatePutS3BucketSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3BucketSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) ResetInputCompressionType() {
	_jsii_.InvokeVoid(
		t,
		"resetInputCompressionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) ResetInputFormatOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetInputFormatOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTable_ImportTablePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

