package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference interface {
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
	InternalValue() *TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationProperty
	// Experimental.
	SetInternalValue(val *TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationProperty)
	// Experimental.
	LaunchTemplateId() *string
	// Experimental.
	SetLaunchTemplateId(val *string)
	// Experimental.
	LaunchTemplateIdInput() *string
	// Experimental.
	LaunchTemplateName() *string
	// Experimental.
	SetLaunchTemplateName(val *string)
	// Experimental.
	LaunchTemplateNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(val *string)
	// Experimental.
	VersionInput() *string
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
	ResetLaunchTemplateId()
	// Experimental.
	ResetLaunchTemplateName()
	// Experimental.
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference
type jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) InternalValue() *TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationProperty {
	var returns *TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) LaunchTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) LaunchTemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) LaunchTemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) LaunchTemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfGroup.MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference_Override(t TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfGroup.MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference)SetInternalValue(val *TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference)SetLaunchTemplateId(val *string) {
	if err := j.validateSetLaunchTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTemplateId",
		val,
	)
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference)SetLaunchTemplateName(val *string) {
	if err := j.validateSetLaunchTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTemplateName",
		val,
	)
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) ResetLaunchTemplateId() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchTemplateId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) ResetLaunchTemplateName() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchTemplateName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

