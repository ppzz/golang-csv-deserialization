# csv-deserialization

Define a struct object, and create an instance of this object from the csv file according to the annotation

## example

csv example：

```csv
id,     power,      is_newbie,  desc,                                       skill,                  score,                          subject
int,    float,      bool,       string,                                     arr:int,                map(int:int),                   map(string:string)
1001,   131.159,    TRUE,       "I can eat glass, it does not hurt me.",    "3701,3722,3380,3752",  "101:88, 102:90, 103:60",       "magic-type: science, weapon: electromagnetic gun"
1002,   2222.3,     FALSE,      "I can eat glass, it does not hurt me.",    "3707,3720,3391,3752",  "101: 90, 102: 99, 103: 99",    "magic-type: fire, weapon: Melville's Bone"
```

golang struct definition:

```golang
type Player struct {
    Id       int               `csv:"id"`
    Power    float64           `csv:"power"`
    IsNewbie bool              `csv:"is_newbie"`
    Desc     string            `csv:"desc"`
    Skill    []int             `csv:"skill"`
    Score    map[int]int       `csv:"score"`
    Subject  map[string]string `csv:"subject"`
}
```

* tag name `csv` will guide us to column name in csv
* id column in csv, means int field in struct definition
* field type must match column type

supported val type in csv:

| csv column type   | golang field type   |
|-------------------|---------------------|
| int               | int                 |
| float             | float               |
| bool              | bool                |
| string            | string              |
| arr:int           | []int               |
| arr:float         | []float             |
| arr:bool          | []bool              |
| arr:string        | []string            |
| map(int:int)      | map[int]int         |
| map(int:float)    | map[int]float64     |
| map(int:bool)     | map[int]bool        |
| map(int:string)   | map[int]string      |
| map(float:int)    | map[float64]int     |
| map(float:float)  | map[float64]float64 |
| map(float:bool)   | map[float64]bool    |
| map(float:string) | map[float64]string  |
| map(string:int)   | map[string]int      |
| map(string:float) | map[string]float64  |
| map(string:bool)  | map[string]bool     |
| map(string:string) | map[string]string   |

demo:

install this pkg : `go get github.com/ppzz/csv-deserialization`

```go
import csvdeserialization "github.com/ppzz/csv-deserialization"

func main() {
    f, _ := os.Open("./example/example.csv")
    defer f.Close()
    
    c := csvdeserialization.Csv{}
    c.Read(f)
    
    var list []Player
    c.Attach(&list)
    
    for i := 0; i < len(list); i++ {
    item := list[i]
    fmt.Println(i, item.Id, item.Power, item.IsNewbie, item.Desc, item.Skill, item.Score, item.Subject)
    }
}
```

or just run `./example/example.go`
