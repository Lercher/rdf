# RDF tools

This is an exploration project on RAM based virtualized graph databases written in Go (golang). I doubt that it will ever reach production quality.

## Get and Build the Code

Assuming the go SDK is installed already, download

````bat
go get github.com/lercher/rdf
````

### Build

````bat
go get
go build
````

### Test

````bat
go test ./...
````

or

````bat
go test ./... -short
````

## Deployment and Runtime Requirements

There are no particular runtime requirements, just copy the executable and run it with
appropriate command line parameters. Go supports a lot of environments out of the box
and so does this project.

## Development Tasks

### Modifing *.g4 Grammars

First of all: don't do it. The generated artefacts of the antlr4 run are included in the repository,
ready to be compiled by go. But in case you really need to do it:

* install a current JRE
* install and use Visual Studio Code
* install this [antlr4 VSCode extension](https://marketplace.visualstudio.com/items?itemName=mike-lischke.vscode-antlr4)

Then use these VSCode settings as provided in [settings.json](.vscode/settings.json) in this repository to have the lexer, parser and listeners generated by the extension:

````json
"antlr4.generation": {
    "mode": "external",
    "language": "Go",
    "listeners": true,
    "visitors": false,
    "outputDir": "parser"
},
````

## Packages

It consists of the following packages.

### package algebra

Query algebra resulting from a parsed SPARQL statement

### package encoding/binary

Binary reading and writing of Graphs

### package encoding/csv

Import of a csv reader with a header line and subject in the first column to a Graph

### package graph

Basic rdf graph datatypes such as Graph, Triple, TriplePattern and Value

### package processor

Engine to execute an algebra instance on a graph

### package sparql

Parsing SPARQL syntax to form an algebra instance

### package values

Primitive datatypes to be used in a graph's assertions

## Using the Packages

To execute a query against a data store, you'll need a `Graph` containing the asserted `Triples`, an `Algebra` representing the SPARQL query and a `processor` method.

### How to Get a Graph

Either create a new one and `Assert` your knowledge

````go
import "github.com/lercher/rdf/graph"

a := graph.IRI(`http://www.w3.org/1999/02/22-rdf-syntax-ns#type`)
g := graph.New()
g.Assert("martin", a, "person")
````

or load one from a csv file with a header line

````go
import "github.com/lercher/rdf/encoding/csv"

const nsEst    = `http://education.data.gov.uk/def/school/establishment/`
const nsSchool = `http://education.data.gov.uk/def/school/`

f, err := os.Open(`path_to.csv`)
defer f.Close()

dec := csv.NewDecoder(f, nsEst, nsSchool)
g, err := dec.Decode()
````

or `import "github.com/lercher/rdf/encoding/binary"` and store and load a binary graph represention.

### How to get an Algebra

Allthough it's possible to construct an algebra from scratch, it's far more convenient
to parse an algebra from a SPARQL query

````go
import "github.com/antlr4-go/antlr/v4"
import "github.com/lercher/rdf/sparql"

input := antlr.NewInputStream(`select * {?s ?p ?o}`)
ast, err := sparql.Parse(input)
a := a.Algebra()
````

Don't forget to optimize your algebra after parsing. Currently, however there is no actual optimization.

````go
a = a.Optimize()
````

**Note:** The `ast` variable holds the Abstract Syntax Tree of the parsed query. It might
contain information more nearer to the parsed query than the algebra created from the AST.

### How to Execute the Algebra on a Graph

There is currently only one processor, that processes the joins and filters of an algebra on a graph.
It is started by calling `processor.Execute/3` on an `Algebra` and a `Graph`. The `processor.Receiver`
function in the 3rd parameter is called for any result line of the execution, here a `func` literal:

````go
import "github.com/lercher/rdf/sparql/processor"

// a algebra, g graph
err := processor.Execute(a, g, func(bs algebra.Binding, vs *algebra.Variables) (bool, error) {
    m := bs.Materialize(g, vs)
    log.Print(m)
    return true, nil
})
````

The `bs` var holds a compact representation of a result line's variables. It has to be materialized
with the help of the graph and the selected variables described in `vs`. The resulting `m` is just
a slice of Variablename (`string`) and Value (`interface{}`) pairs.

The `bool` return value expresses the receiver function's wish to continue processing (`true`)
or stop it (`false`). An `error` is passed up to the `processor.Execute` call, ends it
immediatley and is the return value of it.

**Note:** Result lines are raised unordered and the sequence will be different on each call, because
the underlying Go maps behave this way. On the other hand, the query processing only takes the
memory it needs to produce the current result line.

**Warning:** Issuing ordered queries implies complete processing and materialization of all result lines
before the first call of the receiver function. This can cost a lot of ressources, if the
result set is large.

## Ressources

* [SPARQL By Example - A Tutorial](https://www.w3.org/2009/Talks/0615-qbe/)
* [Gossary](https://www.futurelearn.com/courses/linked-data/0/steps/16092)
