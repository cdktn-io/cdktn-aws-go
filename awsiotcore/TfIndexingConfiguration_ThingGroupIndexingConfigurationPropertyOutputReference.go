package awsiotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiotcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference interface {
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
	CustomField() TfIndexingConfiguration_ThingGroupIndexingConfigurationCustomFieldPropertyList
	// Experimental.
	CustomFieldInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfIndexingConfiguration_ThingGroupIndexingConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfIndexingConfiguration_ThingGroupIndexingConfigurationProperty)
	// Experimental.
	ManagedField() TfIndexingConfiguration_ThingGroupIndexingConfigurationManagedFieldPropertyList
	// Experimental.
	ManagedFieldInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ThingGroupIndexingMode() *string
	// Experimental.
	SetThingGroupIndexingMode(val *string)
	// Experimental.
	ThingGroupIndexingModeInput() *string
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
	PutCustomField(value interface{})
	// Experimental.
	PutManagedField(value interface{})
	// Experimental.
	ResetCustomField()
	// Experimental.
	ResetManagedField()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference
type jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) CustomField() TfIndexingConfiguration_ThingGroupIndexingConfigurationCustomFieldPropertyList {
	var returns TfIndexingConfiguration_ThingGroupIndexingConfigurationCustomFieldPropertyList
	_jsii_.Get(
		j,
		"customField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) CustomFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) InternalValue() *TfIndexingConfiguration_ThingGroupIndexingConfigurationProperty {
	var returns *TfIndexingConfiguration_ThingGroupIndexingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ManagedField() TfIndexingConfiguration_ThingGroupIndexingConfigurationManagedFieldPropertyList {
	var returns TfIndexingConfiguration_ThingGroupIndexingConfigurationManagedFieldPropertyList
	_jsii_.Get(
		j,
		"managedField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ManagedFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ThingGroupIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingGroupIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ThingGroupIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingGroupIndexingModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfIndexingConfiguration.ThingGroupIndexingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference_Override(t TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfIndexingConfiguration.ThingGroupIndexingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetInternalValue(val *TfIndexingConfiguration_ThingGroupIndexingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetThingGroupIndexingMode(val *string) {
	if err := j.validateSetThingGroupIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thingGroupIndexingMode",
		val,
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) PutCustomField(value interface{}) {
	if err := t.validatePutCustomFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomField",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) PutManagedField(value interface{}) {
	if err := t.validatePutManagedFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedField",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ResetCustomField() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomField",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ResetManagedField() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedField",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

