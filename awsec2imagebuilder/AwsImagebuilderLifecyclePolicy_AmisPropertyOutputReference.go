package awsec2imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IsPublic() interface{}
	// Experimental.
	SetIsPublic(val interface{})
	// Experimental.
	IsPublicInput() interface{}
	// Experimental.
	LastLaunched() AwsImagebuilderLifecyclePolicy_LastLaunchedPropertyList
	// Experimental.
	LastLaunchedInput() interface{}
	// Experimental.
	Regions() *[]*string
	// Experimental.
	SetRegions(val *[]*string)
	// Experimental.
	RegionsInput() *[]*string
	// Experimental.
	SharedAccounts() *[]*string
	// Experimental.
	SetSharedAccounts(val *[]*string)
	// Experimental.
	SharedAccountsInput() *[]*string
	// Experimental.
	TagMap() *map[string]*string
	// Experimental.
	SetTagMap(val *map[string]*string)
	// Experimental.
	TagMapInput() *map[string]*string
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
	PutLastLaunched(value interface{})
	// Experimental.
	ResetIsPublic()
	// Experimental.
	ResetLastLaunched()
	// Experimental.
	ResetRegions()
	// Experimental.
	ResetSharedAccounts()
	// Experimental.
	ResetTagMap()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference
type jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) IsPublic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) IsPublicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPublicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) LastLaunched() AwsImagebuilderLifecyclePolicy_LastLaunchedPropertyList {
	var returns AwsImagebuilderLifecyclePolicy_LastLaunchedPropertyList
	_jsii_.Get(
		j,
		"lastLaunched",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) LastLaunchedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastLaunchedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) SharedAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) SharedAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) TagMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) TagMapInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsImagebuilderLifecyclePolicy_AmisPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.AwsImagebuilderLifecyclePolicy.AmisPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference_Override(a AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.AwsImagebuilderLifecyclePolicy.AmisPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference)SetIsPublic(val interface{}) {
	if err := j.validateSetIsPublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPublic",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference)SetSharedAccounts(val *[]*string) {
	if err := j.validateSetSharedAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAccounts",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference)SetTagMap(val *map[string]*string) {
	if err := j.validateSetTagMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagMap",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) PutLastLaunched(value interface{}) {
	if err := a.validatePutLastLaunchedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLastLaunched",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) ResetIsPublic() {
	_jsii_.InvokeVoid(
		a,
		"resetIsPublic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) ResetLastLaunched() {
	_jsii_.InvokeVoid(
		a,
		"resetLastLaunched",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		a,
		"resetRegions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) ResetSharedAccounts() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedAccounts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) ResetTagMap() {
	_jsii_.InvokeVoid(
		a,
		"resetTagMap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsImagebuilderLifecyclePolicy_AmisPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

