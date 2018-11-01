package algebra

// Variable is just an index and therefore depends on Variables
type Variable int

// NotAVariable is the value for something that aggregates a Variable optionally
const NotAVariable = Variable(-1)