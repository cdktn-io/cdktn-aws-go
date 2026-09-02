package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAlgorithm_ContainersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdditionalS3DataSource() TfAlgorithm_InferenceSpecificationContainersAdditionalS3DataSourcePropertyList
	// Experimental.
	AdditionalS3DataSourceInput() interface{}
	// Experimental.
	BaseModel() TfAlgorithm_BaseModelPropertyList
	// Experimental.
	BaseModelInput() interface{}
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
	ContainerHostname() *string
	// Experimental.
	SetContainerHostname(val *string)
	// Experimental.
	ContainerHostnameInput() *string
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
	Framework() *string
	// Experimental.
	SetFramework(val *string)
	// Experimental.
	FrameworkInput() *string
	// Experimental.
	FrameworkVersion() *string
	// Experimental.
	SetFrameworkVersion(val *string)
	// Experimental.
	FrameworkVersionInput() *string
	// Experimental.
	Image() *string
	// Experimental.
	SetImage(val *string)
	// Experimental.
	ImageDigest() *string
	// Experimental.
	SetImageDigest(val *string)
	// Experimental.
	ImageDigestInput() *string
	// Experimental.
	ImageInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IsCheckpoint() interface{}
	// Experimental.
	SetIsCheckpoint(val interface{})
	// Experimental.
	IsCheckpointInput() interface{}
	// Experimental.
	ModelDataEtag() *string
	// Experimental.
	SetModelDataEtag(val *string)
	// Experimental.
	ModelDataEtagInput() *string
	// Experimental.
	ModelDataSource() TfAlgorithm_ModelDataSourcePropertyList
	// Experimental.
	ModelDataSourceInput() interface{}
	// Experimental.
	ModelDataUrl() *string
	// Experimental.
	SetModelDataUrl(val *string)
	// Experimental.
	ModelDataUrlInput() *string
	// Experimental.
	ModelInput() TfAlgorithm_ModelInputPropertyList
	// Experimental.
	ModelInputInput() interface{}
	// Experimental.
	NearestModelName() *string
	// Experimental.
	SetNearestModelName(val *string)
	// Experimental.
	NearestModelNameInput() *string
	// Experimental.
	ProductId() *string
	// Experimental.
	SetProductId(val *string)
	// Experimental.
	ProductIdInput() *string
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
	PutAdditionalS3DataSource(value interface{})
	// Experimental.
	PutBaseModel(value interface{})
	// Experimental.
	PutModelDataSource(value interface{})
	// Experimental.
	PutModelInput(value interface{})
	// Experimental.
	ResetAdditionalS3DataSource()
	// Experimental.
	ResetBaseModel()
	// Experimental.
	ResetContainerHostname()
	// Experimental.
	ResetEnvironment()
	// Experimental.
	ResetFramework()
	// Experimental.
	ResetFrameworkVersion()
	// Experimental.
	ResetImage()
	// Experimental.
	ResetImageDigest()
	// Experimental.
	ResetIsCheckpoint()
	// Experimental.
	ResetModelDataEtag()
	// Experimental.
	ResetModelDataSource()
	// Experimental.
	ResetModelDataUrl()
	// Experimental.
	ResetModelInput()
	// Experimental.
	ResetNearestModelName()
	// Experimental.
	ResetProductId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAlgorithm_ContainersPropertyOutputReference
type jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) AdditionalS3DataSource() TfAlgorithm_InferenceSpecificationContainersAdditionalS3DataSourcePropertyList {
	var returns TfAlgorithm_InferenceSpecificationContainersAdditionalS3DataSourcePropertyList
	_jsii_.Get(
		j,
		"additionalS3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) AdditionalS3DataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalS3DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) BaseModel() TfAlgorithm_BaseModelPropertyList {
	var returns TfAlgorithm_BaseModelPropertyList
	_jsii_.Get(
		j,
		"baseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) BaseModelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ContainerHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ContainerHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) Framework() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framework",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) FrameworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) FrameworkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) FrameworkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameworkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ImageDigest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageDigest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ImageDigestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageDigestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) IsCheckpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCheckpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) IsCheckpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCheckpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ModelDataEtag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataEtag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ModelDataEtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataEtagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ModelDataSource() TfAlgorithm_ModelDataSourcePropertyList {
	var returns TfAlgorithm_ModelDataSourcePropertyList
	_jsii_.Get(
		j,
		"modelDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ModelDataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ModelDataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ModelDataUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ModelInput() TfAlgorithm_ModelInputPropertyList {
	var returns TfAlgorithm_ModelInputPropertyList
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ModelInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) NearestModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nearestModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) NearestModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nearestModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAlgorithm_ContainersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfAlgorithm_ContainersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAlgorithm_ContainersPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAlgorithm.ContainersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAlgorithm_ContainersPropertyOutputReference_Override(t TfAlgorithm_ContainersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAlgorithm.ContainersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetContainerHostname(val *string) {
	if err := j.validateSetContainerHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerHostname",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetFramework(val *string) {
	if err := j.validateSetFrameworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framework",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetFrameworkVersion(val *string) {
	if err := j.validateSetFrameworkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frameworkVersion",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetImageDigest(val *string) {
	if err := j.validateSetImageDigestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageDigest",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetIsCheckpoint(val interface{}) {
	if err := j.validateSetIsCheckpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCheckpoint",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetModelDataEtag(val *string) {
	if err := j.validateSetModelDataEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataEtag",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetModelDataUrl(val *string) {
	if err := j.validateSetModelDataUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataUrl",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetNearestModelName(val *string) {
	if err := j.validateSetNearestModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nearestModelName",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetProductId(val *string) {
	if err := j.validateSetProductIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) PutAdditionalS3DataSource(value interface{}) {
	if err := t.validatePutAdditionalS3DataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdditionalS3DataSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) PutBaseModel(value interface{}) {
	if err := t.validatePutBaseModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBaseModel",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) PutModelDataSource(value interface{}) {
	if err := t.validatePutModelDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putModelDataSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) PutModelInput(value interface{}) {
	if err := t.validatePutModelInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putModelInput",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetAdditionalS3DataSource() {
	_jsii_.InvokeVoid(
		t,
		"resetAdditionalS3DataSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetBaseModel() {
	_jsii_.InvokeVoid(
		t,
		"resetBaseModel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetContainerHostname() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerHostname",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetFramework() {
	_jsii_.InvokeVoid(
		t,
		"resetFramework",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetFrameworkVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetFrameworkVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		t,
		"resetImage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetImageDigest() {
	_jsii_.InvokeVoid(
		t,
		"resetImageDigest",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetIsCheckpoint() {
	_jsii_.InvokeVoid(
		t,
		"resetIsCheckpoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetModelDataEtag() {
	_jsii_.InvokeVoid(
		t,
		"resetModelDataEtag",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetModelDataSource() {
	_jsii_.InvokeVoid(
		t,
		"resetModelDataSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetModelDataUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetModelDataUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetModelInput() {
	_jsii_.InvokeVoid(
		t,
		"resetModelInput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetNearestModelName() {
	_jsii_.InvokeVoid(
		t,
		"resetNearestModelName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ResetProductId() {
	_jsii_.InvokeVoid(
		t,
		"resetProductId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAlgorithm_ContainersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

