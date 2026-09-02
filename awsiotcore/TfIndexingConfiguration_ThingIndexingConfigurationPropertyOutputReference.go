package awsiotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiotcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference interface {
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
	CustomField() TfIndexingConfiguration_ThingIndexingConfigurationCustomFieldPropertyList
	// Experimental.
	CustomFieldInput() interface{}
	// Experimental.
	DeviceDefenderIndexingMode() *string
	// Experimental.
	SetDeviceDefenderIndexingMode(val *string)
	// Experimental.
	DeviceDefenderIndexingModeInput() *string
	// Experimental.
	Filter() TfIndexingConfiguration_FilterPropertyOutputReference
	// Experimental.
	FilterInput() *TfIndexingConfiguration_FilterProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfIndexingConfiguration_ThingIndexingConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfIndexingConfiguration_ThingIndexingConfigurationProperty)
	// Experimental.
	ManagedField() TfIndexingConfiguration_ThingIndexingConfigurationManagedFieldPropertyList
	// Experimental.
	ManagedFieldInput() interface{}
	// Experimental.
	NamedShadowIndexingMode() *string
	// Experimental.
	SetNamedShadowIndexingMode(val *string)
	// Experimental.
	NamedShadowIndexingModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ThingConnectivityIndexingMode() *string
	// Experimental.
	SetThingConnectivityIndexingMode(val *string)
	// Experimental.
	ThingConnectivityIndexingModeInput() *string
	// Experimental.
	ThingIndexingMode() *string
	// Experimental.
	SetThingIndexingMode(val *string)
	// Experimental.
	ThingIndexingModeInput() *string
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
	PutFilter(value *TfIndexingConfiguration_FilterProperty)
	// Experimental.
	PutManagedField(value interface{})
	// Experimental.
	ResetCustomField()
	// Experimental.
	ResetDeviceDefenderIndexingMode()
	// Experimental.
	ResetFilter()
	// Experimental.
	ResetManagedField()
	// Experimental.
	ResetNamedShadowIndexingMode()
	// Experimental.
	ResetThingConnectivityIndexingMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference
type jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) CustomField() TfIndexingConfiguration_ThingIndexingConfigurationCustomFieldPropertyList {
	var returns TfIndexingConfiguration_ThingIndexingConfigurationCustomFieldPropertyList
	_jsii_.Get(
		j,
		"customField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) CustomFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) DeviceDefenderIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceDefenderIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) DeviceDefenderIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceDefenderIndexingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) Filter() TfIndexingConfiguration_FilterPropertyOutputReference {
	var returns TfIndexingConfiguration_FilterPropertyOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) FilterInput() *TfIndexingConfiguration_FilterProperty {
	var returns *TfIndexingConfiguration_FilterProperty
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) InternalValue() *TfIndexingConfiguration_ThingIndexingConfigurationProperty {
	var returns *TfIndexingConfiguration_ThingIndexingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ManagedField() TfIndexingConfiguration_ThingIndexingConfigurationManagedFieldPropertyList {
	var returns TfIndexingConfiguration_ThingIndexingConfigurationManagedFieldPropertyList
	_jsii_.Get(
		j,
		"managedField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ManagedFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) NamedShadowIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namedShadowIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) NamedShadowIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namedShadowIndexingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ThingConnectivityIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingConnectivityIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ThingConnectivityIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingConnectivityIndexingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ThingIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ThingIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingIndexingModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfIndexingConfiguration.ThingIndexingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference_Override(t TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfIndexingConfiguration.ThingIndexingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetDeviceDefenderIndexingMode(val *string) {
	if err := j.validateSetDeviceDefenderIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceDefenderIndexingMode",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetInternalValue(val *TfIndexingConfiguration_ThingIndexingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetNamedShadowIndexingMode(val *string) {
	if err := j.validateSetNamedShadowIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namedShadowIndexingMode",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetThingConnectivityIndexingMode(val *string) {
	if err := j.validateSetThingConnectivityIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thingConnectivityIndexingMode",
		val,
	)
}

func (j *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetThingIndexingMode(val *string) {
	if err := j.validateSetThingIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thingIndexingMode",
		val,
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) PutCustomField(value interface{}) {
	if err := t.validatePutCustomFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomField",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) PutFilter(value *TfIndexingConfiguration_FilterProperty) {
	if err := t.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFilter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) PutManagedField(value interface{}) {
	if err := t.validatePutManagedFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedField",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetCustomField() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomField",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetDeviceDefenderIndexingMode() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceDefenderIndexingMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetManagedField() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedField",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetNamedShadowIndexingMode() {
	_jsii_.InvokeVoid(
		t,
		"resetNamedShadowIndexingMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetThingConnectivityIndexingMode() {
	_jsii_.InvokeVoid(
		t,
		"resetThingConnectivityIndexingMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

