package config

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/config/jsii"

	"github.com/cdktn-io/cdktn-aws-go/config/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConfigurationRecorder_RecordingGroupPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllSupported() interface{}
	// Experimental.
	SetAllSupported(val interface{})
	// Experimental.
	AllSupportedInput() interface{}
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
	ExclusionByResourceTypes() AwsConfigurationRecorder_ExclusionByResourceTypesPropertyList
	// Experimental.
	ExclusionByResourceTypesInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludeGlobalResourceTypes() interface{}
	// Experimental.
	SetIncludeGlobalResourceTypes(val interface{})
	// Experimental.
	IncludeGlobalResourceTypesInput() interface{}
	// Experimental.
	InternalValue() *AwsConfigurationRecorder_RecordingGroupProperty
	// Experimental.
	SetInternalValue(val *AwsConfigurationRecorder_RecordingGroupProperty)
	// Experimental.
	RecordingStrategy() AwsConfigurationRecorder_RecordingStrategyPropertyList
	// Experimental.
	RecordingStrategyInput() interface{}
	// Experimental.
	ResourceTypes() *[]*string
	// Experimental.
	SetResourceTypes(val *[]*string)
	// Experimental.
	ResourceTypesInput() *[]*string
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
	PutExclusionByResourceTypes(value interface{})
	// Experimental.
	PutRecordingStrategy(value interface{})
	// Experimental.
	ResetAllSupported()
	// Experimental.
	ResetExclusionByResourceTypes()
	// Experimental.
	ResetIncludeGlobalResourceTypes()
	// Experimental.
	ResetRecordingStrategy()
	// Experimental.
	ResetResourceTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsConfigurationRecorder_RecordingGroupPropertyOutputReference
type jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) AllSupported() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) AllSupportedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allSupportedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ExclusionByResourceTypes() AwsConfigurationRecorder_ExclusionByResourceTypesPropertyList {
	var returns AwsConfigurationRecorder_ExclusionByResourceTypesPropertyList
	_jsii_.Get(
		j,
		"exclusionByResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ExclusionByResourceTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exclusionByResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) IncludeGlobalResourceTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) IncludeGlobalResourceTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) InternalValue() *AwsConfigurationRecorder_RecordingGroupProperty {
	var returns *AwsConfigurationRecorder_RecordingGroupProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) RecordingStrategy() AwsConfigurationRecorder_RecordingStrategyPropertyList {
	var returns AwsConfigurationRecorder_RecordingStrategyPropertyList
	_jsii_.Get(
		j,
		"recordingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) RecordingStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConfigurationRecorder_RecordingGroupPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConfigurationRecorder_RecordingGroupPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConfigurationRecorder_RecordingGroupPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-config.AwsConfigurationRecorder.RecordingGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConfigurationRecorder_RecordingGroupPropertyOutputReference_Override(a AwsConfigurationRecorder_RecordingGroupPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-config.AwsConfigurationRecorder.RecordingGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference)SetAllSupported(val interface{}) {
	if err := j.validateSetAllSupportedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allSupported",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference)SetIncludeGlobalResourceTypes(val interface{}) {
	if err := j.validateSetIncludeGlobalResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeGlobalResourceTypes",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference)SetInternalValue(val *AwsConfigurationRecorder_RecordingGroupProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference)SetResourceTypes(val *[]*string) {
	if err := j.validateSetResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) PutExclusionByResourceTypes(value interface{}) {
	if err := a.validatePutExclusionByResourceTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExclusionByResourceTypes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) PutRecordingStrategy(value interface{}) {
	if err := a.validatePutRecordingStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRecordingStrategy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ResetAllSupported() {
	_jsii_.InvokeVoid(
		a,
		"resetAllSupported",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ResetExclusionByResourceTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetExclusionByResourceTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ResetIncludeGlobalResourceTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeGlobalResourceTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ResetRecordingStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetRecordingStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ResetResourceTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConfigurationRecorder_RecordingGroupPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

