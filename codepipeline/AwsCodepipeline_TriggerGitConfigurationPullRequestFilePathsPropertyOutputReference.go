package codepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/codepipeline/jsii"

	"github.com/cdktn-io/cdktn-aws-go/codepipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference interface {
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
	Excludes() *[]*string
	// Experimental.
	SetExcludes(val *[]*string)
	// Experimental.
	ExcludesInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	Includes() *[]*string
	// Experimental.
	SetIncludes(val *[]*string)
	// Experimental.
	IncludesInput() *[]*string
	// Experimental.
	InternalValue() *AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty
	// Experimental.
	SetInternalValue(val *AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty)
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
	ResetExcludes()
	// Experimental.
	ResetIncludes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference
type jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) Excludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) ExcludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) Includes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) IncludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) InternalValue() *AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty {
	var returns *AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codepipeline.AwsCodepipeline.TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference_Override(a AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codepipeline.AwsCodepipeline.TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference)SetExcludes(val *[]*string) {
	if err := j.validateSetExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludes",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference)SetIncludes(val *[]*string) {
	if err := j.validateSetIncludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includes",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference)SetInternalValue(val *AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) ResetExcludes() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) ResetIncludes() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCodepipeline_TriggerGitConfigurationPullRequestFilePathsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

