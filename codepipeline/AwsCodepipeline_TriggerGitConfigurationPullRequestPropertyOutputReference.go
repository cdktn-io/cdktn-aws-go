package codepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/codepipeline/jsii"

	"github.com/cdktn-io/cdktn-aws-go/codepipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Branches() AwsCodepipeline_TriggerGitConfigurationPullRequestBranchesPropertyOutputReference
	// Experimental.
	BranchesInput() *AwsCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty
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
	FilePaths() AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference
	// Experimental.
	FilePathsInput() *AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty
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
	PutBranches(value *AwsCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty)
	// Experimental.
	PutFilePaths(value *AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty)
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

// The jsii proxy struct for AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference
type jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) Branches() AwsCodepipeline_TriggerGitConfigurationPullRequestBranchesPropertyOutputReference {
	var returns AwsCodepipeline_TriggerGitConfigurationPullRequestBranchesPropertyOutputReference
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) BranchesInput() *AwsCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty {
	var returns *AwsCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty
	_jsii_.Get(
		j,
		"branchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) Events() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) EventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) FilePaths() AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference {
	var returns AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference
	_jsii_.Get(
		j,
		"filePaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) FilePathsInput() *AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty {
	var returns *AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty
	_jsii_.Get(
		j,
		"filePathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codepipeline.AwsCodepipeline.TriggerGitConfigurationPullRequestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference_Override(a AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codepipeline.AwsCodepipeline.TriggerGitConfigurationPullRequestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetEvents(val *[]*string) {
	if err := j.validateSetEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) PutBranches(value *AwsCodepipeline_TriggerGitConfigurationPullRequestBranchesProperty) {
	if err := a.validatePutBranchesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBranches",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) PutFilePaths(value *AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty) {
	if err := a.validatePutFilePathsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilePaths",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ResetBranches() {
	_jsii_.InvokeVoid(
		a,
		"resetBranches",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ResetEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ResetFilePaths() {
	_jsii_.InvokeVoid(
		a,
		"resetFilePaths",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

