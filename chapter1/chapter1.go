/*
SUMMMARY NOTES:

MODULES
-------------------
- This is a collection of Go packages and their dependencies and how they can be built into each other:-
- To initialize a new module you use the go mod init ...path ...
- preferrably use the mod name that refers to where the module can be stored in a VCS like github and can be accessed later on :-
- for example go mod init github.com/Felixoh/demo

PACKAGES
-------------------
- >A package is a collection of go files that share the same import path
they are almost always in the same directory as each other , and the directory name becomes the package name for the code
. As an example below :
|---- go.mod
|---- smtp
|   |---- smtp.go
|   '---- template.go
'---- sound
|---- analyzer.go
'---- sound.go
- Code within a package can reference any identifier, such as a function, variable, or type,
within the package,regardless of whether that identifier is exported. Any exported
identifier can be used by any other package.

NAMING PACKAGES
---------------------
- >The name of the package must be declared at the top of every Go source file using the
package keyword.
- >There are naming rules, and conventions, for naming packages in Go. Package names
are lowercased and should contain only contain letters and numbers. You are not allowed
to start package names with a number
- >The below are examples of proper package naming :-
// Good
*package http
*package sql
*package oauth2
*package base64
-> First, your package name should be short and descriptive. For example, sql is a good
package name, but structured-query-language is not
-> To help prevent these clashes, you should always try to avoid using thenames of popular packages already in the Go standard library
or names used for popular variables, such as user, conn, db, etc.

FILES,FOLDERS AND ORGANIZATION
-------------------------------------
-> Below is a simple package organization structure commonly used

|---- go.mod
|---- smtp
|   |---- smtp.go
Folders, Files, and Organization 7
|   '---- template.go
'---- sound
|---- analyzer.go
'---- sound.go
-> With rare exceptions, a folder can only contain one package
-> All files within the same directory/package have to have the same package name in
the file
-> File names are less strict than package names. The convention is to use lowercase letters
and underscores, as shown
|---- bad
|   |---- ServiceManagerTest.go
|   |---- USER_test.go
|   |---- User.go
|   '---- serviceManager.go
|---- go.mod
'---- good
|---- service_manager.go
|---- service_manager_test.go
|---- user.go
'---- user_test.go

-> Packages should be singularly focused. For example, if you have a package that contains
tools for working with time and for working with HTTP, then that package should be split
into two packages: one for time and one for HTTP. Now the packages have a clear,
singular purpose.
-> Package files should be organized in a way that is clear and obvious to the reader. Avoid
using one or two large files for a package. Instead, break the package into smaller files,
with each having a clear purpose.
-> An example of a good Folder and its source files
.
|---- order.go
|---- order_test.go
|---- product.go
|---- product_test.go
|---- user.go
'---- user_test.go

+
IMPORTING PACKAGES AND MODULES
------------------------------------------------------
------------------------------------------------------
-> The import path for a package is the relative path from the module root to the package
directory.
->The import path for a package is the relative path from the module root to the package
directory. The import path for a package, or module, always starts with a module name
-> Example :

	import "fmt" // import path

func Greet() {
fmt.Println("Hello, world!")
}

IMPORTS
----------------
-> The import keyword is used along with the fully qualified package name to import,
-> the following examples illustrate the import statement and paths and how they can become useful in the longwhile :-
mport "time"
import "os/exec"
import "github.com/user/repo/foo"
-> Multiple import statements can be condensed into a single import statement
import (
"time"
"os/exec"
"github.com/user/repo/foo"
)
-> Sometimes to avoid name collisions when doing package imports we can apply the alias keyword as seen below :-

package main
import (
"demo/bar"
pub "demo/foo/bar"
)
func main() {
var _ pub.Pub
var _ bar.Bar
}

ADDING DEPENDANCIES
--------------------
-> To add dependencies to our module, we can use the go get command,
-> dependencies are other programmes we need in our code to make it reusable that are not found in our standard library in that case
-> For example we can apply the go get command as follows :-
go get github.com/gobuffalo/flect
-> When adding a third-party dependency to the module, the Go toolchain will generate and manage a go.sum file.
->
|---- go.mod
|---- go.sum
'---- main.go

-> A typical directory outlook in go project
-> The go.sum file contains a list of all the dependencies, direct and indirect, that
are required by the module.
->Additionally, the go.sum file contains a hash of the source
code and go.mod file for each dependency
-> The go.sum file should be checked into version control. The go.sum acts both as a
security measure to prevent malicious code from being added to the module and as a
source of truth for the dependencies that are required by the module

UPDATING DEPENDENCIES
--------------------------------------------------
-> To update a direct dependency of our module, we can use the go get command,
Listing 1.42, with the -u flag.
->> go get -u github.com/gobuffalo/flect
-> Go modules are versioned using semantic versioning.6 For example, v1.2.3 is a semantic version that
tells us the major version of the release, 1,

	the minor version of the release, 2,

and the patch version of the release, 3.
-> Because of semantic import versioning, we can have multiple versions of a package in our
module.
-> When importing the two, we can assign import aliases to the two versions to allow us
to use them side by side.
-> For example in codebase we can use :-
import (
"fmt"
"log"
one "github.com/gofrs/uuid"
three "github.com/gofrs/uuid/v3"
*/
package chapter1
