package cloudwatchlogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudwatchlogs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudwatchlogs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTransformer_TransformerConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AddKeys() AwsTransformer_AddKeysPropertyList
	// Experimental.
	AddKeysInput() interface{}
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
	CopyValue() AwsTransformer_CopyValuePropertyList
	// Experimental.
	CopyValueInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Csv() AwsTransformer_CsvPropertyList
	// Experimental.
	CsvInput() interface{}
	// Experimental.
	DateTimeConverter() AwsTransformer_DateTimeConverterPropertyList
	// Experimental.
	DateTimeConverterInput() interface{}
	// Experimental.
	DeleteKeys() AwsTransformer_DeleteKeysPropertyList
	// Experimental.
	DeleteKeysInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Grok() AwsTransformer_GrokPropertyList
	// Experimental.
	GrokInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ListToMap() AwsTransformer_ListToMapPropertyList
	// Experimental.
	ListToMapInput() interface{}
	// Experimental.
	LowerCaseString() AwsTransformer_LowerCaseStringPropertyList
	// Experimental.
	LowerCaseStringInput() interface{}
	// Experimental.
	MoveKeys() AwsTransformer_MoveKeysPropertyList
	// Experimental.
	MoveKeysInput() interface{}
	// Experimental.
	ParseCloudfront() AwsTransformer_ParseCloudfrontPropertyList
	// Experimental.
	ParseCloudfrontInput() interface{}
	// Experimental.
	ParseJson() AwsTransformer_ParseJsonPropertyList
	// Experimental.
	ParseJsonInput() interface{}
	// Experimental.
	ParseKeyValue() AwsTransformer_ParseKeyValuePropertyList
	// Experimental.
	ParseKeyValueInput() interface{}
	// Experimental.
	ParsePostgres() AwsTransformer_ParsePostgresPropertyList
	// Experimental.
	ParsePostgresInput() interface{}
	// Experimental.
	ParseRoute53() AwsTransformer_ParseRoute53PropertyList
	// Experimental.
	ParseRoute53Input() interface{}
	// Experimental.
	ParseToOcsf() AwsTransformer_ParseToOcsfPropertyList
	// Experimental.
	ParseToOcsfInput() interface{}
	// Experimental.
	ParseVpc() AwsTransformer_ParseVpcPropertyList
	// Experimental.
	ParseVpcInput() interface{}
	// Experimental.
	ParseWaf() AwsTransformer_ParseWafPropertyList
	// Experimental.
	ParseWafInput() interface{}
	// Experimental.
	RenameKeys() AwsTransformer_RenameKeysPropertyList
	// Experimental.
	RenameKeysInput() interface{}
	// Experimental.
	SplitString() AwsTransformer_SplitStringPropertyList
	// Experimental.
	SplitStringInput() interface{}
	// Experimental.
	SubstituteString() AwsTransformer_SubstituteStringPropertyList
	// Experimental.
	SubstituteStringInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrimString() AwsTransformer_TrimStringPropertyList
	// Experimental.
	TrimStringInput() interface{}
	// Experimental.
	TypeConverter() AwsTransformer_TypeConverterPropertyList
	// Experimental.
	TypeConverterInput() interface{}
	// Experimental.
	UpperCaseString() AwsTransformer_UpperCaseStringPropertyList
	// Experimental.
	UpperCaseStringInput() interface{}
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
	PutAddKeys(value interface{})
	// Experimental.
	PutCopyValue(value interface{})
	// Experimental.
	PutCsv(value interface{})
	// Experimental.
	PutDateTimeConverter(value interface{})
	// Experimental.
	PutDeleteKeys(value interface{})
	// Experimental.
	PutGrok(value interface{})
	// Experimental.
	PutListToMap(value interface{})
	// Experimental.
	PutLowerCaseString(value interface{})
	// Experimental.
	PutMoveKeys(value interface{})
	// Experimental.
	PutParseCloudfront(value interface{})
	// Experimental.
	PutParseJson(value interface{})
	// Experimental.
	PutParseKeyValue(value interface{})
	// Experimental.
	PutParsePostgres(value interface{})
	// Experimental.
	PutParseRoute53(value interface{})
	// Experimental.
	PutParseToOcsf(value interface{})
	// Experimental.
	PutParseVpc(value interface{})
	// Experimental.
	PutParseWaf(value interface{})
	// Experimental.
	PutRenameKeys(value interface{})
	// Experimental.
	PutSplitString(value interface{})
	// Experimental.
	PutSubstituteString(value interface{})
	// Experimental.
	PutTrimString(value interface{})
	// Experimental.
	PutTypeConverter(value interface{})
	// Experimental.
	PutUpperCaseString(value interface{})
	// Experimental.
	ResetAddKeys()
	// Experimental.
	ResetCopyValue()
	// Experimental.
	ResetCsv()
	// Experimental.
	ResetDateTimeConverter()
	// Experimental.
	ResetDeleteKeys()
	// Experimental.
	ResetGrok()
	// Experimental.
	ResetListToMap()
	// Experimental.
	ResetLowerCaseString()
	// Experimental.
	ResetMoveKeys()
	// Experimental.
	ResetParseCloudfront()
	// Experimental.
	ResetParseJson()
	// Experimental.
	ResetParseKeyValue()
	// Experimental.
	ResetParsePostgres()
	// Experimental.
	ResetParseRoute53()
	// Experimental.
	ResetParseToOcsf()
	// Experimental.
	ResetParseVpc()
	// Experimental.
	ResetParseWaf()
	// Experimental.
	ResetRenameKeys()
	// Experimental.
	ResetSplitString()
	// Experimental.
	ResetSubstituteString()
	// Experimental.
	ResetTrimString()
	// Experimental.
	ResetTypeConverter()
	// Experimental.
	ResetUpperCaseString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTransformer_TransformerConfigPropertyOutputReference
type jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) AddKeys() AwsTransformer_AddKeysPropertyList {
	var returns AwsTransformer_AddKeysPropertyList
	_jsii_.Get(
		j,
		"addKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) AddKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) CopyValue() AwsTransformer_CopyValuePropertyList {
	var returns AwsTransformer_CopyValuePropertyList
	_jsii_.Get(
		j,
		"copyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) CopyValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) Csv() AwsTransformer_CsvPropertyList {
	var returns AwsTransformer_CsvPropertyList
	_jsii_.Get(
		j,
		"csv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) CsvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"csvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) DateTimeConverter() AwsTransformer_DateTimeConverterPropertyList {
	var returns AwsTransformer_DateTimeConverterPropertyList
	_jsii_.Get(
		j,
		"dateTimeConverter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) DateTimeConverterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dateTimeConverterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) DeleteKeys() AwsTransformer_DeleteKeysPropertyList {
	var returns AwsTransformer_DeleteKeysPropertyList
	_jsii_.Get(
		j,
		"deleteKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) DeleteKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) Grok() AwsTransformer_GrokPropertyList {
	var returns AwsTransformer_GrokPropertyList
	_jsii_.Get(
		j,
		"grok",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) GrokInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grokInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ListToMap() AwsTransformer_ListToMapPropertyList {
	var returns AwsTransformer_ListToMapPropertyList
	_jsii_.Get(
		j,
		"listToMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ListToMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"listToMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) LowerCaseString() AwsTransformer_LowerCaseStringPropertyList {
	var returns AwsTransformer_LowerCaseStringPropertyList
	_jsii_.Get(
		j,
		"lowerCaseString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) LowerCaseStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lowerCaseStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) MoveKeys() AwsTransformer_MoveKeysPropertyList {
	var returns AwsTransformer_MoveKeysPropertyList
	_jsii_.Get(
		j,
		"moveKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) MoveKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"moveKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseCloudfront() AwsTransformer_ParseCloudfrontPropertyList {
	var returns AwsTransformer_ParseCloudfrontPropertyList
	_jsii_.Get(
		j,
		"parseCloudfront",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseCloudfrontInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseCloudfrontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseJson() AwsTransformer_ParseJsonPropertyList {
	var returns AwsTransformer_ParseJsonPropertyList
	_jsii_.Get(
		j,
		"parseJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseJsonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseKeyValue() AwsTransformer_ParseKeyValuePropertyList {
	var returns AwsTransformer_ParseKeyValuePropertyList
	_jsii_.Get(
		j,
		"parseKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseKeyValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParsePostgres() AwsTransformer_ParsePostgresPropertyList {
	var returns AwsTransformer_ParsePostgresPropertyList
	_jsii_.Get(
		j,
		"parsePostgres",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParsePostgresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parsePostgresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseRoute53() AwsTransformer_ParseRoute53PropertyList {
	var returns AwsTransformer_ParseRoute53PropertyList
	_jsii_.Get(
		j,
		"parseRoute53",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseRoute53Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseRoute53Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseToOcsf() AwsTransformer_ParseToOcsfPropertyList {
	var returns AwsTransformer_ParseToOcsfPropertyList
	_jsii_.Get(
		j,
		"parseToOcsf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseToOcsfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseToOcsfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseVpc() AwsTransformer_ParseVpcPropertyList {
	var returns AwsTransformer_ParseVpcPropertyList
	_jsii_.Get(
		j,
		"parseVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseVpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseVpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseWaf() AwsTransformer_ParseWafPropertyList {
	var returns AwsTransformer_ParseWafPropertyList
	_jsii_.Get(
		j,
		"parseWaf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ParseWafInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseWafInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) RenameKeys() AwsTransformer_RenameKeysPropertyList {
	var returns AwsTransformer_RenameKeysPropertyList
	_jsii_.Get(
		j,
		"renameKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) RenameKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renameKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) SplitString() AwsTransformer_SplitStringPropertyList {
	var returns AwsTransformer_SplitStringPropertyList
	_jsii_.Get(
		j,
		"splitString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) SplitStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) SubstituteString() AwsTransformer_SubstituteStringPropertyList {
	var returns AwsTransformer_SubstituteStringPropertyList
	_jsii_.Get(
		j,
		"substituteString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) SubstituteStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"substituteStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) TrimString() AwsTransformer_TrimStringPropertyList {
	var returns AwsTransformer_TrimStringPropertyList
	_jsii_.Get(
		j,
		"trimString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) TrimStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trimStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) TypeConverter() AwsTransformer_TypeConverterPropertyList {
	var returns AwsTransformer_TypeConverterPropertyList
	_jsii_.Get(
		j,
		"typeConverter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) TypeConverterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeConverterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) UpperCaseString() AwsTransformer_UpperCaseStringPropertyList {
	var returns AwsTransformer_UpperCaseStringPropertyList
	_jsii_.Get(
		j,
		"upperCaseString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) UpperCaseStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upperCaseStringInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTransformer_TransformerConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsTransformer_TransformerConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTransformer_TransformerConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-logs.AwsTransformer.TransformerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTransformer_TransformerConfigPropertyOutputReference_Override(a AwsTransformer_TransformerConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-logs.AwsTransformer.TransformerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutAddKeys(value interface{}) {
	if err := a.validatePutAddKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAddKeys",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutCopyValue(value interface{}) {
	if err := a.validatePutCopyValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCopyValue",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutCsv(value interface{}) {
	if err := a.validatePutCsvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCsv",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutDateTimeConverter(value interface{}) {
	if err := a.validatePutDateTimeConverterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDateTimeConverter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutDeleteKeys(value interface{}) {
	if err := a.validatePutDeleteKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeleteKeys",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutGrok(value interface{}) {
	if err := a.validatePutGrokParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGrok",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutListToMap(value interface{}) {
	if err := a.validatePutListToMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putListToMap",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutLowerCaseString(value interface{}) {
	if err := a.validatePutLowerCaseStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLowerCaseString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutMoveKeys(value interface{}) {
	if err := a.validatePutMoveKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMoveKeys",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutParseCloudfront(value interface{}) {
	if err := a.validatePutParseCloudfrontParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParseCloudfront",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutParseJson(value interface{}) {
	if err := a.validatePutParseJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParseJson",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutParseKeyValue(value interface{}) {
	if err := a.validatePutParseKeyValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParseKeyValue",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutParsePostgres(value interface{}) {
	if err := a.validatePutParsePostgresParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParsePostgres",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutParseRoute53(value interface{}) {
	if err := a.validatePutParseRoute53Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParseRoute53",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutParseToOcsf(value interface{}) {
	if err := a.validatePutParseToOcsfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParseToOcsf",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutParseVpc(value interface{}) {
	if err := a.validatePutParseVpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParseVpc",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutParseWaf(value interface{}) {
	if err := a.validatePutParseWafParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParseWaf",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutRenameKeys(value interface{}) {
	if err := a.validatePutRenameKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRenameKeys",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutSplitString(value interface{}) {
	if err := a.validatePutSplitStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSplitString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutSubstituteString(value interface{}) {
	if err := a.validatePutSubstituteStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubstituteString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutTrimString(value interface{}) {
	if err := a.validatePutTrimStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrimString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutTypeConverter(value interface{}) {
	if err := a.validatePutTypeConverterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTypeConverter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) PutUpperCaseString(value interface{}) {
	if err := a.validatePutUpperCaseStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpperCaseString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetAddKeys() {
	_jsii_.InvokeVoid(
		a,
		"resetAddKeys",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetCopyValue() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetCsv() {
	_jsii_.InvokeVoid(
		a,
		"resetCsv",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetDateTimeConverter() {
	_jsii_.InvokeVoid(
		a,
		"resetDateTimeConverter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetDeleteKeys() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteKeys",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetGrok() {
	_jsii_.InvokeVoid(
		a,
		"resetGrok",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetListToMap() {
	_jsii_.InvokeVoid(
		a,
		"resetListToMap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetLowerCaseString() {
	_jsii_.InvokeVoid(
		a,
		"resetLowerCaseString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetMoveKeys() {
	_jsii_.InvokeVoid(
		a,
		"resetMoveKeys",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetParseCloudfront() {
	_jsii_.InvokeVoid(
		a,
		"resetParseCloudfront",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetParseJson() {
	_jsii_.InvokeVoid(
		a,
		"resetParseJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetParseKeyValue() {
	_jsii_.InvokeVoid(
		a,
		"resetParseKeyValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetParsePostgres() {
	_jsii_.InvokeVoid(
		a,
		"resetParsePostgres",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetParseRoute53() {
	_jsii_.InvokeVoid(
		a,
		"resetParseRoute53",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetParseToOcsf() {
	_jsii_.InvokeVoid(
		a,
		"resetParseToOcsf",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetParseVpc() {
	_jsii_.InvokeVoid(
		a,
		"resetParseVpc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetParseWaf() {
	_jsii_.InvokeVoid(
		a,
		"resetParseWaf",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetRenameKeys() {
	_jsii_.InvokeVoid(
		a,
		"resetRenameKeys",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetSplitString() {
	_jsii_.InvokeVoid(
		a,
		"resetSplitString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetSubstituteString() {
	_jsii_.InvokeVoid(
		a,
		"resetSubstituteString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetTrimString() {
	_jsii_.InvokeVoid(
		a,
		"resetTrimString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetTypeConverter() {
	_jsii_.InvokeVoid(
		a,
		"resetTypeConverter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ResetUpperCaseString() {
	_jsii_.InvokeVoid(
		a,
		"resetUpperCaseString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTransformer_TransformerConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

