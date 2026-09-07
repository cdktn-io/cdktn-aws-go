package dlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/dlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference interface {
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
	ExcludeBootVolume() interface{}
	// Experimental.
	SetExcludeBootVolume(val interface{})
	// Experimental.
	ExcludeBootVolumeInput() interface{}
	// Experimental.
	ExcludeDataVolumeTags() *map[string]*string
	// Experimental.
	SetExcludeDataVolumeTags(val *map[string]*string)
	// Experimental.
	ExcludeDataVolumeTagsInput() *map[string]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsLifecyclePolicy_PolicyDetailsParametersProperty
	// Experimental.
	SetInternalValue(val *AwsLifecyclePolicy_PolicyDetailsParametersProperty)
	// Experimental.
	NoReboot() interface{}
	// Experimental.
	SetNoReboot(val interface{})
	// Experimental.
	NoRebootInput() interface{}
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
	ResetExcludeBootVolume()
	// Experimental.
	ResetExcludeDataVolumeTags()
	// Experimental.
	ResetNoReboot()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference
type jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) ExcludeBootVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeBootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) ExcludeBootVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeBootVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) ExcludeDataVolumeTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"excludeDataVolumeTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) ExcludeDataVolumeTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"excludeDataVolumeTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) InternalValue() *AwsLifecyclePolicy_PolicyDetailsParametersProperty {
	var returns *AwsLifecyclePolicy_PolicyDetailsParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) NoReboot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noReboot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) NoRebootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRebootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsLifecyclePolicy.PolicyDetailsParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference_Override(a AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsLifecyclePolicy.PolicyDetailsParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference)SetExcludeBootVolume(val interface{}) {
	if err := j.validateSetExcludeBootVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeBootVolume",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference)SetExcludeDataVolumeTags(val *map[string]*string) {
	if err := j.validateSetExcludeDataVolumeTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeDataVolumeTags",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference)SetInternalValue(val *AwsLifecyclePolicy_PolicyDetailsParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference)SetNoReboot(val interface{}) {
	if err := j.validateSetNoRebootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noReboot",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) ResetExcludeBootVolume() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludeBootVolume",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) ResetExcludeDataVolumeTags() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludeDataVolumeTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) ResetNoReboot() {
	_jsii_.InvokeVoid(
		a,
		"resetNoReboot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLifecyclePolicy_PolicyDetailsParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

