package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference interface {
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
	CompressionType() *string
	// Experimental.
	SetCompressionType(val *string)
	// Experimental.
	CompressionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Etag() *string
	// Experimental.
	SetEtag(val *string)
	// Experimental.
	EtagInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	HubAccessConfig() TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourceHubAccessConfigPropertyList
	// Experimental.
	HubAccessConfigInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ManifestEtag() *string
	// Experimental.
	SetManifestEtag(val *string)
	// Experimental.
	ManifestEtagInput() *string
	// Experimental.
	ManifestS3Uri() *string
	// Experimental.
	SetManifestS3Uri(val *string)
	// Experimental.
	ManifestS3UriInput() *string
	// Experimental.
	ModelAccessConfig() TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourceModelAccessConfigPropertyList
	// Experimental.
	ModelAccessConfigInput() interface{}
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
	ResetEtag()
	// Experimental.
	ResetHubAccessConfig()
	// Experimental.
	ResetManifestEtag()
	// Experimental.
	ResetManifestS3Uri()
	// Experimental.
	ResetModelAccessConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference
type jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) HubAccessConfig() TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourceHubAccessConfigPropertyList {
	var returns TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourceHubAccessConfigPropertyList
	_jsii_.Get(
		j,
		"hubAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) HubAccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hubAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ManifestEtag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestEtag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ManifestEtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestEtagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ManifestS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ManifestS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ModelAccessConfig() TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourceModelAccessConfigPropertyList {
	var returns TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourceModelAccessConfigPropertyList
	_jsii_.Get(
		j,
		"modelAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ModelAccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) S3DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) S3DataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) S3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) S3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAlgorithm.InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference_Override(t TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAlgorithm.InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference)SetManifestEtag(val *string) {
	if err := j.validateSetManifestEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestEtag",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference)SetManifestS3Uri(val *string) {
	if err := j.validateSetManifestS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestS3Uri",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference)SetS3DataType(val *string) {
	if err := j.validateSetS3DataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3DataType",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference)SetS3Uri(val *string) {
	if err := j.validateSetS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Uri",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) PutHubAccessConfig(value interface{}) {
	if err := t.validatePutHubAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHubAccessConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) PutModelAccessConfig(value interface{}) {
	if err := t.validatePutModelAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putModelAccessConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ResetEtag() {
	_jsii_.InvokeVoid(
		t,
		"resetEtag",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ResetHubAccessConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetHubAccessConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ResetManifestEtag() {
	_jsii_.InvokeVoid(
		t,
		"resetManifestEtag",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ResetManifestS3Uri() {
	_jsii_.InvokeVoid(
		t,
		"resetManifestS3Uri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ResetModelAccessConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetModelAccessConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAlgorithm_InferenceSpecificationContainersModelDataSourceS3DataSourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

