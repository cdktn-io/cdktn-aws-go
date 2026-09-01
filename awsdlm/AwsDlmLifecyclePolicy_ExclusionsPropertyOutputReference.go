package awsdlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference interface {
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
	ExcludeBootVolumes() interface{}
	// Experimental.
	SetExcludeBootVolumes(val interface{})
	// Experimental.
	ExcludeBootVolumesInput() interface{}
	// Experimental.
	ExcludeTags() *map[string]*string
	// Experimental.
	SetExcludeTags(val *map[string]*string)
	// Experimental.
	ExcludeTagsInput() *map[string]*string
	// Experimental.
	ExcludeVolumeTypes() *[]*string
	// Experimental.
	SetExcludeVolumeTypes(val *[]*string)
	// Experimental.
	ExcludeVolumeTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDlmLifecyclePolicy_ExclusionsProperty
	// Experimental.
	SetInternalValue(val *AwsDlmLifecyclePolicy_ExclusionsProperty)
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
	ResetExcludeBootVolumes()
	// Experimental.
	ResetExcludeTags()
	// Experimental.
	ResetExcludeVolumeTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference
type jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ExcludeBootVolumes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeBootVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ExcludeBootVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeBootVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ExcludeTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"excludeTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ExcludeTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"excludeTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ExcludeVolumeTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeVolumeTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ExcludeVolumeTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeVolumeTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) InternalValue() *AwsDlmLifecyclePolicy_ExclusionsProperty {
	var returns *AwsDlmLifecyclePolicy_ExclusionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDlmLifecyclePolicy_ExclusionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsDlmLifecyclePolicy.ExclusionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference_Override(a AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsDlmLifecyclePolicy.ExclusionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference)SetExcludeBootVolumes(val interface{}) {
	if err := j.validateSetExcludeBootVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeBootVolumes",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference)SetExcludeTags(val *map[string]*string) {
	if err := j.validateSetExcludeTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeTags",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference)SetExcludeVolumeTypes(val *[]*string) {
	if err := j.validateSetExcludeVolumeTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeVolumeTypes",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference)SetInternalValue(val *AwsDlmLifecyclePolicy_ExclusionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ResetExcludeBootVolumes() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludeBootVolumes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ResetExcludeTags() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludeTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ResetExcludeVolumeTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludeVolumeTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDlmLifecyclePolicy_ExclusionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

