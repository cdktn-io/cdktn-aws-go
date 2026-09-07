package inspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/inspector/jsii"

	"github.com/cdktn-io/cdktn-aws-go/inspector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFilter_VulnerablePackagesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Architecture() AwsFilter_ArchitecturePropertyList
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
	Epoch() AwsFilter_EpochPropertyList
	// Experimental.
	EpochInput() interface{}
	// Experimental.
	FilePath() AwsFilter_FilePathPropertyList
	// Experimental.
	FilePathInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() AwsFilter_NamePropertyList
	// Experimental.
	NameInput() interface{}
	// Experimental.
	Release() AwsFilter_ReleasePropertyList
	// Experimental.
	ReleaseInput() interface{}
	// Experimental.
	SourceLambdaLayerArn() AwsFilter_SourceLambdaLayerArnPropertyList
	// Experimental.
	SourceLambdaLayerArnInput() interface{}
	// Experimental.
	SourceLayerHash() AwsFilter_SourceLayerHashPropertyList
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
	Version() AwsFilter_VersionPropertyList
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

// The jsii proxy struct for AwsFilter_VulnerablePackagesPropertyOutputReference
type jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) Architecture() AwsFilter_ArchitecturePropertyList {
	var returns AwsFilter_ArchitecturePropertyList
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ArchitectureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) Epoch() AwsFilter_EpochPropertyList {
	var returns AwsFilter_EpochPropertyList
	_jsii_.Get(
		j,
		"epoch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) EpochInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"epochInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) FilePath() AwsFilter_FilePathPropertyList {
	var returns AwsFilter_FilePathPropertyList
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) FilePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) Name() AwsFilter_NamePropertyList {
	var returns AwsFilter_NamePropertyList
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) NameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) Release() AwsFilter_ReleasePropertyList {
	var returns AwsFilter_ReleasePropertyList
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ReleaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"releaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) SourceLambdaLayerArn() AwsFilter_SourceLambdaLayerArnPropertyList {
	var returns AwsFilter_SourceLambdaLayerArnPropertyList
	_jsii_.Get(
		j,
		"sourceLambdaLayerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) SourceLambdaLayerArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceLambdaLayerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) SourceLayerHash() AwsFilter_SourceLayerHashPropertyList {
	var returns AwsFilter_SourceLayerHashPropertyList
	_jsii_.Get(
		j,
		"sourceLayerHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) SourceLayerHashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceLayerHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) Version() AwsFilter_VersionPropertyList {
	var returns AwsFilter_VersionPropertyList
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) VersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFilter_VulnerablePackagesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsFilter_VulnerablePackagesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFilter_VulnerablePackagesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-inspector.AwsFilter.VulnerablePackagesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFilter_VulnerablePackagesPropertyOutputReference_Override(a AwsFilter_VulnerablePackagesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-inspector.AwsFilter.VulnerablePackagesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) PutArchitecture(value interface{}) {
	if err := a.validatePutArchitectureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArchitecture",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) PutEpoch(value interface{}) {
	if err := a.validatePutEpochParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEpoch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) PutFilePath(value interface{}) {
	if err := a.validatePutFilePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilePath",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) PutName(value interface{}) {
	if err := a.validatePutNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) PutRelease(value interface{}) {
	if err := a.validatePutReleaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelease",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) PutSourceLambdaLayerArn(value interface{}) {
	if err := a.validatePutSourceLambdaLayerArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceLambdaLayerArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) PutSourceLayerHash(value interface{}) {
	if err := a.validatePutSourceLayerHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceLayerHash",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) PutVersion(value interface{}) {
	if err := a.validatePutVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVersion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ResetArchitecture() {
	_jsii_.InvokeVoid(
		a,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ResetEpoch() {
	_jsii_.InvokeVoid(
		a,
		"resetEpoch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ResetFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ResetRelease() {
	_jsii_.InvokeVoid(
		a,
		"resetRelease",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ResetSourceLambdaLayerArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceLambdaLayerArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ResetSourceLayerHash() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceLayerHash",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFilter_VulnerablePackagesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

