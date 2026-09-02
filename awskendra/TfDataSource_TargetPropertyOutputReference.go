package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSource_TargetPropertyOutputReference interface {
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
	InternalValue() *TfDataSource_TargetProperty
	// Experimental.
	SetInternalValue(val *TfDataSource_TargetProperty)
	// Experimental.
	TargetDocumentAttributeKey() *string
	// Experimental.
	SetTargetDocumentAttributeKey(val *string)
	// Experimental.
	TargetDocumentAttributeKeyInput() *string
	// Experimental.
	TargetDocumentAttributeValue() TfDataSource_TargetDocumentAttributeValuePropertyOutputReference
	// Experimental.
	TargetDocumentAttributeValueDeletion() interface{}
	// Experimental.
	SetTargetDocumentAttributeValueDeletion(val interface{})
	// Experimental.
	TargetDocumentAttributeValueDeletionInput() interface{}
	// Experimental.
	TargetDocumentAttributeValueInput() *TfDataSource_TargetDocumentAttributeValueProperty
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
	PutTargetDocumentAttributeValue(value *TfDataSource_TargetDocumentAttributeValueProperty)
	// Experimental.
	ResetTargetDocumentAttributeKey()
	// Experimental.
	ResetTargetDocumentAttributeValue()
	// Experimental.
	ResetTargetDocumentAttributeValueDeletion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSource_TargetPropertyOutputReference
type jsiiProxy_TfDataSource_TargetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) InternalValue() *TfDataSource_TargetProperty {
	var returns *TfDataSource_TargetProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) TargetDocumentAttributeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDocumentAttributeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) TargetDocumentAttributeKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDocumentAttributeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) TargetDocumentAttributeValue() TfDataSource_TargetDocumentAttributeValuePropertyOutputReference {
	var returns TfDataSource_TargetDocumentAttributeValuePropertyOutputReference
	_jsii_.Get(
		j,
		"targetDocumentAttributeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) TargetDocumentAttributeValueDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetDocumentAttributeValueDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) TargetDocumentAttributeValueDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetDocumentAttributeValueDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) TargetDocumentAttributeValueInput() *TfDataSource_TargetDocumentAttributeValueProperty {
	var returns *TfDataSource_TargetDocumentAttributeValueProperty
	_jsii_.Get(
		j,
		"targetDocumentAttributeValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSource_TargetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDataSource_TargetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSource_TargetPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSource_TargetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.TfDataSource.TargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSource_TargetPropertyOutputReference_Override(t TfDataSource_TargetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.TfDataSource.TargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference)SetInternalValue(val *TfDataSource_TargetProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference)SetTargetDocumentAttributeKey(val *string) {
	if err := j.validateSetTargetDocumentAttributeKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDocumentAttributeKey",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference)SetTargetDocumentAttributeValueDeletion(val interface{}) {
	if err := j.validateSetTargetDocumentAttributeValueDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDocumentAttributeValueDeletion",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_TargetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) PutTargetDocumentAttributeValue(value *TfDataSource_TargetDocumentAttributeValueProperty) {
	if err := t.validatePutTargetDocumentAttributeValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetDocumentAttributeValue",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) ResetTargetDocumentAttributeKey() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetDocumentAttributeKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) ResetTargetDocumentAttributeValue() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetDocumentAttributeValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) ResetTargetDocumentAttributeValueDeletion() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetDocumentAttributeValueDeletion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSource_TargetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

