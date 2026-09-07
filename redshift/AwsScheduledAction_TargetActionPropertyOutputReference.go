package redshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/redshift/jsii"

	"github.com/cdktn-io/cdktn-aws-go/redshift/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsScheduledAction_TargetActionPropertyOutputReference interface {
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
	InternalValue() *AwsScheduledAction_TargetActionProperty
	// Experimental.
	SetInternalValue(val *AwsScheduledAction_TargetActionProperty)
	// Experimental.
	PauseCluster() AwsScheduledAction_PauseClusterPropertyOutputReference
	// Experimental.
	PauseClusterInput() *AwsScheduledAction_PauseClusterProperty
	// Experimental.
	ResizeCluster() AwsScheduledAction_ResizeClusterPropertyOutputReference
	// Experimental.
	ResizeClusterInput() *AwsScheduledAction_ResizeClusterProperty
	// Experimental.
	ResumeCluster() AwsScheduledAction_ResumeClusterPropertyOutputReference
	// Experimental.
	ResumeClusterInput() *AwsScheduledAction_ResumeClusterProperty
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
	PutPauseCluster(value *AwsScheduledAction_PauseClusterProperty)
	// Experimental.
	PutResizeCluster(value *AwsScheduledAction_ResizeClusterProperty)
	// Experimental.
	PutResumeCluster(value *AwsScheduledAction_ResumeClusterProperty)
	// Experimental.
	ResetPauseCluster()
	// Experimental.
	ResetResizeCluster()
	// Experimental.
	ResetResumeCluster()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsScheduledAction_TargetActionPropertyOutputReference
type jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) InternalValue() *AwsScheduledAction_TargetActionProperty {
	var returns *AwsScheduledAction_TargetActionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) PauseCluster() AwsScheduledAction_PauseClusterPropertyOutputReference {
	var returns AwsScheduledAction_PauseClusterPropertyOutputReference
	_jsii_.Get(
		j,
		"pauseCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) PauseClusterInput() *AwsScheduledAction_PauseClusterProperty {
	var returns *AwsScheduledAction_PauseClusterProperty
	_jsii_.Get(
		j,
		"pauseClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) ResizeCluster() AwsScheduledAction_ResizeClusterPropertyOutputReference {
	var returns AwsScheduledAction_ResizeClusterPropertyOutputReference
	_jsii_.Get(
		j,
		"resizeCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) ResizeClusterInput() *AwsScheduledAction_ResizeClusterProperty {
	var returns *AwsScheduledAction_ResizeClusterProperty
	_jsii_.Get(
		j,
		"resizeClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) ResumeCluster() AwsScheduledAction_ResumeClusterPropertyOutputReference {
	var returns AwsScheduledAction_ResumeClusterPropertyOutputReference
	_jsii_.Get(
		j,
		"resumeCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) ResumeClusterInput() *AwsScheduledAction_ResumeClusterProperty {
	var returns *AwsScheduledAction_ResumeClusterProperty
	_jsii_.Get(
		j,
		"resumeClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsScheduledAction_TargetActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsScheduledAction_TargetActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsScheduledAction_TargetActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-redshift.AwsScheduledAction.TargetActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsScheduledAction_TargetActionPropertyOutputReference_Override(a AwsScheduledAction_TargetActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-redshift.AwsScheduledAction.TargetActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference)SetInternalValue(val *AwsScheduledAction_TargetActionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) PutPauseCluster(value *AwsScheduledAction_PauseClusterProperty) {
	if err := a.validatePutPauseClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPauseCluster",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) PutResizeCluster(value *AwsScheduledAction_ResizeClusterProperty) {
	if err := a.validatePutResizeClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResizeCluster",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) PutResumeCluster(value *AwsScheduledAction_ResumeClusterProperty) {
	if err := a.validatePutResumeClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResumeCluster",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) ResetPauseCluster() {
	_jsii_.InvokeVoid(
		a,
		"resetPauseCluster",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) ResetResizeCluster() {
	_jsii_.InvokeVoid(
		a,
		"resetResizeCluster",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) ResetResumeCluster() {
	_jsii_.InvokeVoid(
		a,
		"resetResumeCluster",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsScheduledAction_TargetActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

