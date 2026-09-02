package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AttributeNames() *[]*string
	// Experimental.
	SetAttributeNames(val *[]*string)
	// Experimental.
	AttributeNamesInput() *[]*string
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
	HubAccessConfig() TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourceHubAccessConfigPropertyList
	// Experimental.
	HubAccessConfigInput() interface{}
	// Experimental.
	InstanceGroupNames() *[]*string
	// Experimental.
	SetInstanceGroupNames(val *[]*string)
	// Experimental.
	InstanceGroupNamesInput() *[]*string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ModelAccessConfig() TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourceModelAccessConfigPropertyList
	// Experimental.
	ModelAccessConfigInput() interface{}
	// Experimental.
	S3DataDistributionType() *string
	// Experimental.
	SetS3DataDistributionType(val *string)
	// Experimental.
	S3DataDistributionTypeInput() *string
	// Experimental.
	S3DataType() *string
	// Experimental.
	SetS3DataType(val *string)
	// Experimental.
	S3DataTypeInput() *string
	// Experimental.
	S3Uri() *string
	// Experimental.
	SetS3Uri(val *string)
	// Experimental.
	S3UriInput() *string
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
	PutHubAccessConfig(value interface{})
	// Experimental.
	PutModelAccessConfig(value interface{})
	// Experimental.
	ResetAttributeNames()
	// Experimental.
	ResetHubAccessConfig()
	// Experimental.
	ResetInstanceGroupNames()
	// Experimental.
	ResetModelAccessConfig()
	// Experimental.
	ResetS3DataDistributionType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference
type jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) AttributeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) AttributeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) HubAccessConfig() TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourceHubAccessConfigPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourceHubAccessConfigPropertyList
	_jsii_.Get(
		j,
		"hubAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) HubAccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hubAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) InstanceGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) InstanceGroupNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGroupNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) ModelAccessConfig() TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourceModelAccessConfigPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourceModelAccessConfigPropertyList
	_jsii_.Get(
		j,
		"modelAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) ModelAccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) S3DataDistributionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) S3DataDistributionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) S3DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) S3DataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) S3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) S3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference_Override(t TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference)SetAttributeNames(val *[]*string) {
	if err := j.validateSetAttributeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeNames",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference)SetInstanceGroupNames(val *[]*string) {
	if err := j.validateSetInstanceGroupNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGroupNames",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference)SetS3DataDistributionType(val *string) {
	if err := j.validateSetS3DataDistributionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3DataDistributionType",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference)SetS3DataType(val *string) {
	if err := j.validateSetS3DataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3DataType",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference)SetS3Uri(val *string) {
	if err := j.validateSetS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Uri",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) PutHubAccessConfig(value interface{}) {
	if err := t.validatePutHubAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHubAccessConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) PutModelAccessConfig(value interface{}) {
	if err := t.validatePutModelAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putModelAccessConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) ResetAttributeNames() {
	_jsii_.InvokeVoid(
		t,
		"resetAttributeNames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) ResetHubAccessConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetHubAccessConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) ResetInstanceGroupNames() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceGroupNames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) ResetModelAccessConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetModelAccessConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) ResetS3DataDistributionType() {
	_jsii_.InvokeVoid(
		t,
		"resetS3DataDistributionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

