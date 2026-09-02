package awsemrcontainers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemrcontainers/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemrcontainers/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfJobTemplate_JobTemplateDataPropertyOutputReference interface {
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
	ConfigurationOverrides() TfJobTemplate_ConfigurationOverridesPropertyOutputReference
	// Experimental.
	ConfigurationOverridesInput() *TfJobTemplate_ConfigurationOverridesProperty
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
	InternalValue() *TfJobTemplate_JobTemplateDataProperty
	// Experimental.
	SetInternalValue(val *TfJobTemplate_JobTemplateDataProperty)
	// Experimental.
	JobDriver() TfJobTemplate_JobDriverPropertyOutputReference
	// Experimental.
	JobDriverInput() *TfJobTemplate_JobDriverProperty
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
	PutConfigurationOverrides(value *TfJobTemplate_ConfigurationOverridesProperty)
	// Experimental.
	PutJobDriver(value *TfJobTemplate_JobDriverProperty)
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

// The jsii proxy struct for TfJobTemplate_JobTemplateDataPropertyOutputReference
type jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ConfigurationOverrides() TfJobTemplate_ConfigurationOverridesPropertyOutputReference {
	var returns TfJobTemplate_ConfigurationOverridesPropertyOutputReference
	_jsii_.Get(
		j,
		"configurationOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ConfigurationOverridesInput() *TfJobTemplate_ConfigurationOverridesProperty {
	var returns *TfJobTemplate_ConfigurationOverridesProperty
	_jsii_.Get(
		j,
		"configurationOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) InternalValue() *TfJobTemplate_JobTemplateDataProperty {
	var returns *TfJobTemplate_JobTemplateDataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) JobDriver() TfJobTemplate_JobDriverPropertyOutputReference {
	var returns TfJobTemplate_JobDriverPropertyOutputReference
	_jsii_.Get(
		j,
		"jobDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) JobDriverInput() *TfJobTemplate_JobDriverProperty {
	var returns *TfJobTemplate_JobDriverProperty
	_jsii_.Get(
		j,
		"jobDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) JobTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"jobTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) JobTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"jobTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfJobTemplate_JobTemplateDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfJobTemplate_JobTemplateDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfJobTemplate_JobTemplateDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr-containers.TfJobTemplate.JobTemplateDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfJobTemplate_JobTemplateDataPropertyOutputReference_Override(t TfJobTemplate_JobTemplateDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-containers.TfJobTemplate.JobTemplateDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference)SetInternalValue(val *TfJobTemplate_JobTemplateDataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference)SetJobTags(val *map[string]*string) {
	if err := j.validateSetJobTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobTags",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) PutConfigurationOverrides(value *TfJobTemplate_ConfigurationOverridesProperty) {
	if err := t.validatePutConfigurationOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConfigurationOverrides",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) PutJobDriver(value *TfJobTemplate_JobDriverProperty) {
	if err := t.validatePutJobDriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJobDriver",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ResetConfigurationOverrides() {
	_jsii_.InvokeVoid(
		t,
		"resetConfigurationOverrides",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ResetJobTags() {
	_jsii_.InvokeVoid(
		t,
		"resetJobTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfJobTemplate_JobTemplateDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

