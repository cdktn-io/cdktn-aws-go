package timestreamwrite

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/timestreamwrite/jsii"

	"github.com/cdktn-io/cdktn-aws-go/timestreamwrite/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTable_MagneticStoreWritePropertiesPropertyOutputReference interface {
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
	InternalValue() *AwsTable_MagneticStoreWritePropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsTable_MagneticStoreWritePropertiesProperty)
	// Experimental.
	MagneticStoreRejectedDataLocation() AwsTable_MagneticStoreRejectedDataLocationPropertyOutputReference
	// Experimental.
	MagneticStoreRejectedDataLocationInput() *AwsTable_MagneticStoreRejectedDataLocationProperty
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
	PutMagneticStoreRejectedDataLocation(value *AwsTable_MagneticStoreRejectedDataLocationProperty)
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

// The jsii proxy struct for AwsTable_MagneticStoreWritePropertiesPropertyOutputReference
type jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) EnableMagneticStoreWrites() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMagneticStoreWrites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) EnableMagneticStoreWritesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMagneticStoreWritesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) InternalValue() *AwsTable_MagneticStoreWritePropertiesProperty {
	var returns *AwsTable_MagneticStoreWritePropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) MagneticStoreRejectedDataLocation() AwsTable_MagneticStoreRejectedDataLocationPropertyOutputReference {
	var returns AwsTable_MagneticStoreRejectedDataLocationPropertyOutputReference
	_jsii_.Get(
		j,
		"magneticStoreRejectedDataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) MagneticStoreRejectedDataLocationInput() *AwsTable_MagneticStoreRejectedDataLocationProperty {
	var returns *AwsTable_MagneticStoreRejectedDataLocationProperty
	_jsii_.Get(
		j,
		"magneticStoreRejectedDataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTable_MagneticStoreWritePropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsTable_MagneticStoreWritePropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTable_MagneticStoreWritePropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-write.AwsTable.MagneticStoreWritePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTable_MagneticStoreWritePropertiesPropertyOutputReference_Override(a AwsTable_MagneticStoreWritePropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-write.AwsTable.MagneticStoreWritePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference)SetEnableMagneticStoreWrites(val interface{}) {
	if err := j.validateSetEnableMagneticStoreWritesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMagneticStoreWrites",
		val,
	)
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference)SetInternalValue(val *AwsTable_MagneticStoreWritePropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) PutMagneticStoreRejectedDataLocation(value *AwsTable_MagneticStoreRejectedDataLocationProperty) {
	if err := a.validatePutMagneticStoreRejectedDataLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMagneticStoreRejectedDataLocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) ResetEnableMagneticStoreWrites() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableMagneticStoreWrites",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) ResetMagneticStoreRejectedDataLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetMagneticStoreRejectedDataLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTable_MagneticStoreWritePropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

