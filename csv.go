package golang_csv_deserialization

import (
	"bufio"
	"encoding/csv"
	"errors"
	"github.com/samber/lo"
	"io"
	"reflect"
	"strconv"
)

// column 的数据类型
const (
	Int              = "int"                // basic type
	Float            = "float"              // basic type
	Bool             = "bool"               // basic type
	String           = "string"             // basic type
	ArrInt           = "arr:int"            // Arr type, eg: 299792458,222,333
	ArrFloat         = "arr:float"          // Arr type, eg: 3.14159,2.71828,1.61803,1.41421
	ArrBool          = "arr:bool"           // Arr type, eg: true,false,true
	ArrString        = "arr:string"         // Arr type, eg: 我可以吃玻璃, 而不伤害身体,不伤害身体
	MapInt2Int       = "map(int:int)"       // Map type, eg：1:111,2:222,3:333
	MapInt2Float     = "map(int:float)"     // Map type, eg: 1:3.14159,2:2.71828,3:1.61803
	MapInt2Bool      = "map(int:bool)"      // Map type
	MapInt2String    = "map(int:string)"    // Map type
	MapFloat2Int     = "map(float:int)"     // Map type
	MapFloat2Float   = "map(float:float)"   // Map type
	MapFloat2Bool    = "map(float:bool)"    // Map type
	MapFloat2String  = "map(float:string)"  // Map type
	MapString2Int    = "map(string:int)"    // Map type
	MapString2Float  = "map(string:float)"  // Map type
	MapString2Bool   = "map(string:bool)"   // Map type
	MapString2String = "map(string:string)" // Map type
)

type header struct {
	name string // 列名
	typ  string // 列数据类型
}

type Csv struct {
	headers []*header
	records [][]string
	cursor  int
}

func (c *Csv) Read(reader io.Reader) {
	r := bufio.NewReader(reader)
	// readRune, _, err := r.ReadRune()
	// AssertNoError("read rune", err)
	// if readRune != '\uFEFF' {
	// 	_ = r.UnreadRune()
	// }

	csvReader := csv.NewReader(r)
	firstLine, err := csvReader.Read()
	AssertNoError("csv read first line", err)
	secondLine, err := csvReader.Read()
	AssertNoError("csv read second line", err)

	list, err := csvReader.ReadAll()
	AssertNoError("csv read all record", err)

	headers := make([]*header, len(firstLine))
	for i := 0; i < len(firstLine); i++ {
		name := firstLine[i]
		typ := secondLine[i]
		headers[i] = &header{name: name, typ: typ}
	}
	c.headers = headers
	c.checkHeaders() // 检查header是否是在预定义的列表中
	c.records = list
	c.Check() // 检查每个值是不是能够按照header中描述的转化成对应类型的val
}

// Check 根据 headers 检查
func (c *Csv) Check() {
	for i := 0; i < len(c.headers); i++ {
		h := c.headers[i]
		for j := 0; j < len(c.records); j++ {
			val := c.records[j][i]
			parseVal(h.typ, val)
		}
	}
}

func (c *Csv) checkHeaders() {
	validTypes := []string{Int, Float, Bool, String, ArrInt, ArrFloat, ArrBool, ArrString, MapInt2Int, MapInt2Float, MapInt2Bool, MapInt2String, MapFloat2Int, MapFloat2Float, MapFloat2Bool, MapFloat2String, MapString2Int, MapString2Float, MapString2Bool, MapString2String}
	for i := 0; i < len(c.headers); i++ {
		h := c.headers[i]
		idx := lo.IndexOf(validTypes, h.typ)
		if idx != -1 {
			continue
		}
		e := errors.New("header type is not valid: " + h.typ + ". index: " + strconv.Itoa(i))
		panic(e)
	}
}

func (c *Csv) AttachOne(r interface{}) bool {

	val := reflect.ValueOf(r).Elem()

	if c.cursor >= len(c.records) {
		return false
	}

	record := c.records[c.cursor] // 使用当前游标所属的位置的记录
	c.cursor++                    // 游标后移

	for i := 0; i < val.NumField(); i++ { // 遍历这个对象所有的属性，
		typeField := val.Type().Field(i)   //  根据序号找到这个 field 数据
		csvTag := typeField.Tag.Get("csv") // tag 名字
		if csvTag == "" {                  // 如果这个属性不是一个 csv 字段，则跳过
			continue
		}

		// 确保 tag 名 在 csv 文件中存在， 如果不存在则 panic
		h, idx, exist := lo.FindIndexOf(c.headers, func(item *header) bool { return item.name == csvTag })
		if !exist {
			e := errors.New("not found csv tag: " + csvTag)
			panic(e)
		}

		setVal(val.Field(i), h.typ, record[idx]) // 将这个属性对应的字符串值 转为对应的golang类型，并写回这个属性上

	}
	return true
}

func (c *Csv) attachOne(value reflect.Value) bool {
	val := value.Elem()

	if c.cursor >= len(c.records) {
		return false
	}

	record := c.records[c.cursor] // 使用当前游标所属的位置的记录
	c.cursor++                    // 游标后移

	for i := 0; i < val.NumField(); i++ { // 遍历这个对象所有的属性，
		typeField := val.Type().Field(i)   //  根据序号找到这个 field 数据
		csvTag := typeField.Tag.Get("csv") // tag 名字
		if csvTag == "" {                  // 如果这个属性不是一个 csv 字段，则跳过
			continue
		}

		// 确保 tag 名 在 csv 文件中存在， 如果不存在则 panic
		h, idx, exist := lo.FindIndexOf(c.headers, func(item *header) bool { return item.name == csvTag })
		if !exist {
			e := errors.New("not found csv tag: " + csvTag)
			panic(e)
		}

		setVal(val.Field(i), h.typ, record[idx]) // 将这个属性对应的字符串值 转为对应的golang类型，并写回这个属性上

	}
	return true
}

func (c *Csv) Attach(list interface{}) {
	rt := reflect.TypeOf(list)
	rv := reflect.ValueOf(list)
	kind := rt.Kind()
	if kind != reflect.Ptr {
		e := errors.New("param must be a pointer of slice, but: " + kind.String())
		panic(e)
	}

	rtEE := rt.Elem().Elem()
	rvE := rv.Elem()

	tempSlice := make([]reflect.Value, 0)
	for c.hasNext() {
		element := reflect.New(rtEE)
		c.attachOne(element)
		tempSlice = append(tempSlice, element.Elem())
	}
	temp := reflect.Append(rvE, tempSlice...)
	rvE.Set(temp)
}

func (c *Csv) recordsCount() int {
	return len(c.records)
}

func (c *Csv) hasNext() bool {
	return c.cursor < len(c.records)
}

func setVal(field reflect.Value, typ string, val string) {
	switch typ {
	case Int:
		temp := P.Int(val)
		field.SetInt(int64(temp))
	case Float:
		temp := P.Float(val)
		field.SetFloat(temp)
	case Bool:
		temp := P.Bool(val)
		field.SetBool(temp)
	case String:
		temp := P.String(val)
		field.SetString(temp)
	case ArrInt:
		temp := P.ArrInt(val)
		field.Set(reflect.ValueOf(temp))
	case ArrFloat:
		temp := P.ArrFloat(val)
		field.Set(reflect.ValueOf(temp))
	case ArrBool:
		temp := P.ArrBool(val)
		field.Set(reflect.ValueOf(temp))
	case ArrString:
		temp := P.ArrString(val)
		field.Set(reflect.ValueOf(temp))
	case MapInt2Int:
		temp := P.MapInt2Int(val)
		field.Set(reflect.ValueOf(temp))
	case MapInt2Float:
		temp := P.MapInt2Float(val)
		field.Set(reflect.ValueOf(temp))
	case MapInt2Bool:
		temp := P.MapInt2Bool(val)
		field.Set(reflect.ValueOf(temp))
	case MapInt2String:
		temp := P.MapInt2String(val)
		field.Set(reflect.ValueOf(temp))
	case MapFloat2Int:
		temp := P.MapFloat2Int(val)
		field.Set(reflect.ValueOf(temp))
	case MapFloat2Float:
		temp := P.MapFloat2Float(val)
		field.Set(reflect.ValueOf(temp))
	case MapFloat2Bool:
		temp := P.MapFloat2Bool(val)
		field.Set(reflect.ValueOf(temp))
	case MapFloat2String:
		temp := P.MapFloat2String(val)
		field.Set(reflect.ValueOf(temp))
	case MapString2Int:
		temp := P.MapString2Int(val)
		field.Set(reflect.ValueOf(temp))
	case MapString2Float:
		temp := P.MapString2Float(val)
		field.Set(reflect.ValueOf(temp))
	case MapString2Bool:
		temp := P.MapString2Bool(val)
		field.Set(reflect.ValueOf(temp))
	case MapString2String:
		temp := P.MapString2String(val)
		field.Set(reflect.ValueOf(temp))
	default:
		e := errors.New("type is not valid: " + typ)
		panic(e)
	}
}

func parseVal(typ string, val string) {
	switch typ {
	case Int:
		P.Int(val)
	case Float:
		P.Float(val)
	case Bool:
		P.Bool(val)
	case String:
		P.String(val)
	case ArrInt:
		P.ArrInt(val)
	case ArrFloat:
		P.ArrFloat(val)
	case ArrBool:
		P.ArrBool(val)
	case ArrString:
		P.ArrString(val)
	case MapInt2Int:
		P.MapInt2Int(val)
	case MapInt2Float:
		P.MapInt2Float(val)
	case MapInt2Bool:
		P.MapInt2Bool(val)
	case MapInt2String:
		P.MapInt2String(val)
	case MapFloat2Int:
		P.MapFloat2Int(val)
	case MapFloat2Float:
		P.MapFloat2Float(val)
	case MapFloat2Bool:
		P.MapFloat2Bool(val)
	case MapFloat2String:
		P.MapFloat2String(val)
	case MapString2Int:
		P.MapString2Int(val)
	case MapString2Float:
		P.MapString2Float(val)
	case MapString2Bool:
		P.MapString2Bool(val)
	case MapString2String:
		P.MapString2String(val)
	default:
		e := errors.New("type is not valid: " + typ)
		panic(e)
	}
}

func AssertNoError(msg string, err error) {
	if err == nil {
		return
	}
	e := errors.New(msg + " failed, error detail: " + err.Error())
	panic(e)
}
