package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodebuild/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodebuild/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfProject_BuildBatchConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CombineArtifacts() interface{}
	// Experimental.
	SetCombineArtifacts(val interface{})
	// Experimental.
	CombineArtifactsInput() interface{}
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
	InternalValue() *TfProject_BuildBatchConfigProperty
	// Experimental.
	SetInternalValue(val *TfProject_BuildBatchConfigProperty)
	// Experimental.
	Restrictions() TfProject_RestrictionsPropertyOutputReference
	// Experimental.
	RestrictionsInput() *TfProject_RestrictionsProperty
	// Experimental.
	ServiceRole() *string
	// Experimental.
	SetServiceRole(val *string)
	// Experimental.
	ServiceRoleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutInMins() *float64
	// Experimental.
	SetTimeoutInMins(val *float64)
	// Experimental.
	TimeoutInMinsInput() *float64
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
	PutRestrictions(value *TfProject_RestrictionsProperty)
	// Experimental.
	ResetCombineArtifacts()
	// Experimental.
	ResetRestrictions()
	// Experimental.
	ResetTimeoutInMins()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfProject_BuildBatchConfigPropertyOutputReference
type jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) CombineArtifacts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"combineArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) CombineArtifactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"combineArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) InternalValue() *TfProject_BuildBatchConfigProperty {
	var returns *TfProject_BuildBatchConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) Restrictions() TfProject_RestrictionsPropertyOutputReference {
	var returns TfProject_RestrictionsPropertyOutputReference
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) RestrictionsInput() *TfProject_RestrictionsProperty {
	var returns *TfProject_RestrictionsProperty
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) TimeoutInMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) TimeoutInMinsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfProject_BuildBatchConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfProject_BuildBatchConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfProject_BuildBatchConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codebuild.TfProject.BuildBatchConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfProject_BuildBatchConfigPropertyOutputReference_Override(t TfProject_BuildBatchConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codebuild.TfProject.BuildBatchConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference)SetCombineArtifacts(val interface{}) {
	if err := j.validateSetCombineArtifactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"combineArtifacts",
		val,
	)
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference)SetInternalValue(val *TfProject_BuildBatchConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference)SetTimeoutInMins(val *float64) {
	if err := j.validateSetTimeoutInMinsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInMins",
		val,
	)
}

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) PutRestrictions(value *TfProject_RestrictionsProperty) {
	if err := t.validatePutRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) ResetCombineArtifacts() {
	_jsii_.InvokeVoid(
		t,
		"resetCombineArtifacts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) ResetRestrictions() {
	_jsii_.InvokeVoid(
		t,
		"resetRestrictions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) ResetTimeoutInMins() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeoutInMins",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfProject_BuildBatchConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

