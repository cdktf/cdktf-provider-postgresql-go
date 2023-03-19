package datapostgresqlschemas

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPostgresqlSchemasConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The PostgreSQL database which will be queried for schema names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/d/schemas#database DataPostgresqlSchemas#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/d/schemas#id DataPostgresqlSchemas#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Determines whether to include system schemas (pg_ prefix and information_schema). 'public' will always be included.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/d/schemas#include_system_schemas DataPostgresqlSchemas#include_system_schemas}
	IncludeSystemSchemas interface{} `field:"optional" json:"includeSystemSchemas" yaml:"includeSystemSchemas"`
	// Expression(s) which will be pattern matched in the query using the PostgreSQL LIKE ALL operator.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/d/schemas#like_all_patterns DataPostgresqlSchemas#like_all_patterns}
	LikeAllPatterns *[]*string `field:"optional" json:"likeAllPatterns" yaml:"likeAllPatterns"`
	// Expression(s) which will be pattern matched in the query using the PostgreSQL LIKE ANY operator.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/d/schemas#like_any_patterns DataPostgresqlSchemas#like_any_patterns}
	LikeAnyPatterns *[]*string `field:"optional" json:"likeAnyPatterns" yaml:"likeAnyPatterns"`
	// Expression(s) which will be pattern matched in the query using the PostgreSQL NOT LIKE ALL operator.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/d/schemas#not_like_all_patterns DataPostgresqlSchemas#not_like_all_patterns}
	NotLikeAllPatterns *[]*string `field:"optional" json:"notLikeAllPatterns" yaml:"notLikeAllPatterns"`
	// Expression which will be pattern matched in the query using the PostgreSQL ~ (regular expression match) operator.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/d/schemas#regex_pattern DataPostgresqlSchemas#regex_pattern}
	RegexPattern *string `field:"optional" json:"regexPattern" yaml:"regexPattern"`
}

