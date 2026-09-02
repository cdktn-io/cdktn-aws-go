package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodepipeline/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodepipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCodepipeline_TriggerGitConfigurationPropertyOutputReference interface {
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
	InternalValue() *TfCodepipeline_TriggerGitConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfCodepipeline_TriggerGitConfigurationProperty)
	// Experimental.
	PullRequest() TfCodepipeline_TriggerGitConfigurationPullRequestPropertyList
	// Experimental.
	PullRequestInput() interface{}
	// Experimental.
	Push() TfCodepipeline_TriggerGitConfigurationPushPropertyList
	// Experimental.
	PushInput() interface{}
	// Experimental.
	SourceActionName() *string
	// Experimental.
	SetSourceActionName(val *string)
	// Experimental.
	SourceActionNameInput() *string
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
	PutPullRequest(value interface{})
	// Experimental.
	PutPush(value interface{})
	// Experimental.
	ResetPullRequest()
	// Experimental.
	ResetPush()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCodepipeline_TriggerGitConfigurationPropertyOutputReference
type jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) InternalValue() *TfCodepipeline_TriggerGitConfigurationProperty {
	var returns *TfCodepipeline_TriggerGitConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) PullRequest() TfCodepipeline_TriggerGitConfigurationPullRequestPropertyList {
	var returns TfCodepipeline_TriggerGitConfigurationPullRequestPropertyList
	_jsii_.Get(
		j,
		"pullRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) PullRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pullRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) Push() TfCodepipeline_TriggerGitConfigurationPushPropertyList {
	var returns TfCodepipeline_TriggerGitConfigurationPushPropertyList
	_jsii_.Get(
		j,
		"push",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) PushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) SourceActionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceActionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) SourceActionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceActionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCodepipeline_TriggerGitConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCodepipeline_TriggerGitConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCodepipeline_TriggerGitConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCodepipeline.TriggerGitConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCodepipeline_TriggerGitConfigurationPropertyOutputReference_Override(t TfCodepipeline_TriggerGitConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCodepipeline.TriggerGitConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference)SetInternalValue(val *TfCodepipeline_TriggerGitConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference)SetSourceActionName(val *string) {
	if err := j.validateSetSourceActionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceActionName",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) PutPullRequest(value interface{}) {
	if err := t.validatePutPullRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPullRequest",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) PutPush(value interface{}) {
	if err := t.validatePutPushParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPush",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) ResetPullRequest() {
	_jsii_.InvokeVoid(
		t,
		"resetPullRequest",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) ResetPush() {
	_jsii_.InvokeVoid(
		t,
		"resetPush",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

