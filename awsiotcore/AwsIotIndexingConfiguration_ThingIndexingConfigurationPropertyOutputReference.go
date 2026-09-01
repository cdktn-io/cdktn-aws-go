package awsiotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiotcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference interface {
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
	CustomField() AwsIotIndexingConfiguration_ThingIndexingConfigurationCustomFieldPropertyList
	// Experimental.
	CustomFieldInput() interface{}
	// Experimental.
	DeviceDefenderIndexingMode() *string
	// Experimental.
	SetDeviceDefenderIndexingMode(val *string)
	// Experimental.
	DeviceDefenderIndexingModeInput() *string
	// Experimental.
	Filter() AwsIotIndexingConfiguration_FilterPropertyOutputReference
	// Experimental.
	FilterInput() *AwsIotIndexingConfiguration_FilterProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsIotIndexingConfiguration_ThingIndexingConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsIotIndexingConfiguration_ThingIndexingConfigurationProperty)
	// Experimental.
	ManagedField() AwsIotIndexingConfiguration_ThingIndexingConfigurationManagedFieldPropertyList
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
	PutFilter(value *AwsIotIndexingConfiguration_FilterProperty)
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

// The jsii proxy struct for AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference
type jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) CustomField() AwsIotIndexingConfiguration_ThingIndexingConfigurationCustomFieldPropertyList {
	var returns AwsIotIndexingConfiguration_ThingIndexingConfigurationCustomFieldPropertyList
	_jsii_.Get(
		j,
		"customField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) CustomFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) DeviceDefenderIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceDefenderIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) DeviceDefenderIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceDefenderIndexingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) Filter() AwsIotIndexingConfiguration_FilterPropertyOutputReference {
	var returns AwsIotIndexingConfiguration_FilterPropertyOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) FilterInput() *AwsIotIndexingConfiguration_FilterProperty {
	var returns *AwsIotIndexingConfiguration_FilterProperty
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) InternalValue() *AwsIotIndexingConfiguration_ThingIndexingConfigurationProperty {
	var returns *AwsIotIndexingConfiguration_ThingIndexingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ManagedField() AwsIotIndexingConfiguration_ThingIndexingConfigurationManagedFieldPropertyList {
	var returns AwsIotIndexingConfiguration_ThingIndexingConfigurationManagedFieldPropertyList
	_jsii_.Get(
		j,
		"managedField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ManagedFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) NamedShadowIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namedShadowIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) NamedShadowIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namedShadowIndexingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ThingConnectivityIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingConnectivityIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ThingConnectivityIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingConnectivityIndexingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ThingIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ThingIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingIndexingModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsIotIndexingConfiguration.ThingIndexingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference_Override(a AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsIotIndexingConfiguration.ThingIndexingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetDeviceDefenderIndexingMode(val *string) {
	if err := j.validateSetDeviceDefenderIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceDefenderIndexingMode",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetInternalValue(val *AwsIotIndexingConfiguration_ThingIndexingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetNamedShadowIndexingMode(val *string) {
	if err := j.validateSetNamedShadowIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namedShadowIndexingMode",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetThingConnectivityIndexingMode(val *string) {
	if err := j.validateSetThingConnectivityIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thingConnectivityIndexingMode",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference)SetThingIndexingMode(val *string) {
	if err := j.validateSetThingIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thingIndexingMode",
		val,
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) PutCustomField(value interface{}) {
	if err := a.validatePutCustomFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomField",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) PutFilter(value *AwsIotIndexingConfiguration_FilterProperty) {
	if err := a.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) PutManagedField(value interface{}) {
	if err := a.validatePutManagedFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedField",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetCustomField() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetDeviceDefenderIndexingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDeviceDefenderIndexingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetManagedField() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetNamedShadowIndexingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetNamedShadowIndexingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ResetThingConnectivityIndexingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetThingConnectivityIndexingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingIndexingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

