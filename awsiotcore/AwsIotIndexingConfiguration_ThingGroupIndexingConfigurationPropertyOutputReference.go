package awsiotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiotcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference interface {
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
	CustomField() AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationCustomFieldPropertyList
	// Experimental.
	CustomFieldInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationProperty)
	// Experimental.
	ManagedField() AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationManagedFieldPropertyList
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

// The jsii proxy struct for AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference
type jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) CustomField() AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationCustomFieldPropertyList {
	var returns AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationCustomFieldPropertyList
	_jsii_.Get(
		j,
		"customField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) CustomFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) InternalValue() *AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationProperty {
	var returns *AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ManagedField() AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationManagedFieldPropertyList {
	var returns AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationManagedFieldPropertyList
	_jsii_.Get(
		j,
		"managedField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ManagedFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ThingGroupIndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingGroupIndexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ThingGroupIndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingGroupIndexingModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsIotIndexingConfiguration.ThingGroupIndexingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference_Override(a AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsIotIndexingConfiguration.ThingGroupIndexingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetInternalValue(val *AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference)SetThingGroupIndexingMode(val *string) {
	if err := j.validateSetThingGroupIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thingGroupIndexingMode",
		val,
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) PutCustomField(value interface{}) {
	if err := a.validatePutCustomFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomField",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) PutManagedField(value interface{}) {
	if err := a.validatePutManagedFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedField",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ResetCustomField() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ResetManagedField() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

