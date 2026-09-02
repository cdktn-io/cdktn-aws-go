package awstimestreamwrite

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstimestreamwrite/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstimestreamwrite/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTable_MagneticStoreWritePropertiesPropertyOutputReference interface {
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
	EnableMagneticStoreWrites() interface{}
	// Experimental.
	SetEnableMagneticStoreWrites(val interface{})
	// Experimental.
	EnableMagneticStoreWritesInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfTable_MagneticStoreWritePropertiesProperty
	// Experimental.
	SetInternalValue(val *TfTable_MagneticStoreWritePropertiesProperty)
	// Experimental.
	MagneticStoreRejectedDataLocation() TfTable_MagneticStoreRejectedDataLocationPropertyOutputReference
	// Experimental.
	MagneticStoreRejectedDataLocationInput() *TfTable_MagneticStoreRejectedDataLocationProperty
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
	PutMagneticStoreRejectedDataLocation(value *TfTable_MagneticStoreRejectedDataLocationProperty)
	// Experimental.
	ResetEnableMagneticStoreWrites()
	// Experimental.
	ResetMagneticStoreRejectedDataLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTable_MagneticStoreWritePropertiesPropertyOutputReference
type jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) EnableMagneticStoreWrites() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMagneticStoreWrites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) EnableMagneticStoreWritesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMagneticStoreWritesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) InternalValue() *TfTable_MagneticStoreWritePropertiesProperty {
	var returns *TfTable_MagneticStoreWritePropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) MagneticStoreRejectedDataLocation() TfTable_MagneticStoreRejectedDataLocationPropertyOutputReference {
	var returns TfTable_MagneticStoreRejectedDataLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"magneticStoreRejectedDataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) MagneticStoreRejectedDataLocationInput() *TfTable_MagneticStoreRejectedDataLocationProperty {
	var returns *TfTable_MagneticStoreRejectedDataLocationProperty
	_jsii_.Get(
		j,
		"magneticStoreRejectedDataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTable_MagneticStoreWritePropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTable_MagneticStoreWritePropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTable_MagneticStoreWritePropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-write.TfTable.MagneticStoreWritePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTable_MagneticStoreWritePropertiesPropertyOutputReference_Override(t TfTable_MagneticStoreWritePropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-write.TfTable.MagneticStoreWritePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference)SetEnableMagneticStoreWrites(val interface{}) {
	if err := j.validateSetEnableMagneticStoreWritesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMagneticStoreWrites",
		val,
	)
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference)SetInternalValue(val *TfTable_MagneticStoreWritePropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) PutMagneticStoreRejectedDataLocation(value *TfTable_MagneticStoreRejectedDataLocationProperty) {
	if err := t.validatePutMagneticStoreRejectedDataLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMagneticStoreRejectedDataLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) ResetEnableMagneticStoreWrites() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableMagneticStoreWrites",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) ResetMagneticStoreRejectedDataLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetMagneticStoreRejectedDataLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTable_MagneticStoreWritePropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

