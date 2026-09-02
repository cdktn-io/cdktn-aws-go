package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsinspector/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsinspector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFilter_VulnerablePackagesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Architecture() TfFilter_ArchitecturePropertyList
	// Experimental.
	ArchitectureInput() interface{}
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
	Epoch() TfFilter_EpochPropertyList
	// Experimental.
	EpochInput() interface{}
	// Experimental.
	FilePath() TfFilter_FilePathPropertyList
	// Experimental.
	FilePathInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() TfFilter_NamePropertyList
	// Experimental.
	NameInput() interface{}
	// Experimental.
	Release() TfFilter_ReleasePropertyList
	// Experimental.
	ReleaseInput() interface{}
	// Experimental.
	SourceLambdaLayerArn() TfFilter_SourceLambdaLayerArnPropertyList
	// Experimental.
	SourceLambdaLayerArnInput() interface{}
	// Experimental.
	SourceLayerHash() TfFilter_SourceLayerHashPropertyList
	// Experimental.
	SourceLayerHashInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Version() TfFilter_VersionPropertyList
	// Experimental.
	VersionInput() interface{}
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
	PutArchitecture(value interface{})
	// Experimental.
	PutEpoch(value interface{})
	// Experimental.
	PutFilePath(value interface{})
	// Experimental.
	PutName(value interface{})
	// Experimental.
	PutRelease(value interface{})
	// Experimental.
	PutSourceLambdaLayerArn(value interface{})
	// Experimental.
	PutSourceLayerHash(value interface{})
	// Experimental.
	PutVersion(value interface{})
	// Experimental.
	ResetArchitecture()
	// Experimental.
	ResetEpoch()
	// Experimental.
	ResetFilePath()
	// Experimental.
	ResetName()
	// Experimental.
	ResetRelease()
	// Experimental.
	ResetSourceLambdaLayerArn()
	// Experimental.
	ResetSourceLayerHash()
	// Experimental.
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFilter_VulnerablePackagesPropertyOutputReference
type jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) Architecture() TfFilter_ArchitecturePropertyList {
	var returns TfFilter_ArchitecturePropertyList
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ArchitectureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) Epoch() TfFilter_EpochPropertyList {
	var returns TfFilter_EpochPropertyList
	_jsii_.Get(
		j,
		"epoch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) EpochInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"epochInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) FilePath() TfFilter_FilePathPropertyList {
	var returns TfFilter_FilePathPropertyList
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) FilePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) Name() TfFilter_NamePropertyList {
	var returns TfFilter_NamePropertyList
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) NameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) Release() TfFilter_ReleasePropertyList {
	var returns TfFilter_ReleasePropertyList
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ReleaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"releaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) SourceLambdaLayerArn() TfFilter_SourceLambdaLayerArnPropertyList {
	var returns TfFilter_SourceLambdaLayerArnPropertyList
	_jsii_.Get(
		j,
		"sourceLambdaLayerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) SourceLambdaLayerArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceLambdaLayerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) SourceLayerHash() TfFilter_SourceLayerHashPropertyList {
	var returns TfFilter_SourceLayerHashPropertyList
	_jsii_.Get(
		j,
		"sourceLayerHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) SourceLayerHashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceLayerHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) Version() TfFilter_VersionPropertyList {
	var returns TfFilter_VersionPropertyList
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) VersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFilter_VulnerablePackagesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfFilter_VulnerablePackagesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFilter_VulnerablePackagesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-inspector.TfFilter.VulnerablePackagesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFilter_VulnerablePackagesPropertyOutputReference_Override(t TfFilter_VulnerablePackagesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-inspector.TfFilter.VulnerablePackagesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) PutArchitecture(value interface{}) {
	if err := t.validatePutArchitectureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putArchitecture",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) PutEpoch(value interface{}) {
	if err := t.validatePutEpochParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEpoch",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) PutFilePath(value interface{}) {
	if err := t.validatePutFilePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFilePath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) PutName(value interface{}) {
	if err := t.validatePutNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) PutRelease(value interface{}) {
	if err := t.validatePutReleaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRelease",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) PutSourceLambdaLayerArn(value interface{}) {
	if err := t.validatePutSourceLambdaLayerArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourceLambdaLayerArn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) PutSourceLayerHash(value interface{}) {
	if err := t.validatePutSourceLayerHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourceLayerHash",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) PutVersion(value interface{}) {
	if err := t.validatePutVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVersion",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ResetArchitecture() {
	_jsii_.InvokeVoid(
		t,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ResetEpoch() {
	_jsii_.InvokeVoid(
		t,
		"resetEpoch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ResetFilePath() {
	_jsii_.InvokeVoid(
		t,
		"resetFilePath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ResetRelease() {
	_jsii_.InvokeVoid(
		t,
		"resetRelease",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ResetSourceLambdaLayerArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceLambdaLayerArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ResetSourceLayerHash() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceLayerHash",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFilter_VulnerablePackagesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

