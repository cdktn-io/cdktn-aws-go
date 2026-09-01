package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference interface {
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
	FileSystemDataSource() AwsSagemakerAlgorithm_FileSystemDataSourcePropertyList
	// Experimental.
	FileSystemDataSourceInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	S3DataSource() AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyList
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

// The jsii proxy struct for AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference
type jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) FileSystemDataSource() AwsSagemakerAlgorithm_FileSystemDataSourcePropertyList {
	var returns AwsSagemakerAlgorithm_FileSystemDataSourcePropertyList
	_jsii_.Get(
		j,
		"fileSystemDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) FileSystemDataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) S3DataSource() AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyList {
	var returns AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourceS3DataSourcePropertyList
	_jsii_.Get(
		j,
		"s3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) S3DataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerAlgorithm.ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference_Override(a AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerAlgorithm.ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) PutFileSystemDataSource(value interface{}) {
	if err := a.validatePutFileSystemDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFileSystemDataSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) PutS3DataSource(value interface{}) {
	if err := a.validatePutS3DataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3DataSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ResetFileSystemDataSource() {
	_jsii_.InvokeVoid(
		a,
		"resetFileSystemDataSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ResetS3DataSource() {
	_jsii_.InvokeVoid(
		a,
		"resetS3DataSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_ValidationSpecificationValidationProfilesTrainingJobDefinitionInputDataConfigDataSourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

