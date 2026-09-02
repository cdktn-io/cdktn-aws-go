package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference interface {
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
	ContainerArguments() *[]*string
	// Experimental.
	SetContainerArguments(val *[]*string)
	// Experimental.
	ContainerArgumentsInput() *[]*string
	// Experimental.
	ContainerEntrypoint() *[]*string
	// Experimental.
	SetContainerEntrypoint(val *[]*string)
	// Experimental.
	ContainerEntrypointInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	ImageUri() *string
	// Experimental.
	SetImageUri(val *string)
	// Experimental.
	ImageUriInput() *string
	// Experimental.
	InternalValue() *TfMonitoringSchedule_MonitoringAppSpecificationProperty
	// Experimental.
	SetInternalValue(val *TfMonitoringSchedule_MonitoringAppSpecificationProperty)
	// Experimental.
	PostAnalyticsProcessorSourceUri() *string
	// Experimental.
	SetPostAnalyticsProcessorSourceUri(val *string)
	// Experimental.
	PostAnalyticsProcessorSourceUriInput() *string
	// Experimental.
	RecordPreprocessorSourceUri() *string
	// Experimental.
	SetRecordPreprocessorSourceUri(val *string)
	// Experimental.
	RecordPreprocessorSourceUriInput() *string
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
	ResetContainerArguments()
	// Experimental.
	ResetContainerEntrypoint()
	// Experimental.
	ResetPostAnalyticsProcessorSourceUri()
	// Experimental.
	ResetRecordPreprocessorSourceUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference
type jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ContainerArguments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ContainerArgumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ContainerEntrypoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ContainerEntrypointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) InternalValue() *TfMonitoringSchedule_MonitoringAppSpecificationProperty {
	var returns *TfMonitoringSchedule_MonitoringAppSpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) PostAnalyticsProcessorSourceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAnalyticsProcessorSourceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) PostAnalyticsProcessorSourceUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAnalyticsProcessorSourceUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) RecordPreprocessorSourceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordPreprocessorSourceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) RecordPreprocessorSourceUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordPreprocessorSourceUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.MonitoringAppSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference_Override(t TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.MonitoringAppSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference)SetContainerArguments(val *[]*string) {
	if err := j.validateSetContainerArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerArguments",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference)SetContainerEntrypoint(val *[]*string) {
	if err := j.validateSetContainerEntrypointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerEntrypoint",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference)SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference)SetInternalValue(val *TfMonitoringSchedule_MonitoringAppSpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference)SetPostAnalyticsProcessorSourceUri(val *string) {
	if err := j.validateSetPostAnalyticsProcessorSourceUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postAnalyticsProcessorSourceUri",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference)SetRecordPreprocessorSourceUri(val *string) {
	if err := j.validateSetRecordPreprocessorSourceUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordPreprocessorSourceUri",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ResetContainerArguments() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerArguments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ResetContainerEntrypoint() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerEntrypoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ResetPostAnalyticsProcessorSourceUri() {
	_jsii_.InvokeVoid(
		t,
		"resetPostAnalyticsProcessorSourceUri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ResetRecordPreprocessorSourceUri() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordPreprocessorSourceUri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

