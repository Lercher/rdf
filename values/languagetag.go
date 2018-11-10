package values

// LanguageTag is Language spec attached to a Literal
// it is of the form '@' PN_CHARS_BASE+ ('-' (PN_CHARS_BASE DIGIT)+)*
type LanguageTag string

// NotALanguage represents no language tag given in a Literal
const NotALanguage = LanguageTag("")
