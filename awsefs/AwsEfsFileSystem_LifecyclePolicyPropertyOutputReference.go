package awsefs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsefs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsefs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransitionToArchive() *string
	// Experimental.
	SetTransitionToArchive(val *string)
	// Experimental.
	TransitionToArchiveInput() *string
	// Experimental.
	TransitionToIa() *string
	// Experimental.
	SetTransitionToIa(val *string)
	// Experimental.
	TransitionToIaInput() *string
	// Experimental.
	TransitionToPrimaryStorageClass() *string
	// Experimental.
	SetTransitionToPrimaryStorageClass(val *string)
	// Experimental.
	TransitionToPrimaryStorageClassInput() *string
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
	ResetTransitionToArchive()
	// Experimental.
	ResetTransitionToIa()
	// Experimental.
	ResetTransitionToPrimaryStorageClass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference
type jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) TransitionToArchive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToArchive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) TransitionToArchiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToArchiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) TransitionToIa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToIa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) TransitionToIaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToIaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) TransitionToPrimaryStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToPrimaryStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) TransitionToPrimaryStorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToPrimaryStorageClassInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEfsFileSystem_LifecyclePolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEfsFileSystem_LifecyclePolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-efs.AwsEfsFileSystem.LifecyclePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEfsFileSystem_LifecyclePolicyPropertyOutputReference_Override(a AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-efs.AwsEfsFileSystem.LifecyclePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference)SetTransitionToArchive(val *string) {
	if err := j.validateSetTransitionToArchiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitionToArchive",
		val,
	)
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference)SetTransitionToIa(val *string) {
	if err := j.validateSetTransitionToIaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitionToIa",
		val,
	)
}

func (j *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference)SetTransitionToPrimaryStorageClass(val *string) {
	if err := j.validateSetTransitionToPrimaryStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitionToPrimaryStorageClass",
		val,
	)
}

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) ResetTransitionToArchive() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitionToArchive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) ResetTransitionToIa() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitionToIa",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) ResetTransitionToPrimaryStorageClass() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitionToPrimaryStorageClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEfsFileSystem_LifecyclePolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

