// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapostgresqlsequences

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-postgresql-go/postgresql/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-postgresql-go/postgresql/v12/datapostgresqlsequences/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/sequences postgresql_sequences}.
type DataPostgresqlSequences interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LikeAllPatterns() *[]*string
	SetLikeAllPatterns(val *[]*string)
	LikeAllPatternsInput() *[]*string
	LikeAnyPatterns() *[]*string
	SetLikeAnyPatterns(val *[]*string)
	LikeAnyPatternsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	NotLikeAllPatterns() *[]*string
	SetNotLikeAllPatterns(val *[]*string)
	NotLikeAllPatternsInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RegexPattern() *string
	SetRegexPattern(val *string)
	RegexPatternInput() *string
	Schemas() *[]*string
	SetSchemas(val *[]*string)
	SchemasInput() *[]*string
	Sequences() DataPostgresqlSequencesSequencesList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	ResetLikeAllPatterns()
	ResetLikeAnyPatterns()
	ResetNotLikeAllPatterns()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegexPattern()
	ResetSchemas()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataPostgresqlSequences
type jsiiProxy_DataPostgresqlSequences struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataPostgresqlSequences) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) LikeAllPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAllPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) LikeAllPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAllPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) LikeAnyPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAnyPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) LikeAnyPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"likeAnyPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) NotLikeAllPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notLikeAllPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) NotLikeAllPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notLikeAllPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) RegexPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) RegexPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) Schemas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) SchemasInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) Sequences() DataPostgresqlSequencesSequencesList {
	var returns DataPostgresqlSequencesSequencesList
	_jsii_.Get(
		j,
		"sequences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPostgresqlSequences) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/sequences postgresql_sequences} Data Source.
func NewDataPostgresqlSequences(scope constructs.Construct, id *string, config *DataPostgresqlSequencesConfig) DataPostgresqlSequences {
	_init_.Initialize()

	if err := validateNewDataPostgresqlSequencesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataPostgresqlSequences{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.dataPostgresqlSequences.DataPostgresqlSequences",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/sequences postgresql_sequences} Data Source.
func NewDataPostgresqlSequences_Override(d DataPostgresqlSequences, scope constructs.Construct, id *string, config *DataPostgresqlSequencesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.dataPostgresqlSequences.DataPostgresqlSequences",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetLikeAllPatterns(val *[]*string) {
	if err := j.validateSetLikeAllPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"likeAllPatterns",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetLikeAnyPatterns(val *[]*string) {
	if err := j.validateSetLikeAnyPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"likeAnyPatterns",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetNotLikeAllPatterns(val *[]*string) {
	if err := j.validateSetNotLikeAllPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notLikeAllPatterns",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetRegexPattern(val *string) {
	if err := j.validateSetRegexPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regexPattern",
		val,
	)
}

func (j *jsiiProxy_DataPostgresqlSequences)SetSchemas(val *[]*string) {
	if err := j.validateSetSchemasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemas",
		val,
	)
}

// Generates CDKTF code for importing a DataPostgresqlSequences resource upon running "cdktf plan <stack-name>".
func DataPostgresqlSequences_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataPostgresqlSequences_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.dataPostgresqlSequences.DataPostgresqlSequences",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataPostgresqlSequences_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPostgresqlSequences_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.dataPostgresqlSequences.DataPostgresqlSequences",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataPostgresqlSequences_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPostgresqlSequences_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.dataPostgresqlSequences.DataPostgresqlSequences",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataPostgresqlSequences_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPostgresqlSequences_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.dataPostgresqlSequences.DataPostgresqlSequences",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataPostgresqlSequences_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.dataPostgresqlSequences.DataPostgresqlSequences",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataPostgresqlSequences) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataPostgresqlSequences) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSequences) ResetLikeAllPatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetLikeAllPatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSequences) ResetLikeAnyPatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetLikeAnyPatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSequences) ResetNotLikeAllPatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetNotLikeAllPatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSequences) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSequences) ResetRegexPattern() {
	_jsii_.InvokeVoid(
		d,
		"resetRegexPattern",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSequences) ResetSchemas() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPostgresqlSequences) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPostgresqlSequences) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

