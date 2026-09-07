package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference interface {
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
	InitialActiveLearningModelArn() *string
	// Experimental.
	SetInitialActiveLearningModelArn(val *string)
	// Experimental.
	InitialActiveLearningModelArnInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LabelingJobAlgorithmSpecificationArn() *string
	// Experimental.
	SetLabelingJobAlgorithmSpecificationArn(val *string)
	// Experimental.
	LabelingJobAlgorithmSpecificationArnInput() *string
	// Experimental.
	LabelingJobResourceConfig() AwsLabelingJob_LabelingJobResourceConfigPropertyList
	// Experimental.
	LabelingJobResourceConfigInput() interface{}
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
	PutLabelingJobResourceConfig(value interface{})
	// Experimental.
	ResetInitialActiveLearningModelArn()
	// Experimental.
	ResetLabelingJobResourceConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference
type jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) InitialActiveLearningModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialActiveLearningModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) InitialActiveLearningModelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialActiveLearningModelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) LabelingJobAlgorithmSpecificationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobAlgorithmSpecificationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) LabelingJobAlgorithmSpecificationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelingJobAlgorithmSpecificationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) LabelingJobResourceConfig() AwsLabelingJob_LabelingJobResourceConfigPropertyList {
	var returns AwsLabelingJob_LabelingJobResourceConfigPropertyList
	_jsii_.Get(
		j,
		"labelingJobResourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) LabelingJobResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelingJobResourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsLabelingJob.LabelingJobAlgorithmsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference_Override(a AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsLabelingJob.LabelingJobAlgorithmsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference)SetInitialActiveLearningModelArn(val *string) {
	if err := j.validateSetInitialActiveLearningModelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialActiveLearningModelArn",
		val,
	)
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference)SetLabelingJobAlgorithmSpecificationArn(val *string) {
	if err := j.validateSetLabelingJobAlgorithmSpecificationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelingJobAlgorithmSpecificationArn",
		val,
	)
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) PutLabelingJobResourceConfig(value interface{}) {
	if err := a.validatePutLabelingJobResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLabelingJobResourceConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) ResetInitialActiveLearningModelArn() {
	_jsii_.InvokeVoid(
		a,
		"resetInitialActiveLearningModelArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) ResetLabelingJobResourceConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLabelingJobResourceConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLabelingJob_LabelingJobAlgorithmsConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

