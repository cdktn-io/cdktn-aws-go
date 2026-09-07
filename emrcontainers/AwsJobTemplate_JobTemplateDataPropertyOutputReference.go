package emrcontainers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/emrcontainers/jsii"

	"github.com/cdktn-io/cdktn-aws-go/emrcontainers/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsJobTemplate_JobTemplateDataPropertyOutputReference interface {
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
	// Experimental.
	ConfigurationOverrides() AwsJobTemplate_ConfigurationOverridesPropertyOutputReference
	// Experimental.
	ConfigurationOverridesInput() *AwsJobTemplate_ConfigurationOverridesProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	ExecutionRoleArn() *string
	// Experimental.
	SetExecutionRoleArn(val *string)
	// Experimental.
	ExecutionRoleArnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsJobTemplate_JobTemplateDataProperty
	// Experimental.
	SetInternalValue(val *AwsJobTemplate_JobTemplateDataProperty)
	// Experimental.
	JobDriver() AwsJobTemplate_JobDriverPropertyOutputReference
	// Experimental.
	JobDriverInput() *AwsJobTemplate_JobDriverProperty
	// Experimental.
	JobTags() *map[string]*string
	// Experimental.
	SetJobTags(val *map[string]*string)
	// Experimental.
	JobTagsInput() *map[string]*string
	// Experimental.
	ReleaseLabel() *string
	// Experimental.
	SetReleaseLabel(val *string)
	// Experimental.
	ReleaseLabelInput() *string
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
	PutConfigurationOverrides(value *AwsJobTemplate_ConfigurationOverridesProperty)
	// Experimental.
	PutJobDriver(value *AwsJobTemplate_JobDriverProperty)
	// Experimental.
	ResetConfigurationOverrides()
	// Experimental.
	ResetJobTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsJobTemplate_JobTemplateDataPropertyOutputReference
type jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ConfigurationOverrides() AwsJobTemplate_ConfigurationOverridesPropertyOutputReference {
	var returns AwsJobTemplate_ConfigurationOverridesPropertyOutputReference
	_jsii_.Get(
		j,
		"configurationOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ConfigurationOverridesInput() *AwsJobTemplate_ConfigurationOverridesProperty {
	var returns *AwsJobTemplate_ConfigurationOverridesProperty
	_jsii_.Get(
		j,
		"configurationOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) InternalValue() *AwsJobTemplate_JobTemplateDataProperty {
	var returns *AwsJobTemplate_JobTemplateDataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) JobDriver() AwsJobTemplate_JobDriverPropertyOutputReference {
	var returns AwsJobTemplate_JobDriverPropertyOutputReference
	_jsii_.Get(
		j,
		"jobDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) JobDriverInput() *AwsJobTemplate_JobDriverProperty {
	var returns *AwsJobTemplate_JobDriverProperty
	_jsii_.Get(
		j,
		"jobDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) JobTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"jobTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) JobTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"jobTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsJobTemplate_JobTemplateDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsJobTemplate_JobTemplateDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsJobTemplate_JobTemplateDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr-containers.AwsJobTemplate.JobTemplateDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsJobTemplate_JobTemplateDataPropertyOutputReference_Override(a AwsJobTemplate_JobTemplateDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-containers.AwsJobTemplate.JobTemplateDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference)SetInternalValue(val *AwsJobTemplate_JobTemplateDataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference)SetJobTags(val *map[string]*string) {
	if err := j.validateSetJobTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobTags",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) PutConfigurationOverrides(value *AwsJobTemplate_ConfigurationOverridesProperty) {
	if err := a.validatePutConfigurationOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfigurationOverrides",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) PutJobDriver(value *AwsJobTemplate_JobDriverProperty) {
	if err := a.validatePutJobDriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJobDriver",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ResetConfigurationOverrides() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigurationOverrides",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ResetJobTags() {
	_jsii_.InvokeVoid(
		a,
		"resetJobTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsJobTemplate_JobTemplateDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

