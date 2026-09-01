package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference interface {
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
	ExperimentName() *string
	// Experimental.
	SetExperimentName(val *string)
	// Experimental.
	ExperimentNameInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RunName() *string
	// Experimental.
	SetRunName(val *string)
	// Experimental.
	RunNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrialComponentDisplayName() *string
	// Experimental.
	SetTrialComponentDisplayName(val *string)
	// Experimental.
	TrialComponentDisplayNameInput() *string
	// Experimental.
	TrialName() *string
	// Experimental.
	SetTrialName(val *string)
	// Experimental.
	TrialNameInput() *string
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
	ResetExperimentName()
	// Experimental.
	ResetRunName()
	// Experimental.
	ResetTrialComponentDisplayName()
	// Experimental.
	ResetTrialName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference
type jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) ExperimentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"experimentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) ExperimentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"experimentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) RunName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) RunNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) TrialComponentDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trialComponentDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) TrialComponentDisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trialComponentDisplayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) TrialName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trialName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) TrialNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trialNameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerTrainingJob.ExperimentConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference_Override(a AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerTrainingJob.ExperimentConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference)SetExperimentName(val *string) {
	if err := j.validateSetExperimentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"experimentName",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference)SetRunName(val *string) {
	if err := j.validateSetRunNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runName",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference)SetTrialComponentDisplayName(val *string) {
	if err := j.validateSetTrialComponentDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trialComponentDisplayName",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference)SetTrialName(val *string) {
	if err := j.validateSetTrialNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trialName",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) ResetExperimentName() {
	_jsii_.InvokeVoid(
		a,
		"resetExperimentName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) ResetRunName() {
	_jsii_.InvokeVoid(
		a,
		"resetRunName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) ResetTrialComponentDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetTrialComponentDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) ResetTrialName() {
	_jsii_.InvokeVoid(
		a,
		"resetTrialName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob_ExperimentConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

