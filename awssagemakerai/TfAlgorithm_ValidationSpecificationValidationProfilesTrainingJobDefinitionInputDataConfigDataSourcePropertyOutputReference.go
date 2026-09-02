package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference interface {
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
	FileSystemDataSource() TfAlgorithm_FileSystemDataSourcePropertyList
	// Experimental.
	FileSystemDataSourceInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	S3DataSource() TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyList
	// Experimental.
	S3DataSourceInput() interface{}
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
	PutFileSystemDataSource(value interface{})
	// Experimental.
	PutS3DataSource(value interface{})
	// Experimental.
	ResetFileSystemDataSource()
	// Experimental.
	ResetS3DataSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference
type jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) FileSystemDataSource() TfAlgorithm_FileSystemDataSourcePropertyList {
	var returns TfAlgorithm_FileSystemDataSourcePropertyList
	_jsii_.Get(
		j,
		"fileSystemDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) FileSystemDataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) S3DataSource() TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyList {
	var returns TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyList
	_jsii_.Get(
		j,
		"s3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) S3DataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAlgorithm.ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference_Override(t TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAlgorithm.ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) PutFileSystemDataSource(value interface{}) {
	if err := t.validatePutFileSystemDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFileSystemDataSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) PutS3DataSource(value interface{}) {
	if err := t.validatePutS3DataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3DataSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ResetFileSystemDataSource() {
	_jsii_.InvokeVoid(
		t,
		"resetFileSystemDataSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ResetS3DataSource() {
	_jsii_.InvokeVoid(
		t,
		"resetS3DataSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

