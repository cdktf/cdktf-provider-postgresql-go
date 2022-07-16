// Prebuilt postgresql Provider for Terraform CDK (cdktf)
package postgresql

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-postgresql-go/postgresql/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-postgresql-go/postgresql/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/postgresql/r/database postgresql_database}.
type Database interface {
	cdktf.TerraformResource
	AllowConnections() interface{}
	SetAllowConnections(val interface{})
	AllowConnectionsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConnectionLimit() *float64
	SetConnectionLimit(val *float64)
	ConnectionLimitInput() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsTemplate() interface{}
	SetIsTemplate(val interface{})
	IsTemplateInput() interface{}
	LcCollate() *string
	SetLcCollate(val *string)
	LcCollateInput() *string
	LcCtype() *string
	SetLcCtype(val *string)
	LcCtypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	TablespaceName() *string
	SetTablespaceName(val *string)
	TablespaceNameInput() *string
	Template() *string
	SetTemplate(val *string)
	TemplateInput() *string
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
	ResetAllowConnections()
	ResetConnectionLimit()
	ResetEncoding()
	ResetId()
	ResetIsTemplate()
	ResetLcCollate()
	ResetLcCtype()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwner()
	ResetTablespaceName()
	ResetTemplate()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Database
type jsiiProxy_Database struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Database) AllowConnections() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) AllowConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) ConnectionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) ConnectionLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) IsTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) IsTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) LcCollate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lcCollate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) LcCollateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lcCollateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) LcCtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lcCtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) LcCtypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lcCtypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TablespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TablespaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablespaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Template() *string {
	var returns *string
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/database postgresql_database} Resource.
func NewDatabase(scope constructs.Construct, id *string, config *DatabaseConfig) Database {
	_init_.Initialize()

	j := jsiiProxy_Database{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.Database",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/database postgresql_database} Resource.
func NewDatabase_Override(d Database, scope constructs.Construct, id *string, config *DatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.Database",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_Database) SetAllowConnections(val interface{}) {
	_jsii_.Set(
		j,
		"allowConnections",
		val,
	)
}

func (j *jsiiProxy_Database) SetConnectionLimit(val *float64) {
	_jsii_.Set(
		j,
		"connectionLimit",
		val,
	)
}

func (j *jsiiProxy_Database) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Database) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Database) SetEncoding(val *string) {
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_Database) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Database) SetIsTemplate(val interface{}) {
	_jsii_.Set(
		j,
		"isTemplate",
		val,
	)
}

func (j *jsiiProxy_Database) SetLcCollate(val *string) {
	_jsii_.Set(
		j,
		"lcCollate",
		val,
	)
}

func (j *jsiiProxy_Database) SetLcCtype(val *string) {
	_jsii_.Set(
		j,
		"lcCtype",
		val,
	)
}

func (j *jsiiProxy_Database) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Database) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Database) SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_Database) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Database) SetTablespaceName(val *string) {
	_jsii_.Set(
		j,
		"tablespaceName",
		val,
	)
}

func (j *jsiiProxy_Database) SetTemplate(val *string) {
	_jsii_.Set(
		j,
		"template",
		val,
	)
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
func Database_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.Database",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Database_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.Database",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_Database) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_Database) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_Database) ResetAllowConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetConnectionLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetEncoding() {
	_jsii_.InvokeVoid(
		d,
		"resetEncoding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetIsTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetIsTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetLcCollate() {
	_jsii_.InvokeVoid(
		d,
		"resetLcCollate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetLcCtype() {
	_jsii_.InvokeVoid(
		d,
		"resetLcCtype",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetTablespaceName() {
	_jsii_.InvokeVoid(
		d,
		"resetTablespaceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DatabaseConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// The PostgreSQL database name to connect to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#name Database#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If false then no one can connect to this database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#allow_connections Database#allow_connections}
	AllowConnections interface{} `field:"optional" json:"allowConnections" yaml:"allowConnections"`
	// How many concurrent connections can be made to this database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#connection_limit Database#connection_limit}
	ConnectionLimit *float64 `field:"optional" json:"connectionLimit" yaml:"connectionLimit"`
	// Character set encoding to use in the new database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#encoding Database#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#id Database#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, then this database can be cloned by any user with CREATEDB privileges.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#is_template Database#is_template}
	IsTemplate interface{} `field:"optional" json:"isTemplate" yaml:"isTemplate"`
	// Collation order (LC_COLLATE) to use in the new database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#lc_collate Database#lc_collate}
	LcCollate *string `field:"optional" json:"lcCollate" yaml:"lcCollate"`
	// Character classification (LC_CTYPE) to use in the new database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#lc_ctype Database#lc_ctype}
	LcCtype *string `field:"optional" json:"lcCtype" yaml:"lcCtype"`
	// The ROLE which owns the database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#owner Database#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The name of the tablespace that will be associated with the new database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#tablespace_name Database#tablespace_name}
	TablespaceName *string `field:"optional" json:"tablespaceName" yaml:"tablespaceName"`
	// The name of the template from which to create the new database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/database#template Database#template}
	Template *string `field:"optional" json:"template" yaml:"template"`
}

// Represents a {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges postgresql_default_privileges}.
type DefaultPrivileges interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	// The tree node.
	Node() constructs.Node
	ObjectType() *string
	SetObjectType(val *string)
	ObjectTypeInput() *string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	Privileges() *[]*string
	SetPrivileges(val *[]*string)
	PrivilegesInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WithGrantOption() interface{}
	SetWithGrantOption(val interface{})
	WithGrantOptionInput() interface{}
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSchema()
	ResetWithGrantOption()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DefaultPrivileges
type jsiiProxy_DefaultPrivileges struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DefaultPrivileges) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) ObjectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) ObjectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) Privileges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privileges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) PrivilegesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privilegesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) WithGrantOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPrivileges) WithGrantOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOptionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges postgresql_default_privileges} Resource.
func NewDefaultPrivileges(scope constructs.Construct, id *string, config *DefaultPrivilegesConfig) DefaultPrivileges {
	_init_.Initialize()

	j := jsiiProxy_DefaultPrivileges{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.DefaultPrivileges",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges postgresql_default_privileges} Resource.
func NewDefaultPrivileges_Override(d DefaultPrivileges, scope constructs.Construct, id *string, config *DefaultPrivilegesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.DefaultPrivileges",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetObjectType(val *string) {
	_jsii_.Set(
		j,
		"objectType",
		val,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetPrivileges(val *[]*string) {
	_jsii_.Set(
		j,
		"privileges",
		val,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetSchema(val *string) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_DefaultPrivileges) SetWithGrantOption(val interface{}) {
	_jsii_.Set(
		j,
		"withGrantOption",
		val,
	)
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
func DefaultPrivileges_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.DefaultPrivileges",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DefaultPrivileges_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.DefaultPrivileges",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DefaultPrivileges) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DefaultPrivileges) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DefaultPrivileges) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultPrivileges) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultPrivileges) ResetSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultPrivileges) ResetWithGrantOption() {
	_jsii_.InvokeVoid(
		d,
		"resetWithGrantOption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DefaultPrivileges) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultPrivileges) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DefaultPrivilegesConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// The database to grant default privileges for this role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#database DefaultPrivileges#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#object_type DefaultPrivileges#object_type}
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Target role for which to alter default privileges.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#owner DefaultPrivileges#owner}
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The list of privileges to apply as default privileges.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#privileges DefaultPrivileges#privileges}
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
	// The name of the role to which grant default privileges on.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#role DefaultPrivileges#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#id DefaultPrivileges#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The database schema to set default privileges for this role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#schema DefaultPrivileges#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Permit the grant recipient to grant it to others.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/default_privileges#with_grant_option DefaultPrivileges#with_grant_option}
	WithGrantOption interface{} `field:"optional" json:"withGrantOption" yaml:"withGrantOption"`
}

// Represents a {@link https://www.terraform.io/docs/providers/postgresql/r/extension postgresql_extension}.
type Extension interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreateCascade() interface{}
	SetCreateCascade(val interface{})
	CreateCascadeInput() interface{}
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DropCascade() interface{}
	SetDropCascade(val interface{})
	DropCascadeInput() interface{}
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetCreateCascade()
	ResetDatabase()
	ResetDropCascade()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSchema()
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Extension
type jsiiProxy_Extension struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Extension) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) CreateCascade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createCascade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) CreateCascadeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createCascadeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) DropCascade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropCascade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) DropCascadeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropCascadeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extension) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/extension postgresql_extension} Resource.
func NewExtension(scope constructs.Construct, id *string, config *ExtensionConfig) Extension {
	_init_.Initialize()

	j := jsiiProxy_Extension{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.Extension",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/extension postgresql_extension} Resource.
func NewExtension_Override(e Extension, scope constructs.Construct, id *string, config *ExtensionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.Extension",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Extension) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Extension) SetCreateCascade(val interface{}) {
	_jsii_.Set(
		j,
		"createCascade",
		val,
	)
}

func (j *jsiiProxy_Extension) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_Extension) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Extension) SetDropCascade(val interface{}) {
	_jsii_.Set(
		j,
		"dropCascade",
		val,
	)
}

func (j *jsiiProxy_Extension) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Extension) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Extension) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Extension) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Extension) SetSchema(val *string) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_Extension) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
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
func Extension_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.Extension",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Extension_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.Extension",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Extension) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Extension) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Extension) ResetCreateCascade() {
	_jsii_.InvokeVoid(
		e,
		"resetCreateCascade",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Extension) ResetDatabase() {
	_jsii_.InvokeVoid(
		e,
		"resetDatabase",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Extension) ResetDropCascade() {
	_jsii_.InvokeVoid(
		e,
		"resetDropCascade",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Extension) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Extension) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Extension) ResetSchema() {
	_jsii_.InvokeVoid(
		e,
		"resetSchema",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Extension) ResetVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Extension) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extension) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ExtensionConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#name Extension#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// When true, will also create any extensions that this extension depends on that are not already installed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#create_cascade Extension#create_cascade}
	CreateCascade interface{} `field:"optional" json:"createCascade" yaml:"createCascade"`
	// Sets the database to add the extension to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#database Extension#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#drop_cascade Extension#drop_cascade}
	DropCascade interface{} `field:"optional" json:"dropCascade" yaml:"dropCascade"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#id Extension#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Sets the schema of an extension.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#schema Extension#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Sets the version number of the extension.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#version Extension#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// Represents a {@link https://www.terraform.io/docs/providers/postgresql/r/function postgresql_function}.
type Function interface {
	cdktf.TerraformResource
	Arg() FunctionArgList
	ArgInput() interface{}
	Body() *string
	SetBody(val *string)
	BodyInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DropCascade() interface{}
	SetDropCascade(val interface{})
	DropCascadeInput() interface{}
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Returns() *string
	SetReturns(val *string)
	ReturnsInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
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
	PutArg(value interface{})
	ResetArg()
	ResetDropCascade()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReturns()
	ResetSchema()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Function
type jsiiProxy_Function struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Function) Arg() FunctionArgList {
	var returns FunctionArgList
	_jsii_.Get(
		j,
		"arg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) ArgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) BodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) DropCascade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropCascade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) DropCascadeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropCascadeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Returns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) ReturnsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/function postgresql_function} Resource.
func NewFunction(scope constructs.Construct, id *string, config *FunctionConfig) Function {
	_init_.Initialize()

	j := jsiiProxy_Function{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.Function",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/function postgresql_function} Resource.
func NewFunction_Override(f Function, scope constructs.Construct, id *string, config *FunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.Function",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_Function) SetBody(val *string) {
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_Function) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Function) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Function) SetDropCascade(val interface{}) {
	_jsii_.Set(
		j,
		"dropCascade",
		val,
	)
}

func (j *jsiiProxy_Function) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Function) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Function) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Function) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Function) SetReturns(val *string) {
	_jsii_.Set(
		j,
		"returns",
		val,
	)
}

func (j *jsiiProxy_Function) SetSchema(val *string) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
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
func Function_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.Function",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Function_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.Function",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_Function) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_Function) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_Function) PutArg(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putArg",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_Function) ResetArg() {
	_jsii_.InvokeVoid(
		f,
		"resetArg",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Function) ResetDropCascade() {
	_jsii_.InvokeVoid(
		f,
		"resetDropCascade",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Function) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Function) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Function) ResetReturns() {
	_jsii_.InvokeVoid(
		f,
		"resetReturns",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Function) ResetSchema() {
	_jsii_.InvokeVoid(
		f,
		"resetSchema",
		nil, // no parameters
	)
}

func (f *jsiiProxy_Function) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Function) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FunctionArg struct {
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#type Function#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// An expression to be used as default value if the parameter is not specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#default Function#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// The argument mode. One of: IN, OUT, INOUT, or VARIADIC.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#mode Function#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The argument name. The name may be required for some languages or depending on the argument mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#name Function#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

type FunctionArgList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FunctionArgOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionArgList
type jsiiProxy_FunctionArgList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FunctionArgList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFunctionArgList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FunctionArgList {
	_init_.Initialize()

	j := jsiiProxy_FunctionArgList{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.FunctionArgList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFunctionArgList_Override(f FunctionArgList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.FunctionArgList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FunctionArgList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionArgList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionArgList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FunctionArgList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FunctionArgList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgList) Get(index *float64) FunctionArgOutputReference {
	var returns FunctionArgOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FunctionArgOutputReference interface {
	cdktf.ComplexObject
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
	Default() *string
	SetDefault(val *string)
	DefaultInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDefault()
	ResetMode()
	ResetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionArgOutputReference
type jsiiProxy_FunctionArgOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FunctionArgOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) Default() *string {
	var returns *string
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) DefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionArgOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewFunctionArgOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FunctionArgOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FunctionArgOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.FunctionArgOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFunctionArgOutputReference_Override(f FunctionArgOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.FunctionArgOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FunctionArgOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FunctionArgOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FunctionArgOutputReference) SetDefault(val *string) {
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_FunctionArgOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionArgOutputReference) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_FunctionArgOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FunctionArgOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionArgOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FunctionArgOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (f *jsiiProxy_FunctionArgOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) ResetDefault() {
	_jsii_.InvokeVoid(
		f,
		"resetDefault",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionArgOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		f,
		"resetMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionArgOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		f,
		"resetName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionArgOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionArgOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FunctionConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Body of the function.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#body Function#body}
	Body *string `field:"required" json:"body" yaml:"body"`
	// Name of the function.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#name Function#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// arg block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#arg Function#arg}
	Arg interface{} `field:"optional" json:"arg" yaml:"arg"`
	// Automatically drop objects that depend on the function (such as operators or triggers), and in turn all objects that depend on those objects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#drop_cascade Function#drop_cascade}
	DropCascade interface{} `field:"optional" json:"dropCascade" yaml:"dropCascade"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#id Function#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Function return type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#returns Function#returns}
	Returns *string `field:"optional" json:"returns" yaml:"returns"`
	// Schema where the function is located. If not specified, the provider default schema is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#schema Function#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

// Represents a {@link https://www.terraform.io/docs/providers/postgresql/r/grant postgresql_grant}.
type Grant interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	// The tree node.
	Node() constructs.Node
	Objects() *[]*string
	SetObjects(val *[]*string)
	ObjectsInput() *[]*string
	ObjectType() *string
	SetObjectType(val *string)
	ObjectTypeInput() *string
	Privileges() *[]*string
	SetPrivileges(val *[]*string)
	PrivilegesInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WithGrantOption() interface{}
	SetWithGrantOption(val interface{})
	WithGrantOptionInput() interface{}
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
	ResetObjects()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSchema()
	ResetWithGrantOption()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Grant
type jsiiProxy_Grant struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Grant) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Objects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) ObjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) ObjectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) ObjectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Privileges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privileges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) PrivilegesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privilegesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) WithGrantOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Grant) WithGrantOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withGrantOptionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/grant postgresql_grant} Resource.
func NewGrant(scope constructs.Construct, id *string, config *GrantConfig) Grant {
	_init_.Initialize()

	j := jsiiProxy_Grant{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.Grant",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/grant postgresql_grant} Resource.
func NewGrant_Override(g Grant, scope constructs.Construct, id *string, config *GrantConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.Grant",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_Grant) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Grant) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_Grant) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Grant) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Grant) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Grant) SetObjects(val *[]*string) {
	_jsii_.Set(
		j,
		"objects",
		val,
	)
}

func (j *jsiiProxy_Grant) SetObjectType(val *string) {
	_jsii_.Set(
		j,
		"objectType",
		val,
	)
}

func (j *jsiiProxy_Grant) SetPrivileges(val *[]*string) {
	_jsii_.Set(
		j,
		"privileges",
		val,
	)
}

func (j *jsiiProxy_Grant) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Grant) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_Grant) SetSchema(val *string) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_Grant) SetWithGrantOption(val interface{}) {
	_jsii_.Set(
		j,
		"withGrantOption",
		val,
	)
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
func Grant_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.Grant",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Grant_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.Grant",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_Grant) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_Grant) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_Grant) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grant) ResetObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grant) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grant) ResetSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grant) ResetWithGrantOption() {
	_jsii_.InvokeVoid(
		g,
		"resetWithGrantOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Grant) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grant) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GrantConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// The database to grant privileges on for this role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant#database Grant#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The PostgreSQL object type to grant the privileges on (one of: database, function, procedure, routine, schema, sequence, table, foreign_data_wrapper, foreign_server).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant#object_type Grant#object_type}
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// The list of privileges to grant.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant#privileges Grant#privileges}
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
	// The name of the role to grant privileges on.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant#role Grant#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant#id Grant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The specific objects to grant privileges on for this role (empty means all objects of the requested type).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant#objects Grant#objects}
	Objects *[]*string `field:"optional" json:"objects" yaml:"objects"`
	// The database schema to grant privileges on for this role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant#schema Grant#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Permit the grant recipient to grant it to others.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant#with_grant_option Grant#with_grant_option}
	WithGrantOption interface{} `field:"optional" json:"withGrantOption" yaml:"withGrantOption"`
}

// Represents a {@link https://www.terraform.io/docs/providers/postgresql/r/grant_role postgresql_grant_role}.
type GrantRole interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GrantRole() *string
	SetGrantRole(val *string)
	GrantRoleInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WithAdminOption() interface{}
	SetWithAdminOption(val interface{})
	WithAdminOptionInput() interface{}
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetWithAdminOption()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GrantRole
type jsiiProxy_GrantRole struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GrantRole) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) GrantRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) GrantRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) WithAdminOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withAdminOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantRole) WithAdminOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withAdminOptionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/grant_role postgresql_grant_role} Resource.
func NewGrantRole(scope constructs.Construct, id *string, config *GrantRoleConfig) GrantRole {
	_init_.Initialize()

	j := jsiiProxy_GrantRole{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.GrantRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/grant_role postgresql_grant_role} Resource.
func NewGrantRole_Override(g GrantRole, scope constructs.Construct, id *string, config *GrantRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.GrantRole",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GrantRole) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GrantRole) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GrantRole) SetGrantRole(val *string) {
	_jsii_.Set(
		j,
		"grantRole",
		val,
	)
}

func (j *jsiiProxy_GrantRole) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GrantRole) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GrantRole) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GrantRole) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_GrantRole) SetWithAdminOption(val interface{}) {
	_jsii_.Set(
		j,
		"withAdminOption",
		val,
	)
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
func GrantRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.GrantRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GrantRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.GrantRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GrantRole) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GrantRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GrantRole) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantRole) ResetWithAdminOption() {
	_jsii_.InvokeVoid(
		g,
		"resetWithAdminOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type GrantRoleConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// The name of the role that is granted to role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant_role#grant_role GrantRole#grant_role}
	GrantRole *string `field:"required" json:"grantRole" yaml:"grantRole"`
	// The name of the role to grant grant_role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant_role#role GrantRole#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant_role#id GrantRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Permit the grant recipient to grant it to others.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/grant_role#with_admin_option GrantRole#with_admin_option}
	WithAdminOption interface{} `field:"optional" json:"withAdminOption" yaml:"withAdminOption"`
}

// Represents a {@link https://www.terraform.io/docs/providers/postgresql/r/physical_replication_slot postgresql_physical_replication_slot}.
type PhysicalReplicationSlot interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PhysicalReplicationSlot
type jsiiProxy_PhysicalReplicationSlot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PhysicalReplicationSlot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalReplicationSlot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/physical_replication_slot postgresql_physical_replication_slot} Resource.
func NewPhysicalReplicationSlot(scope constructs.Construct, id *string, config *PhysicalReplicationSlotConfig) PhysicalReplicationSlot {
	_init_.Initialize()

	j := jsiiProxy_PhysicalReplicationSlot{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.PhysicalReplicationSlot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/physical_replication_slot postgresql_physical_replication_slot} Resource.
func NewPhysicalReplicationSlot_Override(p PhysicalReplicationSlot, scope constructs.Construct, id *string, config *PhysicalReplicationSlotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.PhysicalReplicationSlot",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PhysicalReplicationSlot) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PhysicalReplicationSlot) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PhysicalReplicationSlot) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PhysicalReplicationSlot) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PhysicalReplicationSlot) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PhysicalReplicationSlot) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
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
func PhysicalReplicationSlot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.PhysicalReplicationSlot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PhysicalReplicationSlot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.PhysicalReplicationSlot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PhysicalReplicationSlot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PhysicalReplicationSlot) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PhysicalReplicationSlot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PhysicalReplicationSlot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalReplicationSlot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PhysicalReplicationSlotConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/physical_replication_slot#name PhysicalReplicationSlot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/physical_replication_slot#id PhysicalReplicationSlot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

// Represents a {@link https://www.terraform.io/docs/providers/postgresql postgresql}.
type PostgresqlProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AwsRdsIamAuth() interface{}
	SetAwsRdsIamAuth(val interface{})
	AwsRdsIamAuthInput() interface{}
	AwsRdsIamProfile() *string
	SetAwsRdsIamProfile(val *string)
	AwsRdsIamProfileInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Clientcert() *PostgresqlProviderClientcert
	SetClientcert(val *PostgresqlProviderClientcert)
	ClientcertInput() *PostgresqlProviderClientcert
	ConnectTimeout() *float64
	SetConnectTimeout(val *float64)
	ConnectTimeoutInput() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	DatabaseUsername() *string
	SetDatabaseUsername(val *string)
	DatabaseUsernameInput() *string
	ExpectedVersion() *string
	SetExpectedVersion(val *string)
	ExpectedVersionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	MaxConnections() *float64
	SetMaxConnections(val *float64)
	MaxConnectionsInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	// Experimental.
	RawOverrides() interface{}
	Scheme() *string
	SetScheme(val *string)
	SchemeInput() *string
	Sslmode() *string
	SetSslmode(val *string)
	SslMode() *string
	SetSslMode(val *string)
	SslmodeInput() *string
	SslModeInput() *string
	Sslrootcert() *string
	SetSslrootcert(val *string)
	SslrootcertInput() *string
	Superuser() interface{}
	SetSuperuser(val interface{})
	SuperuserInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetAwsRdsIamAuth()
	ResetAwsRdsIamProfile()
	ResetClientcert()
	ResetConnectTimeout()
	ResetDatabase()
	ResetDatabaseUsername()
	ResetExpectedVersion()
	ResetHost()
	ResetMaxConnections()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPort()
	ResetScheme()
	ResetSslmode()
	ResetSslMode()
	ResetSslrootcert()
	ResetSuperuser()
	ResetUsername()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PostgresqlProvider
type jsiiProxy_PostgresqlProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_PostgresqlProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) AwsRdsIamAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsRdsIamAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) AwsRdsIamAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsRdsIamAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) AwsRdsIamProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRdsIamProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) AwsRdsIamProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRdsIamProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Clientcert() *PostgresqlProviderClientcert {
	var returns *PostgresqlProviderClientcert
	_jsii_.Get(
		j,
		"clientcert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) ClientcertInput() *PostgresqlProviderClientcert {
	var returns *PostgresqlProviderClientcert
	_jsii_.Get(
		j,
		"clientcertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) ConnectTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) ConnectTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) DatabaseUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) DatabaseUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) ExpectedVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) ExpectedVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) MaxConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) MaxConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Scheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) SchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Sslmode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslmode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) SslmodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslmodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) SslModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Sslrootcert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslrootcert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) SslrootcertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslrootcertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Superuser() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"superuser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) SuperuserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"superuserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlProvider) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql postgresql} Resource.
func NewPostgresqlProvider(scope constructs.Construct, id *string, config *PostgresqlProviderConfig) PostgresqlProvider {
	_init_.Initialize()

	j := jsiiProxy_PostgresqlProvider{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.PostgresqlProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql postgresql} Resource.
func NewPostgresqlProvider_Override(p PostgresqlProvider, scope constructs.Construct, id *string, config *PostgresqlProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.PostgresqlProvider",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetAwsRdsIamAuth(val interface{}) {
	_jsii_.Set(
		j,
		"awsRdsIamAuth",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetAwsRdsIamProfile(val *string) {
	_jsii_.Set(
		j,
		"awsRdsIamProfile",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetClientcert(val *PostgresqlProviderClientcert) {
	_jsii_.Set(
		j,
		"clientcert",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetConnectTimeout(val *float64) {
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetDatabaseUsername(val *string) {
	_jsii_.Set(
		j,
		"databaseUsername",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetExpectedVersion(val *string) {
	_jsii_.Set(
		j,
		"expectedVersion",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetMaxConnections(val *float64) {
	_jsii_.Set(
		j,
		"maxConnections",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetScheme(val *string) {
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetSslmode(val *string) {
	_jsii_.Set(
		j,
		"sslmode",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetSslMode(val *string) {
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetSslrootcert(val *string) {
	_jsii_.Set(
		j,
		"sslrootcert",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetSuperuser(val interface{}) {
	_jsii_.Set(
		j,
		"superuser",
		val,
	)
}

func (j *jsiiProxy_PostgresqlProvider) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
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
func PostgresqlProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.PostgresqlProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PostgresqlProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.PostgresqlProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PostgresqlProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PostgresqlProvider) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		p,
		"resetAlias",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetAwsRdsIamAuth() {
	_jsii_.InvokeVoid(
		p,
		"resetAwsRdsIamAuth",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetAwsRdsIamProfile() {
	_jsii_.InvokeVoid(
		p,
		"resetAwsRdsIamProfile",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetClientcert() {
	_jsii_.InvokeVoid(
		p,
		"resetClientcert",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		p,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetDatabase() {
	_jsii_.InvokeVoid(
		p,
		"resetDatabase",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetDatabaseUsername() {
	_jsii_.InvokeVoid(
		p,
		"resetDatabaseUsername",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetExpectedVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetExpectedVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetHost() {
	_jsii_.InvokeVoid(
		p,
		"resetHost",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetMaxConnections() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxConnections",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetPassword() {
	_jsii_.InvokeVoid(
		p,
		"resetPassword",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetPort() {
	_jsii_.InvokeVoid(
		p,
		"resetPort",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetScheme() {
	_jsii_.InvokeVoid(
		p,
		"resetScheme",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetSslmode() {
	_jsii_.InvokeVoid(
		p,
		"resetSslmode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetSslMode() {
	_jsii_.InvokeVoid(
		p,
		"resetSslMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetSslrootcert() {
	_jsii_.InvokeVoid(
		p,
		"resetSslrootcert",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetSuperuser() {
	_jsii_.InvokeVoid(
		p,
		"resetSuperuser",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) ResetUsername() {
	_jsii_.InvokeVoid(
		p,
		"resetUsername",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PostgresqlProviderClientcert struct {
	// The SSL client certificate file path. The file must contain PEM encoded data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#cert PostgresqlProvider#cert}
	Cert *string `field:"required" json:"cert" yaml:"cert"`
	// The SSL client certificate private key file path. The file must contain PEM encoded data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#key PostgresqlProvider#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

type PostgresqlProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#alias PostgresqlProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Use rds_iam instead of password authentication (see: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#aws_rds_iam_auth PostgresqlProvider#aws_rds_iam_auth}
	AwsRdsIamAuth interface{} `field:"optional" json:"awsRdsIamAuth" yaml:"awsRdsIamAuth"`
	// AWS profile to use for IAM auth.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#aws_rds_iam_profile PostgresqlProvider#aws_rds_iam_profile}
	AwsRdsIamProfile *string `field:"optional" json:"awsRdsIamProfile" yaml:"awsRdsIamProfile"`
	// clientcert block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#clientcert PostgresqlProvider#clientcert}
	Clientcert *PostgresqlProviderClientcert `field:"optional" json:"clientcert" yaml:"clientcert"`
	// Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#connect_timeout PostgresqlProvider#connect_timeout}
	ConnectTimeout *float64 `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// The name of the database to connect to in order to conenct to (defaults to `postgres`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#database PostgresqlProvider#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Database username associated to the connected user (for user name maps).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#database_username PostgresqlProvider#database_username}
	DatabaseUsername *string `field:"optional" json:"databaseUsername" yaml:"databaseUsername"`
	// Specify the expected version of PostgreSQL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#expected_version PostgresqlProvider#expected_version}
	ExpectedVersion *string `field:"optional" json:"expectedVersion" yaml:"expectedVersion"`
	// Name of PostgreSQL server address to connect to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#host PostgresqlProvider#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Maximum number of connections to establish to the database. Zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#max_connections PostgresqlProvider#max_connections}
	MaxConnections *float64 `field:"optional" json:"maxConnections" yaml:"maxConnections"`
	// Password to be used if the PostgreSQL server demands password authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#password PostgresqlProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#port PostgresqlProvider#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#scheme PostgresqlProvider#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the PostgreSQL server.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#sslmode PostgresqlProvider#sslmode}
	Sslmode *string `field:"optional" json:"sslmode" yaml:"sslmode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#ssl_mode PostgresqlProvider#ssl_mode}.
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// The SSL server root certificate file path. The file must contain PEM encoded data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#sslrootcert PostgresqlProvider#sslrootcert}
	Sslrootcert *string `field:"optional" json:"sslrootcert" yaml:"sslrootcert"`
	// Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.: Refreshing state password from Postgres).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#superuser PostgresqlProvider#superuser}
	Superuser interface{} `field:"optional" json:"superuser" yaml:"superuser"`
	// PostgreSQL user name to connect as.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql#username PostgresqlProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Represents a {@link https://www.terraform.io/docs/providers/postgresql/r/publication postgresql_publication}.
type Publication interface {
	cdktf.TerraformResource
	AllTables() interface{}
	SetAllTables(val interface{})
	AllTablesInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DropCascade() interface{}
	SetDropCascade(val interface{})
	DropCascadeInput() interface{}
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	PublishParam() *[]*string
	SetPublishParam(val *[]*string)
	PublishParamInput() *[]*string
	PublishViaPartitionRootParam() interface{}
	SetPublishViaPartitionRootParam(val interface{})
	PublishViaPartitionRootParamInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Tables() *[]*string
	SetTables(val *[]*string)
	TablesInput() *[]*string
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
	ResetAllTables()
	ResetDatabase()
	ResetDropCascade()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwner()
	ResetPublishParam()
	ResetPublishViaPartitionRootParam()
	ResetTables()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Publication
type jsiiProxy_Publication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Publication) AllTables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) AllTablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) DropCascade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropCascade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) DropCascadeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropCascadeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) PublishParam() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publishParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) PublishParamInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publishParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) PublishViaPartitionRootParam() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishViaPartitionRootParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) PublishViaPartitionRootParamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishViaPartitionRootParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) Tables() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) TablesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/publication postgresql_publication} Resource.
func NewPublication(scope constructs.Construct, id *string, config *PublicationConfig) Publication {
	_init_.Initialize()

	j := jsiiProxy_Publication{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.Publication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/publication postgresql_publication} Resource.
func NewPublication_Override(p Publication, scope constructs.Construct, id *string, config *PublicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.Publication",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_Publication) SetAllTables(val interface{}) {
	_jsii_.Set(
		j,
		"allTables",
		val,
	)
}

func (j *jsiiProxy_Publication) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Publication) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_Publication) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Publication) SetDropCascade(val interface{}) {
	_jsii_.Set(
		j,
		"dropCascade",
		val,
	)
}

func (j *jsiiProxy_Publication) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Publication) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Publication) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Publication) SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_Publication) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Publication) SetPublishParam(val *[]*string) {
	_jsii_.Set(
		j,
		"publishParam",
		val,
	)
}

func (j *jsiiProxy_Publication) SetPublishViaPartitionRootParam(val interface{}) {
	_jsii_.Set(
		j,
		"publishViaPartitionRootParam",
		val,
	)
}

func (j *jsiiProxy_Publication) SetTables(val *[]*string) {
	_jsii_.Set(
		j,
		"tables",
		val,
	)
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
func Publication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.Publication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Publication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.Publication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Publication) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_Publication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_Publication) ResetAllTables() {
	_jsii_.InvokeVoid(
		p,
		"resetAllTables",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publication) ResetDatabase() {
	_jsii_.InvokeVoid(
		p,
		"resetDatabase",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publication) ResetDropCascade() {
	_jsii_.InvokeVoid(
		p,
		"resetDropCascade",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publication) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publication) ResetOwner() {
	_jsii_.InvokeVoid(
		p,
		"resetOwner",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publication) ResetPublishParam() {
	_jsii_.InvokeVoid(
		p,
		"resetPublishParam",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publication) ResetPublishViaPartitionRootParam() {
	_jsii_.InvokeVoid(
		p,
		"resetPublishViaPartitionRootParam",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publication) ResetTables() {
	_jsii_.InvokeVoid(
		p,
		"resetTables",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Publication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Publication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PublicationConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/publication#name Publication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Sets the tables list to publish to ALL tables.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/publication#all_tables Publication#all_tables}
	AllTables interface{} `field:"optional" json:"allTables" yaml:"allTables"`
	// Sets the database to add the publication for.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/publication#database Publication#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// When true, will also drop all the objects that depend on the publication, and in turn all objects that depend on those objects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/publication#drop_cascade Publication#drop_cascade}
	DropCascade interface{} `field:"optional" json:"dropCascade" yaml:"dropCascade"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/publication#id Publication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Sets the owner of the publication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/publication#owner Publication#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Sets which DML operations will be published.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/publication#publish_param Publication#publish_param}
	PublishParam *[]*string `field:"optional" json:"publishParam" yaml:"publishParam"`
	// Sets whether changes in a partitioned table using the identity and schema of the partitioned table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/publication#publish_via_partition_root_param Publication#publish_via_partition_root_param}
	PublishViaPartitionRootParam interface{} `field:"optional" json:"publishViaPartitionRootParam" yaml:"publishViaPartitionRootParam"`
	// Sets the tables list to publish.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/publication#tables Publication#tables}
	Tables *[]*string `field:"optional" json:"tables" yaml:"tables"`
}

// Represents a {@link https://www.terraform.io/docs/providers/postgresql/r/replication_slot postgresql_replication_slot}.
type ReplicationSlot interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Plugin() *string
	SetPlugin(val *string)
	PluginInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
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
	ResetDatabase()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ReplicationSlot
type jsiiProxy_ReplicationSlot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ReplicationSlot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) Plugin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plugin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) PluginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationSlot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/replication_slot postgresql_replication_slot} Resource.
func NewReplicationSlot(scope constructs.Construct, id *string, config *ReplicationSlotConfig) ReplicationSlot {
	_init_.Initialize()

	j := jsiiProxy_ReplicationSlot{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.ReplicationSlot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/replication_slot postgresql_replication_slot} Resource.
func NewReplicationSlot_Override(r ReplicationSlot, scope constructs.Construct, id *string, config *ReplicationSlotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.ReplicationSlot",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_ReplicationSlot) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ReplicationSlot) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_ReplicationSlot) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ReplicationSlot) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ReplicationSlot) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ReplicationSlot) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ReplicationSlot) SetPlugin(val *string) {
	_jsii_.Set(
		j,
		"plugin",
		val,
	)
}

func (j *jsiiProxy_ReplicationSlot) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
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
func ReplicationSlot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.ReplicationSlot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ReplicationSlot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.ReplicationSlot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_ReplicationSlot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_ReplicationSlot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_ReplicationSlot) ResetDatabase() {
	_jsii_.InvokeVoid(
		r,
		"resetDatabase",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationSlot) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationSlot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationSlot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationSlot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ReplicationSlotConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/replication_slot#name ReplicationSlot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Sets the output plugin to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/replication_slot#plugin ReplicationSlot#plugin}
	Plugin *string `field:"required" json:"plugin" yaml:"plugin"`
	// Sets the database to add the replication slot to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/replication_slot#database ReplicationSlot#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/replication_slot#id ReplicationSlot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

// Represents a {@link https://www.terraform.io/docs/providers/postgresql/r/role postgresql_role}.
type Role interface {
	cdktf.TerraformResource
	BypassRowLevelSecurity() interface{}
	SetBypassRowLevelSecurity(val interface{})
	BypassRowLevelSecurityInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConnectionLimit() *float64
	SetConnectionLimit(val *float64)
	ConnectionLimitInput() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreateDatabase() interface{}
	SetCreateDatabase(val interface{})
	CreateDatabaseInput() interface{}
	CreateRole() interface{}
	SetCreateRole(val interface{})
	CreateRoleInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Encrypted() *string
	SetEncrypted(val *string)
	EncryptedInput() *string
	EncryptedPassword() interface{}
	SetEncryptedPassword(val interface{})
	EncryptedPasswordInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IdleInTransactionSessionTimeout() *float64
	SetIdleInTransactionSessionTimeout(val *float64)
	IdleInTransactionSessionTimeoutInput() *float64
	Inherit() interface{}
	SetInherit(val interface{})
	InheritInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Login() interface{}
	SetLogin(val interface{})
	LoginInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Replication() interface{}
	SetReplication(val interface{})
	ReplicationInput() interface{}
	Roles() *[]*string
	SetRoles(val *[]*string)
	RolesInput() *[]*string
	SearchPath() *[]*string
	SetSearchPath(val *[]*string)
	SearchPathInput() *[]*string
	SkipDropRole() interface{}
	SetSkipDropRole(val interface{})
	SkipDropRoleInput() interface{}
	SkipReassignOwned() interface{}
	SetSkipReassignOwned(val interface{})
	SkipReassignOwnedInput() interface{}
	StatementTimeout() *float64
	SetStatementTimeout(val *float64)
	StatementTimeoutInput() *float64
	Superuser() interface{}
	SetSuperuser(val interface{})
	SuperuserInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ValidUntil() *string
	SetValidUntil(val *string)
	ValidUntilInput() *string
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
	ResetBypassRowLevelSecurity()
	ResetConnectionLimit()
	ResetCreateDatabase()
	ResetCreateRole()
	ResetEncrypted()
	ResetEncryptedPassword()
	ResetId()
	ResetIdleInTransactionSessionTimeout()
	ResetInherit()
	ResetLogin()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetReplication()
	ResetRoles()
	ResetSearchPath()
	ResetSkipDropRole()
	ResetSkipReassignOwned()
	ResetStatementTimeout()
	ResetSuperuser()
	ResetValidUntil()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Role
type jsiiProxy_Role struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Role) BypassRowLevelSecurity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassRowLevelSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) BypassRowLevelSecurityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassRowLevelSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ConnectionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ConnectionLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) CreateDatabase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) CreateDatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) CreateRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) CreateRoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Encrypted() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) EncryptedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) EncryptedPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) EncryptedPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) IdleInTransactionSessionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleInTransactionSessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) IdleInTransactionSessionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleInTransactionSessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Inherit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inherit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) InheritInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inheritInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Login() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) LoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Replication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ReplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Roles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) RolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) SearchPath() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) SearchPathInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) SkipDropRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDropRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) SkipDropRoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDropRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) SkipReassignOwned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipReassignOwned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) SkipReassignOwnedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipReassignOwnedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) StatementTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) StatementTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) Superuser() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"superuser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) SuperuserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"superuserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Role) ValidUntilInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntilInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/role postgresql_role} Resource.
func NewRole(scope constructs.Construct, id *string, config *RoleConfig) Role {
	_init_.Initialize()

	j := jsiiProxy_Role{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.Role",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/role postgresql_role} Resource.
func NewRole_Override(r Role, scope constructs.Construct, id *string, config *RoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.Role",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Role) SetBypassRowLevelSecurity(val interface{}) {
	_jsii_.Set(
		j,
		"bypassRowLevelSecurity",
		val,
	)
}

func (j *jsiiProxy_Role) SetConnectionLimit(val *float64) {
	_jsii_.Set(
		j,
		"connectionLimit",
		val,
	)
}

func (j *jsiiProxy_Role) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Role) SetCreateDatabase(val interface{}) {
	_jsii_.Set(
		j,
		"createDatabase",
		val,
	)
}

func (j *jsiiProxy_Role) SetCreateRole(val interface{}) {
	_jsii_.Set(
		j,
		"createRole",
		val,
	)
}

func (j *jsiiProxy_Role) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Role) SetEncrypted(val *string) {
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_Role) SetEncryptedPassword(val interface{}) {
	_jsii_.Set(
		j,
		"encryptedPassword",
		val,
	)
}

func (j *jsiiProxy_Role) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Role) SetIdleInTransactionSessionTimeout(val *float64) {
	_jsii_.Set(
		j,
		"idleInTransactionSessionTimeout",
		val,
	)
}

func (j *jsiiProxy_Role) SetInherit(val interface{}) {
	_jsii_.Set(
		j,
		"inherit",
		val,
	)
}

func (j *jsiiProxy_Role) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Role) SetLogin(val interface{}) {
	_jsii_.Set(
		j,
		"login",
		val,
	)
}

func (j *jsiiProxy_Role) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Role) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_Role) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Role) SetReplication(val interface{}) {
	_jsii_.Set(
		j,
		"replication",
		val,
	)
}

func (j *jsiiProxy_Role) SetRoles(val *[]*string) {
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

func (j *jsiiProxy_Role) SetSearchPath(val *[]*string) {
	_jsii_.Set(
		j,
		"searchPath",
		val,
	)
}

func (j *jsiiProxy_Role) SetSkipDropRole(val interface{}) {
	_jsii_.Set(
		j,
		"skipDropRole",
		val,
	)
}

func (j *jsiiProxy_Role) SetSkipReassignOwned(val interface{}) {
	_jsii_.Set(
		j,
		"skipReassignOwned",
		val,
	)
}

func (j *jsiiProxy_Role) SetStatementTimeout(val *float64) {
	_jsii_.Set(
		j,
		"statementTimeout",
		val,
	)
}

func (j *jsiiProxy_Role) SetSuperuser(val interface{}) {
	_jsii_.Set(
		j,
		"superuser",
		val,
	)
}

func (j *jsiiProxy_Role) SetValidUntil(val *string) {
	_jsii_.Set(
		j,
		"validUntil",
		val,
	)
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
func Role_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.Role",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Role_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.Role",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Role) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Role) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Role) ResetBypassRowLevelSecurity() {
	_jsii_.InvokeVoid(
		r,
		"resetBypassRowLevelSecurity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetConnectionLimit() {
	_jsii_.InvokeVoid(
		r,
		"resetConnectionLimit",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetCreateDatabase() {
	_jsii_.InvokeVoid(
		r,
		"resetCreateDatabase",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetCreateRole() {
	_jsii_.InvokeVoid(
		r,
		"resetCreateRole",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetEncrypted() {
	_jsii_.InvokeVoid(
		r,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetEncryptedPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetEncryptedPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetIdleInTransactionSessionTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetIdleInTransactionSessionTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetInherit() {
	_jsii_.InvokeVoid(
		r,
		"resetInherit",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetLogin() {
	_jsii_.InvokeVoid(
		r,
		"resetLogin",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetReplication() {
	_jsii_.InvokeVoid(
		r,
		"resetReplication",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetRoles() {
	_jsii_.InvokeVoid(
		r,
		"resetRoles",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetSearchPath() {
	_jsii_.InvokeVoid(
		r,
		"resetSearchPath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetSkipDropRole() {
	_jsii_.InvokeVoid(
		r,
		"resetSkipDropRole",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetSkipReassignOwned() {
	_jsii_.InvokeVoid(
		r,
		"resetSkipReassignOwned",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetStatementTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetStatementTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetSuperuser() {
	_jsii_.InvokeVoid(
		r,
		"resetSuperuser",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) ResetValidUntil() {
	_jsii_.InvokeVoid(
		r,
		"resetValidUntil",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Role) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Role) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type RoleConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// The name of the role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#name Role#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Determine whether a role bypasses every row-level security (RLS) policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#bypass_row_level_security Role#bypass_row_level_security}
	BypassRowLevelSecurity interface{} `field:"optional" json:"bypassRowLevelSecurity" yaml:"bypassRowLevelSecurity"`
	// How many concurrent connections can be made with this role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#connection_limit Role#connection_limit}
	ConnectionLimit *float64 `field:"optional" json:"connectionLimit" yaml:"connectionLimit"`
	// Define a role's ability to create databases.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#create_database Role#create_database}
	CreateDatabase interface{} `field:"optional" json:"createDatabase" yaml:"createDatabase"`
	// Determine whether this role will be permitted to create new roles.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#create_role Role#create_role}
	CreateRole interface{} `field:"optional" json:"createRole" yaml:"createRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#encrypted Role#encrypted}.
	Encrypted *string `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Control whether the password is stored encrypted in the system catalogs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#encrypted_password Role#encrypted_password}
	EncryptedPassword interface{} `field:"optional" json:"encryptedPassword" yaml:"encryptedPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#id Role#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Terminate any session with an open transaction that has been idle for longer than the specified duration in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#idle_in_transaction_session_timeout Role#idle_in_transaction_session_timeout}
	IdleInTransactionSessionTimeout *float64 `field:"optional" json:"idleInTransactionSessionTimeout" yaml:"idleInTransactionSessionTimeout"`
	// Determine whether a role "inherits" the privileges of roles it is a member of.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#inherit Role#inherit}
	Inherit interface{} `field:"optional" json:"inherit" yaml:"inherit"`
	// Determine whether a role is allowed to log in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#login Role#login}
	Login interface{} `field:"optional" json:"login" yaml:"login"`
	// Sets the role's password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#password Role#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Determine whether a role is allowed to initiate streaming replication or put the system in and out of backup mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#replication Role#replication}
	Replication interface{} `field:"optional" json:"replication" yaml:"replication"`
	// Role(s) to grant to this new role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#roles Role#roles}
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// Sets the role's search path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#search_path Role#search_path}
	SearchPath *[]*string `field:"optional" json:"searchPath" yaml:"searchPath"`
	// Skip actually running the DROP ROLE command when removing a ROLE from PostgreSQL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#skip_drop_role Role#skip_drop_role}
	SkipDropRole interface{} `field:"optional" json:"skipDropRole" yaml:"skipDropRole"`
	// Skip actually running the REASSIGN OWNED command when removing a role from PostgreSQL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#skip_reassign_owned Role#skip_reassign_owned}
	SkipReassignOwned interface{} `field:"optional" json:"skipReassignOwned" yaml:"skipReassignOwned"`
	// Abort any statement that takes more than the specified number of milliseconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#statement_timeout Role#statement_timeout}
	StatementTimeout *float64 `field:"optional" json:"statementTimeout" yaml:"statementTimeout"`
	// Determine whether the new role is a "superuser".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#superuser Role#superuser}
	Superuser interface{} `field:"optional" json:"superuser" yaml:"superuser"`
	// Sets a date and time after which the role's password is no longer valid.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/role#valid_until Role#valid_until}
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}

// Represents a {@link https://www.terraform.io/docs/providers/postgresql/r/schema postgresql_schema}.
type Schema interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DropCascade() interface{}
	SetDropCascade(val interface{})
	DropCascadeInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IfNotExists() interface{}
	SetIfNotExists(val interface{})
	IfNotExistsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	Policy() SchemaPolicyList
	PolicyInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
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
	PutPolicy(value interface{})
	ResetDatabase()
	ResetDropCascade()
	ResetId()
	ResetIfNotExists()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwner()
	ResetPolicy()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Schema
type jsiiProxy_Schema struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Schema) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) DropCascade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropCascade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) DropCascadeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropCascadeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) IfNotExists() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ifNotExists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) IfNotExistsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ifNotExistsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) Policy() SchemaPolicyList {
	var returns SchemaPolicyList
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) PolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schema) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/schema postgresql_schema} Resource.
func NewSchema(scope constructs.Construct, id *string, config *SchemaConfig) Schema {
	_init_.Initialize()

	j := jsiiProxy_Schema{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.Schema",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/postgresql/r/schema postgresql_schema} Resource.
func NewSchema_Override(s Schema, scope constructs.Construct, id *string, config *SchemaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.Schema",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_Schema) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Schema) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_Schema) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Schema) SetDropCascade(val interface{}) {
	_jsii_.Set(
		j,
		"dropCascade",
		val,
	)
}

func (j *jsiiProxy_Schema) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Schema) SetIfNotExists(val interface{}) {
	_jsii_.Set(
		j,
		"ifNotExists",
		val,
	)
}

func (j *jsiiProxy_Schema) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Schema) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Schema) SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_Schema) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
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
func Schema_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-postgresql.Schema",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Schema_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-postgresql.Schema",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Schema) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_Schema) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_Schema) PutPolicy(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Schema) ResetDatabase() {
	_jsii_.InvokeVoid(
		s,
		"resetDatabase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Schema) ResetDropCascade() {
	_jsii_.InvokeVoid(
		s,
		"resetDropCascade",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Schema) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Schema) ResetIfNotExists() {
	_jsii_.InvokeVoid(
		s,
		"resetIfNotExists",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Schema) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Schema) ResetOwner() {
	_jsii_.InvokeVoid(
		s,
		"resetOwner",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Schema) ResetPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Schema) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Schema) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SchemaConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// The name of the schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#name Schema#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The database name to alter schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#database Schema#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// When true, will also drop all the objects that are contained in the schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#drop_cascade Schema#drop_cascade}
	DropCascade interface{} `field:"optional" json:"dropCascade" yaml:"dropCascade"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#id Schema#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When true, use the existing schema if it exists.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#if_not_exists Schema#if_not_exists}
	IfNotExists interface{} `field:"optional" json:"ifNotExists" yaml:"ifNotExists"`
	// The ROLE name who owns the schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#owner Schema#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#policy Schema#policy}
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
}

type SchemaPolicy struct {
	// If true, allow the specified ROLEs to CREATE new objects within the schema(s).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#create Schema#create}
	Create interface{} `field:"optional" json:"create" yaml:"create"`
	// If true, allow the specified ROLEs to CREATE new objects within the schema(s) and GRANT the same CREATE privilege to different ROLEs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#create_with_grant Schema#create_with_grant}
	CreateWithGrant interface{} `field:"optional" json:"createWithGrant" yaml:"createWithGrant"`
	// ROLE who will receive this policy (default: PUBLIC).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#role Schema#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// If true, allow the specified ROLEs to use objects within the schema(s).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#usage Schema#usage}
	Usage interface{} `field:"optional" json:"usage" yaml:"usage"`
	// If true, allow the specified ROLEs to use objects within the schema(s) and GRANT the same USAGE privilege to different ROLEs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#usage_with_grant Schema#usage_with_grant}
	UsageWithGrant interface{} `field:"optional" json:"usageWithGrant" yaml:"usageWithGrant"`
}

type SchemaPolicyList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) SchemaPolicyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SchemaPolicyList
type jsiiProxy_SchemaPolicyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SchemaPolicyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSchemaPolicyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SchemaPolicyList {
	_init_.Initialize()

	j := jsiiProxy_SchemaPolicyList{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.SchemaPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSchemaPolicyList_Override(s SchemaPolicyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.SchemaPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SchemaPolicyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SchemaPolicyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyList) Get(index *float64) SchemaPolicyOutputReference {
	var returns SchemaPolicyOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SchemaPolicyOutputReference interface {
	cdktf.ComplexObject
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
	Create() interface{}
	SetCreate(val interface{})
	CreateInput() interface{}
	CreateWithGrant() interface{}
	SetCreateWithGrant(val interface{})
	CreateWithGrantInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Usage() interface{}
	SetUsage(val interface{})
	UsageInput() interface{}
	UsageWithGrant() interface{}
	SetUsageWithGrant(val interface{})
	UsageWithGrantInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetCreateWithGrant()
	ResetRole()
	ResetUsage()
	ResetUsageWithGrant()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SchemaPolicyOutputReference
type jsiiProxy_SchemaPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SchemaPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) Create() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) CreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) CreateWithGrant() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createWithGrant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) CreateWithGrantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createWithGrantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) Usage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) UsageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) UsageWithGrant() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usageWithGrant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaPolicyOutputReference) UsageWithGrantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usageWithGrantInput",
		&returns,
	)
	return returns
}


func NewSchemaPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SchemaPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SchemaPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-postgresql.SchemaPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSchemaPolicyOutputReference_Override(s SchemaPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-postgresql.SchemaPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SchemaPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyOutputReference) SetCreate(val interface{}) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyOutputReference) SetCreateWithGrant(val interface{}) {
	_jsii_.Set(
		j,
		"createWithGrant",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyOutputReference) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyOutputReference) SetUsage(val interface{}) {
	_jsii_.Set(
		j,
		"usage",
		val,
	)
}

func (j *jsiiProxy_SchemaPolicyOutputReference) SetUsageWithGrant(val interface{}) {
	_jsii_.Set(
		j,
		"usageWithGrant",
		val,
	)
}

func (s *jsiiProxy_SchemaPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchemaPolicyOutputReference) ResetCreateWithGrant() {
	_jsii_.InvokeVoid(
		s,
		"resetCreateWithGrant",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchemaPolicyOutputReference) ResetRole() {
	_jsii_.InvokeVoid(
		s,
		"resetRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchemaPolicyOutputReference) ResetUsage() {
	_jsii_.InvokeVoid(
		s,
		"resetUsage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchemaPolicyOutputReference) ResetUsageWithGrant() {
	_jsii_.InvokeVoid(
		s,
		"resetUsageWithGrant",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchemaPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

