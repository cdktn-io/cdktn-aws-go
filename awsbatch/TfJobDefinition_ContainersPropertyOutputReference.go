package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbatch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbatch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfJobDefinition_ContainersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Args() *[]*string
	// Experimental.
	SetArgs(val *[]*string)
	// Experimental.
	ArgsInput() *[]*string
	// Experimental.
	Command() *[]*string
	// Experimental.
	SetCommand(val *[]*string)
	// Experimental.
	CommandInput() *[]*string
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
	Env() TfJobDefinition_EksPropertiesPodPropertiesContainersEnvPropertyList
	// Experimental.
	EnvInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Image() *string
	// Experimental.
	SetImage(val *string)
	// Experimental.
	ImageInput() *string
	// Experimental.
	ImagePullPolicy() *string
	// Experimental.
	SetImagePullPolicy(val *string)
	// Experimental.
	ImagePullPolicyInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	Resources() TfJobDefinition_EksPropertiesPodPropertiesContainersResourcesPropertyOutputReference
	// Experimental.
	ResourcesInput() *TfJobDefinition_EksPropertiesPodPropertiesContainersResourcesProperty
	// Experimental.
	SecurityContext() TfJobDefinition_EksPropertiesPodPropertiesContainersSecurityContextPropertyOutputReference
	// Experimental.
	SecurityContextInput() *TfJobDefinition_EksPropertiesPodPropertiesContainersSecurityContextProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VolumeMounts() TfJobDefinition_EksPropertiesPodPropertiesContainersVolumeMountsPropertyList
	// Experimental.
	VolumeMountsInput() interface{}
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
	PutEnv(value interface{})
	// Experimental.
	PutResources(value *TfJobDefinition_EksPropertiesPodPropertiesContainersResourcesProperty)
	// Experimental.
	PutSecurityContext(value *TfJobDefinition_EksPropertiesPodPropertiesContainersSecurityContextProperty)
	// Experimental.
	PutVolumeMounts(value interface{})
	// Experimental.
	ResetArgs()
	// Experimental.
	ResetCommand()
	// Experimental.
	ResetEnv()
	// Experimental.
	ResetImagePullPolicy()
	// Experimental.
	ResetName()
	// Experimental.
	ResetResources()
	// Experimental.
	ResetSecurityContext()
	// Experimental.
	ResetVolumeMounts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfJobDefinition_ContainersPropertyOutputReference
type jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) Env() TfJobDefinition_EksPropertiesPodPropertiesContainersEnvPropertyList {
	var returns TfJobDefinition_EksPropertiesPodPropertiesContainersEnvPropertyList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) Resources() TfJobDefinition_EksPropertiesPodPropertiesContainersResourcesPropertyOutputReference {
	var returns TfJobDefinition_EksPropertiesPodPropertiesContainersResourcesPropertyOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ResourcesInput() *TfJobDefinition_EksPropertiesPodPropertiesContainersResourcesProperty {
	var returns *TfJobDefinition_EksPropertiesPodPropertiesContainersResourcesProperty
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) SecurityContext() TfJobDefinition_EksPropertiesPodPropertiesContainersSecurityContextPropertyOutputReference {
	var returns TfJobDefinition_EksPropertiesPodPropertiesContainersSecurityContextPropertyOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) SecurityContextInput() *TfJobDefinition_EksPropertiesPodPropertiesContainersSecurityContextProperty {
	var returns *TfJobDefinition_EksPropertiesPodPropertiesContainersSecurityContextProperty
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) VolumeMounts() TfJobDefinition_EksPropertiesPodPropertiesContainersVolumeMountsPropertyList {
	var returns TfJobDefinition_EksPropertiesPodPropertiesContainersVolumeMountsPropertyList
	_jsii_.Get(
		j,
		"volumeMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) VolumeMountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeMountsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfJobDefinition_ContainersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfJobDefinition_ContainersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfJobDefinition_ContainersPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-batch.TfJobDefinition.ContainersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfJobDefinition_ContainersPropertyOutputReference_Override(t TfJobDefinition_ContainersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.TfJobDefinition.ContainersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference)SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) PutEnv(value interface{}) {
	if err := t.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnv",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) PutResources(value *TfJobDefinition_EksPropertiesPodPropertiesContainersResourcesProperty) {
	if err := t.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResources",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) PutSecurityContext(value *TfJobDefinition_EksPropertiesPodPropertiesContainersSecurityContextProperty) {
	if err := t.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) PutVolumeMounts(value interface{}) {
	if err := t.validatePutVolumeMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVolumeMounts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		t,
		"resetArgs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		t,
		"resetCommand",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		t,
		"resetEnv",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		t,
		"resetResources",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ResetVolumeMounts() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeMounts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfJobDefinition_ContainersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

