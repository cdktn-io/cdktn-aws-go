package appsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFunction_SyncConfigPropertyOutputReference interface {
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
	ConflictDetection() *string
	// Experimental.
	SetConflictDetection(val *string)
	// Experimental.
	ConflictDetectionInput() *string
	// Experimental.
	ConflictHandler() *string
	// Experimental.
	SetConflictHandler(val *string)
	// Experimental.
	ConflictHandlerInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsFunction_SyncConfigProperty
	// Experimental.
	SetInternalValue(val *AwsFunction_SyncConfigProperty)
	// Experimental.
	LambdaConflictHandlerConfig() AwsFunction_LambdaConflictHandlerConfigPropertyOutputReference
	// Experimental.
	LambdaConflictHandlerConfigInput() *AwsFunction_LambdaConflictHandlerConfigProperty
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
	PutLambdaConflictHandlerConfig(value *AwsFunction_LambdaConflictHandlerConfigProperty)
	// Experimental.
	ResetConflictDetection()
	// Experimental.
	ResetConflictHandler()
	// Experimental.
	ResetLambdaConflictHandlerConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFunction_SyncConfigPropertyOutputReference
type jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) ConflictDetection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) ConflictDetectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) ConflictHandler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) ConflictHandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) InternalValue() *AwsFunction_SyncConfigProperty {
	var returns *AwsFunction_SyncConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) LambdaConflictHandlerConfig() AwsFunction_LambdaConflictHandlerConfigPropertyOutputReference {
	var returns AwsFunction_LambdaConflictHandlerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaConflictHandlerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) LambdaConflictHandlerConfigInput() *AwsFunction_LambdaConflictHandlerConfigProperty {
	var returns *AwsFunction_LambdaConflictHandlerConfigProperty
	_jsii_.Get(
		j,
		"lambdaConflictHandlerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFunction_SyncConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFunction_SyncConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFunction_SyncConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsFunction.SyncConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFunction_SyncConfigPropertyOutputReference_Override(a AwsFunction_SyncConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsFunction.SyncConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference)SetConflictDetection(val *string) {
	if err := j.validateSetConflictDetectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conflictDetection",
		val,
	)
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference)SetConflictHandler(val *string) {
	if err := j.validateSetConflictHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conflictHandler",
		val,
	)
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference)SetInternalValue(val *AwsFunction_SyncConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) PutLambdaConflictHandlerConfig(value *AwsFunction_LambdaConflictHandlerConfigProperty) {
	if err := a.validatePutLambdaConflictHandlerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaConflictHandlerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) ResetConflictDetection() {
	_jsii_.InvokeVoid(
		a,
		"resetConflictDetection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) ResetConflictHandler() {
	_jsii_.InvokeVoid(
		a,
		"resetConflictHandler",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) ResetLambdaConflictHandlerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaConflictHandlerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFunction_SyncConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

