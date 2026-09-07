package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference interface {
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
	EnableMultipleJobs() interface{}
	// Experimental.
	SetEnableMultipleJobs(val interface{})
	// Experimental.
	EnableMultipleJobsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PlacementSpecifications() AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPlacementSpecificationsPropertyList
	// Experimental.
	PlacementSpecificationsInput() interface{}
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
	PutPlacementSpecifications(value interface{})
	// Experimental.
	ResetEnableMultipleJobs()
	// Experimental.
	ResetPlacementSpecifications()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference
type jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) EnableMultipleJobs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultipleJobs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) EnableMultipleJobsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultipleJobsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) PlacementSpecifications() AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPlacementSpecificationsPropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPlacementSpecificationsPropertyList
	_jsii_.Get(
		j,
		"placementSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) PlacementSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsHyperParameterTuningJob.TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference_Override(a AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsHyperParameterTuningJob.TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference)SetEnableMultipleJobs(val interface{}) {
	if err := j.validateSetEnableMultipleJobsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMultipleJobs",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) PutPlacementSpecifications(value interface{}) {
	if err := a.validatePutPlacementSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPlacementSpecifications",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) ResetEnableMultipleJobs() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableMultipleJobs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) ResetPlacementSpecifications() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementSpecifications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigInstancePlacementConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

