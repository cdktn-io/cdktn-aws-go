package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference interface {
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
	InstanceType() *string
	// Experimental.
	SetInstanceType(val *string)
	// Experimental.
	InstanceTypeInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LocalPath() *string
	// Experimental.
	SetLocalPath(val *string)
	// Experimental.
	LocalPathInput() *string
	// Experimental.
	RuleConfigurationName() *string
	// Experimental.
	SetRuleConfigurationName(val *string)
	// Experimental.
	RuleConfigurationNameInput() *string
	// Experimental.
	RuleEvaluatorImage() *string
	// Experimental.
	SetRuleEvaluatorImage(val *string)
	// Experimental.
	RuleEvaluatorImageInput() *string
	// Experimental.
	RuleParameters() *map[string]*string
	// Experimental.
	SetRuleParameters(val *map[string]*string)
	// Experimental.
	RuleParametersInput() *map[string]*string
	// Experimental.
	S3OutputPath() *string
	// Experimental.
	SetS3OutputPath(val *string)
	// Experimental.
	S3OutputPathInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VolumeSizeInGb() *float64
	// Experimental.
	SetVolumeSizeInGb(val *float64)
	// Experimental.
	VolumeSizeInGbInput() *float64
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
	ResetInstanceType()
	// Experimental.
	ResetLocalPath()
	// Experimental.
	ResetRuleParameters()
	// Experimental.
	ResetS3OutputPath()
	// Experimental.
	ResetVolumeSizeInGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference
type jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) LocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) RuleConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) RuleConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) RuleEvaluatorImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleEvaluatorImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) RuleEvaluatorImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleEvaluatorImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) RuleParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"ruleParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) RuleParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"ruleParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) S3OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3OutputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) S3OutputPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3OutputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) VolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) VolumeSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGbInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsTrainingJob.ProfilerRuleConfigurationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference_Override(a AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsTrainingJob.ProfilerRuleConfigurationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetLocalPath(val *string) {
	if err := j.validateSetLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPath",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetRuleConfigurationName(val *string) {
	if err := j.validateSetRuleConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleConfigurationName",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetRuleEvaluatorImage(val *string) {
	if err := j.validateSetRuleEvaluatorImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleEvaluatorImage",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetRuleParameters(val *map[string]*string) {
	if err := j.validateSetRuleParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleParameters",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetS3OutputPath(val *string) {
	if err := j.validateSetS3OutputPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3OutputPath",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference)SetVolumeSizeInGb(val *float64) {
	if err := j.validateSetVolumeSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSizeInGb",
		val,
	)
}

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) ResetLocalPath() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) ResetRuleParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRuleParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) ResetS3OutputPath() {
	_jsii_.InvokeVoid(
		a,
		"resetS3OutputPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) ResetVolumeSizeInGb() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumeSizeInGb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTrainingJob_ProfilerRuleConfigurationsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

