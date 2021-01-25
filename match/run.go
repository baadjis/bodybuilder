package match
import(
	"strings"
	"fmt"
)
func TestSearch(q *QueryMatch){
	var builder strings.Builder
	query:=&Query{Type:q}
	highlight := NewHighlight("_").WithTags("styled")
    highlight1:= NewHighlight("kfs","kfs2").WithNumberOfFragment(5,3)
	highlighter:=NewHighlightBuilder(false,highlight,highlight1)

	search:=NewSearch(query,highlighter)
	search.Build(&builder)
	search_string :=strings.NewReader(builder.String())
	fmt.Println(search_string)
}
func TestMultiMatchs(){

	multimatchs:= MultiMatchs("search","search1","search2")
    TestSearch(multimatchs)
}
func TestMatch(){

	match :=Match("location",2)
	match.WithTieBreaker(0.3)
	match.WithFuzzyTranspositions()
	TestSearch(match)
}

func Run(){
	
	TestMultiMatchs()
	//TestMatch()
	}