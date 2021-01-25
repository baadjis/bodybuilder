package match
 import(
	 "fmt"
	 "strconv"
	 "reflect"
 )
// Contains check if a slice contains an item
func Contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}

 // Return keys of the given map
func Keys(m map[string]interface{}) (keys []string) {
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

// write the begining with defer
func Begin(builder Builder,str string){
	WriteQuoted(builder,str)
	builder.WriteString(":{\n")
	defer builder.WriteString("\n}")
}



 //WriteArray to a builder

func WriteArray(builder Builder,array []string){
	builder.WriteString("[")
	defer builder.WriteString("]")
	for i,el :=range array{
		quoted:=fmt.Sprintf(`%q`,el)
		builder.WriteString(quoted)
		if i!=len(array)-1{
			builder.WriteString(" ,")
		}
		
		
	
	}
	
}

//
func WriteQuoted(builder Builder,str interface{}){
	strs:=fmt.Sprintf(`%q`,str)
	builder.WriteString(strs)
}
//WriteWithFormat
func WriteWithFormat(builder Builder,key string,value interface{}){
	
	str := fmt.Sprintf(`%q: %v`,key ,value)
	if reflect.TypeOf(value).String()=="string"{
		str = fmt.Sprintf(`%q: %q`,key ,value)
	}

	
	builder.WriteString(str)
}



func WriteMap(builder Builder,maps map[string]interface{},map_name string){
	var sol string
	builder.WriteString(map_name)
	builder.WriteString(":{\n")
    for key, val := range maps {
        switch concreteVal := val.(type) {
        case map[string]interface{}:
            WriteMap(builder,concreteVal,key)
        case string:
            sol = key + ":" + concreteVal + " "
            builder.WriteString(sol)
        case float64:
            sol = key + ":" + strconv.FormatFloat(concreteVal, 'f', -1, 64) + " "
            builder.WriteString(sol)
			
		case int64:
            sol = key + ":" + strconv.FormatInt(concreteVal, 10) + " "
            builder.WriteString(sol)
        default:
            //What?
            panic("Unsupported")
		}
		builder.WriteString(",")
		
	}
	builder.WriteString("\n}\n")
}


func Encode(builder Builder,prag interface{}){
	switch v := reflect.ValueOf(prag); v.Kind(){
		
		case reflect.String:

			WriteQuoted(builder,v) 
		
		case reflect.Array,reflect.Slice:

			builder.WriteString("[")
			defer builder.WriteString("]")
		    slice:= prag.([]string)
			for i,el :=range slice{
				Encode(builder,el)
				if i!=len(slice)-1{
				      builder.WriteString(",")
                    }
			}
			
		case reflect.Map:
		    

			_map:= prag.(map[string]interface{})
			keys:=Keys(_map)
			if len(keys)>0{
				builder.WriteString("{\n")
				defer builder.WriteString("\n}")
				for i,el :=range keys{
					Encode(builder,el)
					builder.WriteString(":")
					Encode(builder,_map[el])
					if i!=len(keys)-1{
						  builder.WriteString(",")
						  builder.WriteString("\n")
						}
				}
			}else{
				builder.WriteString("{}")
			}     
			
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,reflect.Int64:
			
			_value:=prag.(int)
			format :=fmt.Sprintf(`%v`,_value)
			builder.WriteString(format)
		case reflect.Float32 ,reflect.Float64:
		    
			format :=fmt.Sprintf(`%v`,prag)
			builder.WriteString(format)
		case reflect.Bool:
			_value:=prag.(bool)
			format :=fmt.Sprintf(`%v`,_value)
			builder.WriteString(format)


		case reflect.Struct:
			struc:=v.Type()
			for i := 0; i < v.NumField(); i++ {
				Encode(builder,struc.Field(i).Name)
				builder.WriteString(":")
				Encode(builder,v.Field(i).Interface())
				if i<v.NumField()-1{
                    builder.WriteString(",\n")
				}
				
		
			}
		
		

	}

}

/*func EncodeMultimatch(builder Builder, mm *QueryMultiMatchs){
	WriteQuoted(builder,"multi_match")
	builder.WriteString(": {")
	builder.WriteString("\n")
	defer builder.WriteString("\n}\n")
	WriteQuoted(builder,"query")
	builder.WriteString(":")
	Encode(builder,mm.Query)
	builder.WriteString(",\n")

	WriteQuoted(builder,"field")
	builder.WriteString(":")
	Encode(builder,mm.Fields)

}*/