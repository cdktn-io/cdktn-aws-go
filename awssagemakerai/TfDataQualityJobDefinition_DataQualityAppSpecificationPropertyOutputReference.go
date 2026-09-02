package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference interface {
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
	Environment() *map[string]*string
	// Experimental.
	SetEnvironment(val *map[string]*string)
	// Experimental.
	EnvironmentInput() *map[string]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	ImageUri() *string
	// Experimental.
	SetImageUri(val *string)
	// Experimental.
	ImageUriInput() *string
	// Experimental.
	InternalValue() *TfDataQualityJobDefinition_DataQualityAppSpecificationProperty
	// Experimental.
	SetInternalValue(val *TfDataQualityJobDefinition_DataQualityAppSpecificationProperty)
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
	ResetEnvironment()
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

// The jsii proxy struct for TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference
type jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) InternalValue() *TfDataQualityJobDefinition_DataQualityAppSpecificationProperty {
	var returns *TfDataQualityJobDefinition_DataQualityAppSpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) PostAnalyticsProcessorSourceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAnalyticsProcessorSourceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) PostAnalyticsProcessorSourceUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAnalyticsProcessorSourceUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) RecordPreprocessorSourceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordPreprocessorSourceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) RecordPreprocessorSourceUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordPreprocessorSourceUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDataQualityJobDefinition.DataQualityAppSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference_Override(t TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDataQualityJobDefinition.DataQualityAppSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference)SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference)SetInternalValue(val *TfDataQualityJobDefinition_DataQualityAppSpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference)SetPostAnalyticsProcessorSourceUri(val *string) {
	if err := j.validateSetPostAnalyticsProcessorSourceUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postAnalyticsProcessorSourceUri",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference)SetRecordPreprocessorSourceUri(val *string) {
	if err := j.validateSetRecordPreprocessorSourceUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordPreprocessorSourceUri",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) ResetPostAnalyticsProcessorSourceUri() {
	_jsii_.InvokeVoid(
		t,
		"resetPostAnalyticsProcessorSourceUri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) ResetRecordPreprocessorSourceUri() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordPreprocessorSourceUri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataQualityJobDefinition_DataQualityAppSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

