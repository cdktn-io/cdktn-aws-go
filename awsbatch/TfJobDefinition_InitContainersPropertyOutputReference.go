package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbatch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbatch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfJobDefinition_InitContainersPropertyOutputReference interface {
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
	Env() TfJobDefinition_EksPropertiesPodPropertiesInitContainersEnvPropertyList
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
	Resources() TfJobDefinition_EksPropertiesPodPropertiesInitContainersResourcesPropertyOutputReference
	// Experimental.
	ResourcesInput() *TfJobDefinition_EksPropertiesPodPropertiesInitContainersResourcesProperty
	// Experimental.
	SecurityContext() TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference
	// Experimental.
	SecurityContextInput() *TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VolumeMounts() TfJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList
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
	PutResources(value *TfJobDefinition_EksPropertiesPodPropertiesInitContainersResourcesProperty)
	// Experimental.
	PutSecurityContext(value *TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty)
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

// The jsii proxy struct for TfJobDefinition_InitContainersPropertyOutputReference
type jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) Env() TfJobDefinition_EksPropertiesPodPropertiesInitContainersEnvPropertyList {
	var returns TfJobDefinition_EksPropertiesPodPropertiesInitContainersEnvPropertyList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) Resources() TfJobDefinition_EksPropertiesPodPropertiesInitContainersResourcesPropertyOutputReference {
	var returns TfJobDefinition_EksPropertiesPodPropertiesInitContainersResourcesPropertyOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ResourcesInput() *TfJobDefinition_EksPropertiesPodPropertiesInitContainersResourcesProperty {
	var returns *TfJobDefinition_EksPropertiesPodPropertiesInitContainersResourcesProperty
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) SecurityContext() TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference {
	var returns TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) SecurityContextInput() *TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty {
	var returns *TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) VolumeMounts() TfJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList {
	var returns TfJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList
	_jsii_.Get(
		j,
		"volumeMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) VolumeMountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeMountsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfJobDefinition_InitContainersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfJobDefinition_InitContainersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfJobDefinition_InitContainersPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-batch.TfJobDefinition.InitContainersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfJobDefinition_InitContainersPropertyOutputReference_Override(t TfJobDefinition_InitContainersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.TfJobDefinition.InitContainersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference)SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) PutEnv(value interface{}) {
	if err := t.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnv",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) PutResources(value *TfJobDefinition_EksPropertiesPodPropertiesInitContainersResourcesProperty) {
	if err := t.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResources",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) PutSecurityContext(value *TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty) {
	if err := t.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) PutVolumeMounts(value interface{}) {
	if err := t.validatePutVolumeMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVolumeMounts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		t,
		"resetArgs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		t,
		"resetCommand",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		t,
		"resetEnv",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		t,
		"resetResources",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ResetVolumeMounts() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeMounts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfJobDefinition_InitContainersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

