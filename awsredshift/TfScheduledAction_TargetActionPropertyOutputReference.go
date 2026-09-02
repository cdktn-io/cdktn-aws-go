package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsredshift/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsredshift/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfScheduledAction_TargetActionPropertyOutputReference interface {
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
	InternalValue() *TfScheduledAction_TargetActionProperty
	// Experimental.
	SetInternalValue(val *TfScheduledAction_TargetActionProperty)
	// Experimental.
	PauseCluster() TfScheduledAction_PauseClusterPropertyOutputReference
	// Experimental.
	PauseClusterInput() *TfScheduledAction_PauseClusterProperty
	// Experimental.
	ResizeCluster() TfScheduledAction_ResizeClusterPropertyOutputReference
	// Experimental.
	ResizeClusterInput() *TfScheduledAction_ResizeClusterProperty
	// Experimental.
	ResumeCluster() TfScheduledAction_ResumeClusterPropertyOutputReference
	// Experimental.
	ResumeClusterInput() *TfScheduledAction_ResumeClusterProperty
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
	PutPauseCluster(value *TfScheduledAction_PauseClusterProperty)
	// Experimental.
	PutResizeCluster(value *TfScheduledAction_ResizeClusterProperty)
	// Experimental.
	PutResumeCluster(value *TfScheduledAction_ResumeClusterProperty)
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

// The jsii proxy struct for TfScheduledAction_TargetActionPropertyOutputReference
type jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) InternalValue() *TfScheduledAction_TargetActionProperty {
	var returns *TfScheduledAction_TargetActionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) PauseCluster() TfScheduledAction_PauseClusterPropertyOutputReference {
	var returns TfScheduledAction_PauseClusterPropertyOutputReference
	_jsii_.Get(
		j,
		"pauseCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) PauseClusterInput() *TfScheduledAction_PauseClusterProperty {
	var returns *TfScheduledAction_PauseClusterProperty
	_jsii_.Get(
		j,
		"pauseClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) ResizeCluster() TfScheduledAction_ResizeClusterPropertyOutputReference {
	var returns TfScheduledAction_ResizeClusterPropertyOutputReference
	_jsii_.Get(
		j,
		"resizeCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) ResizeClusterInput() *TfScheduledAction_ResizeClusterProperty {
	var returns *TfScheduledAction_ResizeClusterProperty
	_jsii_.Get(
		j,
		"resizeClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) ResumeCluster() TfScheduledAction_ResumeClusterPropertyOutputReference {
	var returns TfScheduledAction_ResumeClusterPropertyOutputReference
	_jsii_.Get(
		j,
		"resumeCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) ResumeClusterInput() *TfScheduledAction_ResumeClusterProperty {
	var returns *TfScheduledAction_ResumeClusterProperty
	_jsii_.Get(
		j,
		"resumeClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfScheduledAction_TargetActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfScheduledAction_TargetActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfScheduledAction_TargetActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-redshift.TfScheduledAction.TargetActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfScheduledAction_TargetActionPropertyOutputReference_Override(t TfScheduledAction_TargetActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-redshift.TfScheduledAction.TargetActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference)SetInternalValue(val *TfScheduledAction_TargetActionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) PutPauseCluster(value *TfScheduledAction_PauseClusterProperty) {
	if err := t.validatePutPauseClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPauseCluster",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) PutResizeCluster(value *TfScheduledAction_ResizeClusterProperty) {
	if err := t.validatePutResizeClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResizeCluster",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) PutResumeCluster(value *TfScheduledAction_ResumeClusterProperty) {
	if err := t.validatePutResumeClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResumeCluster",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) ResetPauseCluster() {
	_jsii_.InvokeVoid(
		t,
		"resetPauseCluster",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) ResetResizeCluster() {
	_jsii_.InvokeVoid(
		t,
		"resetResizeCluster",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) ResetResumeCluster() {
	_jsii_.InvokeVoid(
		t,
		"resetResumeCluster",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfScheduledAction_TargetActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

