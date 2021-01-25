package match

import(
	"fmt"
)




// ---------------------------------- Multimatch queries implementation --------------------

/*type QueryMultiMatchs struct{
	Query string
	Type string
	Fields [] string
	Operator string
}*/
/*func (mm *QueryMultiMatchs)Name()string{
	return"multi_match"
}*/

//WithOperator add operator 
//default operator is OR so i will put "AND" if required
func(m *QueryMatch)WithOperator(){
	m.Options["oprator"]="AND"
}

//WithAnalyzer add analyzer options
func(m *QueryMatch)WithAnalyzer(an string){
	m.Options["analyzer"]=an
}

//WithTieBreaker add tie_breaker
func(m *QueryMatch)WithTieBreaker(tb float32){
	m.Options["tie_breaker"]=tb
}
//WithLenient add lenient options
// Default is false so i will set it to true if required
func (m *QueryMatch)WithLenient(){
	m.Options["lenient"]=true
}

func (m *QueryMatch)WithMinimumShouldMatch(min string){
	m.Options["minimum_should_match"]=min
}

//WithZerosTermQuery add zero_terms_query option
//Default is None so i will put "all" if required
func (m *QueryMatch)WithZeroTermsQuery(){
	m.Options["zero_terms_query"]="all"
}

//WithFuzzyRewrite add fuzzy_rewrite option
func(m *QueryMatch) WithFuzzyRewrite(fuzz string){

	Allowed:=`
	----------------------------
	|    "constant_score"       |
	----------------------------
	|   "constant_score_boolean"|
	-----------------------------
	|    "scoring_boolean"      |
	-----------------------------
	|"top_terms_blended_freqs_N"|
	-----------------------------
	|      top_terms_boost_N"   |
	-----------------------------
	|       "top_terms_N"       |
	-----------------------------
 `
  AllowedParameters:=[]string{
	"constant_score",
	"constant_score_boolean",
	"scoring_boolean",
	"top_terms_blended_freqs_N",
	"top_terms_boost_N",
	"top_terms_N"}
  if Contains(AllowedParameters,fuzz){
	  m.Options["fuzzy_rewrite"]=fuzz
  }else{
	  fmt.Println("allowed parametters are:")
	  fmt.Println(Allowed)
	  panic("InvalidParameter for fuzzy_rewrite")
  }
}

//WithFuzzyTranspositions add fuzzy_transpositions option
//default parameter is true so i will set it to false if requiredx
func(m *QueryMatch) WithFuzzyTranspositions(){
	m.Options["fuzzy_transpositions"]=false

}

//WithPrefixLength add prefix_length option
func(m *QueryMatch) WithPrefixLength(_length int){
	m.Options["prefix_length"]=_length

}

//WithMaxExpansions add max_expansions option
func(m *QueryMatch) WithMaxExpansions(_max int){
	m.Options["max_expansions"]=_max

}

//WithFuzziness add fuzziness option
func(m *QueryMatch) WithMaxFuzziness(fuzz string){
	m.Options["fuzziness"]=fuzz

}

//WithAutoSynonyms add auto_generate_synonyms_phrase_query option
//default parameter is true so i will set it to false if required
func(m *QueryMatch) WithAutoSynonyms(){
	m.Options["auto_generate_synonyms_phrase_query"]=false

}



//NewMultiMatchs create a multi match query
func NewMultiMatchs(query string,_type string ,fields...string)*QueryMatch{
	fmt.Println(fields)

	
	 return &QueryMatch{
		Name:"multi_matchs",
	   Query:query,
	   Fields:fields,
	   Options:map[string]interface{}{"type":_type}}

}
//CrossFieldsMultiMatchs create a Multimatch query of type cross_fields
func CrossFieldsMultiMatchs(query string ,fields...string)*QueryMatch{
	return NewMultiMatchs(query,"cross_fields",fields...)
}

//PhraseMultiMatchs create a Multimatch query of type phrase
func PhraseMultiMatchs(query string ,fields[]string)*QueryMatch{
	return NewMultiMatchs(query,"phrase",fields...)

}
//PhrasePrefixMultiMatchs create a Multimatch query of type phrase_prefix
func PhrasePrefixMultiMatchs(query string ,fields[]string)*QueryMatch{
	return NewMultiMatchs(query,"phrase_prefix",fields...)

}
//BoolPrefixMultiMatchs create a Multimatchs query of type bool_prefix
func BoolPrefixMultiMatchs(query string ,fields[]string)*QueryMatch{
	return NewMultiMatchs(query,"bool_prefix",fields...)

}

//MostFieldsMultiMatchs create a Multimatchs query of type most_fields
func MostFieldsMultiMatchs(query string ,fields[]string) *QueryMatch{
	return NewMultiMatchs(query,"most_fields",fields...)
}

//MultiMatchs create a Multimatchs query of type best_fields
func MultiMatchs(query string,fields...string) *QueryMatch{
	return NewMultiMatchs(query,"best_fields",fields...)
}
//Build build MultiMatch query body
/*func(mm *QueryMultiMatchs)Build(builder Builder){
	WriteQuoted(builder,"multi_match")
	builder.WriteString(":{\n")
	WriteWithFormat(builder,"query",mm.Query)
	builder.WriteString(",")
	builder.WriteString("\n")
	WriteWithFormat(builder,"type",mm.Type)
	if len(mm.Fields)>0{
		builder.WriteString(",")
		builder.WriteString("\n")
		WriteQuoted(builder,"fields")
		builder.WriteString(":")
		WriteArray(builder,mm.Fields)
	}

	if mm.Analyzer!=""{
		builder.WriteString(",")
		builder.WriteString("\n")
		WriteWithFormat(builder,"analyzer",mm.Analyzer)
	}
	if mm.Tie_Breaker!=0{

		builder.WriteString(",")
		builder.WriteString("\n")
		WriteWithFormat(builder,"tie_breaker",mm.Tie_Breaker)
		
	}
	builder.WriteString("\n}")


}*/
