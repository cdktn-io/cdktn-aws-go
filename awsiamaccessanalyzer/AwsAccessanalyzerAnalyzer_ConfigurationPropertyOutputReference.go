package awsiamaccessanalyzer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiamaccessanalyzer/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiamaccessanalyzer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalAccess() AwsAccessanalyzerAnalyzer_InternalAccessPropertyOutputReference
	// Experimental.
	InternalAccessInput() *AwsAccessanalyzerAnalyzer_InternalAccessProperty
	// Experimental.
	InternalValue() *AwsAccessanalyzerAnalyzer_ConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsAccessanalyzerAnalyzer_ConfigurationProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UnusedAccess() AwsAccessanalyzerAnalyzer_UnusedAccessPropertyOutputReference
	// Experimental.
	UnusedAccessInput() *AwsAccessanalyzerAnalyzer_UnusedAccessProperty
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
	PutInternalAccess(value *AwsAccessanalyzerAnalyzer_InternalAccessProperty)
	// Experimental.
	PutUnusedAccess(value *AwsAccessanalyzerAnalyzer_UnusedAccessProperty)
	// Experimental.
	ResetInternalAccess()
	// Experimental.
	ResetUnusedAccess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference
type jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) InternalAccess() AwsAccessanalyzerAnalyzer_InternalAccessPropertyOutputReference {
	var returns AwsAccessanalyzerAnalyzer_InternalAccessPropertyOutputReference
	_jsii_.Get(
		j,
		"internalAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) InternalAccessInput() *AwsAccessanalyzerAnalyzer_InternalAccessProperty {
	var returns *AwsAccessanalyzerAnalyzer_InternalAccessProperty
	_jsii_.Get(
		j,
		"internalAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) InternalValue() *AwsAccessanalyzerAnalyzer_ConfigurationProperty {
	var returns *AwsAccessanalyzerAnalyzer_ConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) UnusedAccess() AwsAccessanalyzerAnalyzer_UnusedAccessPropertyOutputReference {
	var returns AwsAccessanalyzerAnalyzer_UnusedAccessPropertyOutputReference
	_jsii_.Get(
		j,
		"unusedAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) UnusedAccessInput() *AwsAccessanalyzerAnalyzer_UnusedAccessProperty {
	var returns *AwsAccessanalyzerAnalyzer_UnusedAccessProperty
	_jsii_.Get(
		j,
		"unusedAccessInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iam-access-analyzer.AwsAccessanalyzerAnalyzer.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference_Override(a AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iam-access-analyzer.AwsAccessanalyzerAnalyzer.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference)SetInternalValue(val *AwsAccessanalyzerAnalyzer_ConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) PutInternalAccess(value *AwsAccessanalyzerAnalyzer_InternalAccessProperty) {
	if err := a.validatePutInternalAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInternalAccess",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) PutUnusedAccess(value *AwsAccessanalyzerAnalyzer_UnusedAccessProperty) {
	if err := a.validatePutUnusedAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUnusedAccess",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) ResetInternalAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetInternalAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) ResetUnusedAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetUnusedAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAccessanalyzerAnalyzer_ConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

