package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodepipeline/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodepipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Branches() TfCodepipeline_TriggerGitConfigurationPullRequestBranchesPropertyOutputReference
	// Experimental.
	BranchesInput() *TfCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty
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
	Events() *[]*string
	// Experimental.
	SetEvents(val *[]*string)
	// Experimental.
	EventsInput() *[]*string
	// Experimental.
	FilePaths() TfCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference
	// Experimental.
	FilePathsInput() *TfCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty
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
	PutBranches(value *TfCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty)
	// Experimental.
	PutFilePaths(value *TfCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty)
	// Experimental.
	ResetBranches()
	// Experimental.
	ResetEvents()
	// Experimental.
	ResetFilePaths()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference
type jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) Branches() TfCodepipeline_TriggerGitConfigurationPullRequestBranchesPropertyOutputReference {
	var returns TfCodepipeline_TriggerGitConfigurationPullRequestBranchesPropertyOutputReference
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) BranchesInput() *TfCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty {
	var returns *TfCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty
	_jsii_.Get(
		j,
		"branchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) Events() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) EventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) FilePaths() TfCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference {
	var returns TfCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference
	_jsii_.Get(
		j,
		"filePaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) FilePathsInput() *TfCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty {
	var returns *TfCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty
	_jsii_.Get(
		j,
		"filePathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCodepipeline.TriggerGitConfigurationPullRequestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference_Override(t TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codepipeline.TfCodepipeline.TriggerGitConfigurationPullRequestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetEvents(val *[]*string) {
	if err := j.validateSetEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) PutBranches(value *TfCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty) {
	if err := t.validatePutBranchesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBranches",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) PutFilePaths(value *TfCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty) {
	if err := t.validatePutFilePathsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFilePaths",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ResetBranches() {
	_jsii_.InvokeVoid(
		t,
		"resetBranches",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ResetEvents() {
	_jsii_.InvokeVoid(
		t,
		"resetEvents",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ResetFilePaths() {
	_jsii_.InvokeVoid(
		t,
		"resetFilePaths",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

