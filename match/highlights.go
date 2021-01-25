package match
import (
	_"fmt"
)

type HighlightBuilder struct{
	MainHighlighter *Highlight
	Overiders []*Highlight
	Ordered bool
}
func(hb *HighlightBuilder)Name()string{
	return("highlight")
}

func NewHighlightBuilder(ordered bool,highlights...*Highlight)*HighlightBuilder{
         hi:=&HighlightBuilder{
		  Ordered :ordered}
		  if highlights[0].Fields[0]=="_"{
			  hi.MainHighlighter=highlights[0]
			  if len(highlights)>0{
				hi.Overiders=highlights[1:]
			  }
			 
		  }else{
			  hi.Overiders=highlights
		  }
		  return hi
}
func AppendMapInterface(m1 interface{},m2 map[string]interface{})map[string]interface{}{
	if m1==nil{
		return m2
	}else{
		merg:=m1.(map[string]interface{})
		for k,v := range m2{
			merg[k]=v
		}
		return merg
	}
	
}
func(hb *HighlightBuilder)Build(builder Builder){
	  if hb.MainHighlighter!=nil{
		hb.MainHighlighter.Build(builder)
	  }
	  
	  if len(hb.Overiders)>0{
		if hb.MainHighlighter!=nil{
            builder.WriteString(",\n")
		}
		
         ranger := func(){
			for k,v :=range hb.Overiders{
				v.Build(builder)
				if k!=len(hb.Overiders)-1{
					builder.WriteString("\n")
					builder.WriteString(",")
				}
			}
		 }

		 WriteQuoted(builder,"fields")
		 if hb.Ordered{
			 builder.WriteString(": [\n")
			 ranger() 
			 builder.WriteString("\n]\n")
		 }else{
		   builder.WriteString(": {\n")
		   ranger() 
		   builder.WriteString("\n}\n")  
		 }

	  }


}


type Highlight struct{
	Fields []string
	Options map[string] interface{}
}


func(hi *Highlight)Build(builder Builder){
    if hi.Fields[0]=="_"{
		keys:=Keys(hi.Options)
		for k,v :=range keys{
			WriteWithFormat(builder,v,hi.Options[v])
			
			if k !=len(keys)-1{
				builder.WriteString(",")
				builder.WriteString("\n")
			}
		}
	}else{
		for k,v:=range hi.Fields{
			WriteQuoted(builder,v)
			builder.WriteString(":")
			Encode(builder,hi.Options[v])
			if k!=len(hi.Fields)-1{
				builder.WriteString(",\n")
			}

			}

		}

	}


//
func NewHighlight(fields ...string)*Highlight{
    return &Highlight{
		Fields:fields,
		Options:make(map[string]interface{})}
}
func(hi *Highlight) Name() string{
	return "highlight"
}

func(hi *Highlight)AddFieldOption(field ,option_key string ,option_value interface{}){
	p :=map[string]interface{}{option_key:option_value}
	if hi.Options[field]==nil{
		hi.Options[field]=p
	}else{
		hi.Options[field]=AppendMapInterface(hi.Options[field],p)
	}	
}
func (hi *Highlight)AddGlobalOption(option_key string ,option_value interface{}){
	
	hi.Options[option_key]=option_value
	
}

func(hi *Highlight)AddOption(option_key string ,option_value interface{}){
	if hi.Fields[0]=="_"{
		hi.Options[option_key]=option_value
		}else{
           for _,v := range hi.Fields{
			   hi.AddFieldOption(v,option_key,option_value)
		   }
		}
}

func (hi*Highlight)WithNumberOfFragment(numbers ...int)*Highlight{
	if len(numbers)==0{
		hi.AddOption("number_of_fragments",numbers[0])
	}else{
		
			for k,v := range numbers{
				hi.AddFieldOption(hi.Fields[k],"number_of_fragments",v)

			}

	}
	return hi
}

func(hi *Highlight)WithFragmentSize(number_of_fragments,fragment_size int)*Highlight{
	hi2:=hi.WithNumberOfFragment(number_of_fragments)
	
	hi2.AddOption("fragment_size",fragment_size)
   return hi2

}
//
func(hi *Highlight)WithFragmenter(fragmenter string,number_of_fragments,fragment_size int)*Highlight{
        hi2:=hi.WithFragmentSize(number_of_fragments,fragment_size)
		
		hi2.AddOption("fragmentmenter",fragmenter)
		
	
	   return hi2
	

}

func(hi *Highlight)WithTags(tags ...string)*Highlight{
	if tags[0] == "styled"{
		hi.AddOption("tags_schema","styled")
	}else{
		pre:=make([]string, 0)
		post:=make([]string, 0)
		for _,tag :=range tags {
			pre= append(pre,"<"+tag+">")
			post=append(post,"</"+tag+">")
		}
		hi.AddOption("pre_tags",pre)
		hi.AddOption("post_tags", post)
		
	}
    return hi
}

/*func PlainHighlighther(fragmenter string,number_of_framents,fragment_size int)*Highlight{
	
	hi.Options["type"]="plain"
	hi.Options["fragment_size"]=fragment_size
	hi.Options["number_of_framents"]=number_of_framents
	hi.Options["fragmenter"]=fragmenter
}*/