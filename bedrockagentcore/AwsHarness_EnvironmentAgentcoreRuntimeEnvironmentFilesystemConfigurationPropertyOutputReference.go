package bedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference interface {
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
	EfsAccessPoint() AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationEfsAccessPointPropertyList
	// Experimental.
	EfsAccessPointInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	S3FilesAccessPoint() AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationS3FilesAccessPointPropertyList
	// Experimental.
	S3FilesAccessPointInput() interface{}
	// Experimental.
	SessionStorage() AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationSessionStoragePropertyList
	// Experimental.
	SessionStorageInput() interface{}
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
	PutEfsAccessPoint(value interface{})
	// Experimental.
	PutS3FilesAccessPoint(value interface{})
	// Experimental.
	PutSessionStorage(value interface{})
	// Experimental.
	ResetEfsAccessPoint()
	// Experimental.
	ResetS3FilesAccessPoint()
	// Experimental.
	ResetSessionStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference
type jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) EfsAccessPoint() AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationEfsAccessPointPropertyList {
	var returns AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationEfsAccessPointPropertyList
	_jsii_.Get(
		j,
		"efsAccessPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) EfsAccessPointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efsAccessPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) S3FilesAccessPoint() AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationS3FilesAccessPointPropertyList {
	var returns AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationS3FilesAccessPointPropertyList
	_jsii_.Get(
		j,
		"s3FilesAccessPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) S3FilesAccessPointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3FilesAccessPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) SessionStorage() AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationSessionStoragePropertyList {
	var returns AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationSessionStoragePropertyList
	_jsii_.Get(
		j,
		"sessionStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) SessionStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsHarness.EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference_Override(a AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsHarness.EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) PutEfsAccessPoint(value interface{}) {
	if err := a.validatePutEfsAccessPointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEfsAccessPoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) PutS3FilesAccessPoint(value interface{}) {
	if err := a.validatePutS3FilesAccessPointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3FilesAccessPoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) PutSessionStorage(value interface{}) {
	if err := a.validatePutSessionStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSessionStorage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) ResetEfsAccessPoint() {
	_jsii_.InvokeVoid(
		a,
		"resetEfsAccessPoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) ResetS3FilesAccessPoint() {
	_jsii_.InvokeVoid(
		a,
		"resetS3FilesAccessPoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) ResetSessionStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

