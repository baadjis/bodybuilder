package match
import(
	_"strconv"

)

type SearchElementInterface interface{
	Build(Builder)
	Name()string
	
}
type Search  struct{
	Elements []SearchElementInterface
}

func (s *Search)Build(builder Builder){
	builder.WriteString("{\n")
	for e,v := range s.Elements{
		WriteQuoted(builder,v.Name())
		builder.WriteString(":{\n")
		v.Build(builder)
		builder.WriteString("\n}")
		if e!=len(s.Elements)-1{
			builder.WriteString(",\n")
		}
	  
		
	}
	builder.WriteString("\n}")
}
func NewSearch(elms ...SearchElementInterface)*Search{
	return &Search{Elements:elms}
}
type MapBuilder func(MapElement, Builder)

type Map struct{
  Query Query 
  Size int
  From int
  Fields [] string
}

func (m*Map)Build(builder Builder){
	if len(m.Fields)>0{
		builder.WriteString("fields: ")
		WriteArray(builder,m.Fields)
	}

	m.Query.Build(builder)

}


type MapElement struct{
	Name string
	Elemnts interface{}
	Builder MapBuilder
}

type Builder interface{
	
	WriteString(string) (int, error)
	
}
/*type Builder interface {
  *Writer
  
}*/

type Query struct {
	Type QueryInterface

}

func (q* Query)Name()string{
  return "query"
}

func(q*Query)Build(builder Builder){

	q.Type.Build(builder)

}

type QueryInterface interface{
	
	Build(Builder)
}


//---------------------- Match type Query implementation --------------------

type QueryMatch struct{
    Name string
	Fields []string
	Query interface{}
	Options map[string]interface{}
}
/*func (m *QueryMatch)Name()string{
	return "match"
}*/

//Build a match type query body
func(m QueryMatch) Build(builder Builder){
	WriteQuoted(builder,m.Name)
	
	builder.WriteString(":{\n")

	if len(m.Options)==0 {

		WriteWithFormat(builder,m.Fields[0],m.Query)
	}else{
		m.Options["query"]=m.Query
		if len(m.Fields)>1{
			m.Options["fields"]= m.Fields
			
			Encode(builder,m.Options)
		}else{
			WriteQuoted(builder,m.Fields[0])
			builder.WriteString(":")
			Encode(builder,m.Options)
		}

	}
	builder.WriteString("\n}")
	
    
	
}
//NewQueryMatch Match type query factory
func NewQueryMatch(_name string, field string,query interface{})*QueryMatch{
	return &QueryMatch{
		Name:_name,
		Fields:[]string{field},
		Query:query,
	    Options:make(map[string]interface{})}
}
// Match create match query
func Match(field string,query interface{}) *QueryMatch{
	return NewQueryMatch("match",field,query)
}
// Match_Phrase_Prefix create  match_phrase_prefix query
func Match_Phrase_Prefix(field string,query interface{})*QueryMatch{
	return NewQueryMatch("match_phrase_prefix",field,query)
}

// Match_Phrase create  match_phrase query
func Match_Phrase(field string,query interface{})*QueryMatch{
	return NewQueryMatch("match_phrase",field,query)
}

// Match_Boolean_Prefix create  match_bool_prefix query
func Match_Boolean_Prefix(field string,query interface{})*QueryMatch{
	return NewQueryMatch("match_bool_prefix",field,query)
}



