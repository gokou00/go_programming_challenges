{"filter":false,"title":"digitsProduct2.go","tooltip":"/codesignal/digitsProduct2.go","ace":{"folds":[],"scrolltop":97,"scrollleft":0,"selection":{"start":{"row":18,"column":0},"end":{"row":18,"column":0},"isBackwards":false},"options":{"guessTabSize":true,"useWrapMode":false,"wrapToView":true},"firstLineState":{"row":5,"state":"start","mode":"ace/mode/golang"}},"hash":"921ca67fc6ed5ee1b265dab3e2383dd5152ac6e6","undoManager":{"mark":56,"position":56,"stack":[[{"start":{"row":0,"column":0},"end":{"row":45,"column":0},"action":"insert","lines":["package main","","import \"fmt\"","import \"strconv\"","import \"sort\"","","func main() {","\tproduct := 450","\tnum := \"\"","","\tvar factorArray []int","\tvar factorArrayStr []int","","\tfor i := 2; i < product; i++ {","\t\tif product%i == 0 {","\t\t\tfactorArray = append(factorArray, i)","\t\t}","\t}","","\tfmt.Println(factorArray)","","\tfor i := 0; i < len(factorArray); i++ {","\t\tfor j := 0; j < len(factorArray); j++ {","\t\t\tif j == i {","\t\t\t\tcontinue","\t\t\t}","","\t\t\tif factorArray[i]*factorArray[j] == product {","\t\t\t\tnum += strconv.Itoa(factorArray[i])","\t\t\t\tnum += strconv.Itoa(factorArray[j])","\t\t\t\ttemp, _ := strconv.Atoi(num)","\t\t\t\tfactorArrayStr = append(factorArrayStr, temp)","\t\t\t\tnum = \"\"","","\t\t\t}","","\t\t}","","\t}","","\tsort.Ints(factorArrayStr)","","\tfmt.Println(factorArrayStr)","","}",""],"id":1}],[{"start":{"row":19,"column":25},"end":{"row":20,"column":0},"action":"insert","lines":["",""],"id":2},{"start":{"row":20,"column":0},"end":{"row":20,"column":1},"action":"insert","lines":["\t"]}],[{"start":{"row":20,"column":1},"end":{"row":21,"column":0},"action":"insert","lines":["",""],"id":3},{"start":{"row":21,"column":0},"end":{"row":21,"column":1},"action":"insert","lines":["\t"]}],[{"start":{"row":21,"column":1},"end":{"row":21,"column":2},"action":"insert","lines":["/"],"id":4}],[{"start":{"row":21,"column":2},"end":{"row":21,"column":3},"action":"insert","lines":["*"],"id":5}],[{"start":{"row":45,"column":0},"end":{"row":45,"column":1},"action":"insert","lines":[" "],"id":6}],[{"start":{"row":45,"column":1},"end":{"row":45,"column":2},"action":"insert","lines":["*"],"id":7}],[{"start":{"row":45,"column":2},"end":{"row":45,"column":3},"action":"insert","lines":["/"],"id":8}],[{"start":{"row":14,"column":21},"end":{"row":14,"column":22},"action":"insert","lines":[" "],"id":9}],[{"start":{"row":14,"column":22},"end":{"row":14,"column":23},"action":"insert","lines":["/"],"id":10}],[{"start":{"row":14,"column":23},"end":{"row":14,"column":24},"action":"insert","lines":["/"],"id":11}],[{"start":{"row":14,"column":24},"end":{"row":14,"column":25},"action":"insert","lines":["i"],"id":12}],[{"start":{"row":14,"column":25},"end":{"row":14,"column":26},"action":"insert","lines":["f"],"id":13}],[{"start":{"row":14,"column":26},"end":{"row":14,"column":27},"action":"insert","lines":[" "],"id":14}],[{"start":{"row":14,"column":27},"end":{"row":14,"column":28},"action":"insert","lines":["i"],"id":15}],[{"start":{"row":14,"column":28},"end":{"row":14,"column":29},"action":"insert","lines":[" "],"id":16}],[{"start":{"row":14,"column":29},"end":{"row":14,"column":30},"action":"insert","lines":[">"],"id":17}],[{"start":{"row":14,"column":30},"end":{"row":14,"column":31},"action":"insert","lines":[" "],"id":18}],[{"start":{"row":14,"column":31},"end":{"row":14,"column":32},"action":"insert","lines":["9"],"id":19}],[{"start":{"row":14,"column":32},"end":{"row":14,"column":33},"action":"insert","lines":[" "],"id":20}],[{"start":{"row":14,"column":33},"end":{"row":14,"column":34},"action":"insert","lines":["c"],"id":21}],[{"start":{"row":14,"column":34},"end":{"row":14,"column":35},"action":"insert","lines":["a"],"id":22}],[{"start":{"row":14,"column":35},"end":{"row":14,"column":36},"action":"insert","lines":["l"],"id":23}],[{"start":{"row":14,"column":36},"end":{"row":14,"column":37},"action":"insert","lines":["l"],"id":24}],[{"start":{"row":14,"column":37},"end":{"row":14,"column":38},"action":"insert","lines":[" "],"id":25}],[{"start":{"row":14,"column":38},"end":{"row":14,"column":39},"action":"insert","lines":["f"],"id":26}],[{"start":{"row":14,"column":39},"end":{"row":14,"column":40},"action":"insert","lines":["u"],"id":27}],[{"start":{"row":14,"column":40},"end":{"row":14,"column":41},"action":"insert","lines":["n"],"id":28}],[{"start":{"row":14,"column":41},"end":{"row":14,"column":42},"action":"insert","lines":["c"],"id":29}],[{"start":{"row":14,"column":42},"end":{"row":14,"column":43},"action":"insert","lines":["t"],"id":30}],[{"start":{"row":14,"column":43},"end":{"row":14,"column":44},"action":"insert","lines":["i"],"id":31}],[{"start":{"row":14,"column":44},"end":{"row":14,"column":45},"action":"insert","lines":["o"],"id":32}],[{"start":{"row":14,"column":45},"end":{"row":14,"column":46},"action":"insert","lines":["n"],"id":33}],[{"start":{"row":14,"column":46},"end":{"row":14,"column":47},"action":"insert","lines":[" "],"id":34}],[{"start":{"row":14,"column":47},"end":{"row":14,"column":48},"action":"insert","lines":["t"],"id":35}],[{"start":{"row":14,"column":48},"end":{"row":14,"column":49},"action":"insert","lines":["o"],"id":36}],[{"start":{"row":14,"column":49},"end":{"row":14,"column":50},"action":"insert","lines":[" "],"id":37}],[{"start":{"row":14,"column":50},"end":{"row":14,"column":51},"action":"insert","lines":["s"],"id":38}],[{"start":{"row":14,"column":51},"end":{"row":14,"column":52},"action":"insert","lines":["e"],"id":39}],[{"start":{"row":14,"column":52},"end":{"row":14,"column":53},"action":"insert","lines":["p"],"id":40}],[{"start":{"row":14,"column":53},"end":{"row":14,"column":54},"action":"insert","lines":["a"],"id":41}],[{"start":{"row":14,"column":54},"end":{"row":14,"column":55},"action":"insert","lines":["r"],"id":42}],[{"start":{"row":14,"column":55},"end":{"row":14,"column":56},"action":"insert","lines":["a"],"id":43}],[{"start":{"row":14,"column":56},"end":{"row":14,"column":57},"action":"insert","lines":["t"],"id":44}],[{"start":{"row":14,"column":57},"end":{"row":14,"column":58},"action":"insert","lines":["e"],"id":45}],[{"start":{"row":14,"column":58},"end":{"row":14,"column":59},"action":"insert","lines":[" "],"id":46}],[{"start":{"row":14,"column":59},"end":{"row":14,"column":60},"action":"insert","lines":["t"],"id":47}],[{"start":{"row":14,"column":60},"end":{"row":14,"column":61},"action":"insert","lines":["h"],"id":48}],[{"start":{"row":14,"column":61},"end":{"row":14,"column":62},"action":"insert","lines":["e"],"id":49}],[{"start":{"row":14,"column":62},"end":{"row":14,"column":63},"action":"insert","lines":[" "],"id":50}],[{"start":{"row":14,"column":63},"end":{"row":14,"column":64},"action":"insert","lines":["d"],"id":51}],[{"start":{"row":14,"column":64},"end":{"row":14,"column":65},"action":"insert","lines":["i"],"id":52}],[{"start":{"row":14,"column":65},"end":{"row":14,"column":66},"action":"insert","lines":["g"],"id":53}],[{"start":{"row":14,"column":66},"end":{"row":14,"column":67},"action":"insert","lines":["i"],"id":54}],[{"start":{"row":14,"column":67},"end":{"row":14,"column":68},"action":"insert","lines":["t"],"id":55}],[{"start":{"row":14,"column":68},"end":{"row":14,"column":69},"action":"insert","lines":["s"],"id":56}],[{"start":{"row":20,"column":0},"end":{"row":20,"column":1},"action":"remove","lines":["\t"],"id":57,"ignore":true},{"start":{"row":23,"column":0},"end":{"row":23,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":24,"column":0},"end":{"row":24,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":25,"column":0},"end":{"row":25,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":26,"column":4},"end":{"row":26,"column":5},"action":"insert","lines":["\t"]},{"start":{"row":27,"column":0},"end":{"row":27,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":29,"column":3},"end":{"row":29,"column":4},"action":"insert","lines":["\t"]},{"start":{"row":30,"column":0},"end":{"row":30,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":31,"column":4},"end":{"row":31,"column":5},"action":"insert","lines":["\t"]},{"start":{"row":32,"column":0},"end":{"row":32,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":33,"column":4},"end":{"row":33,"column":5},"action":"insert","lines":["\t"]},{"start":{"row":34,"column":4},"end":{"row":34,"column":5},"action":"insert","lines":["\t"]},{"start":{"row":36,"column":3},"end":{"row":36,"column":4},"action":"insert","lines":["\t"]},{"start":{"row":38,"column":0},"end":{"row":38,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":40,"column":0},"end":{"row":40,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":42,"column":0},"end":{"row":42,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":44,"column":0},"end":{"row":44,"column":1},"action":"insert","lines":["\t"]},{"start":{"row":45,"column":0},"end":{"row":45,"column":1},"action":"remove","lines":[" "]},{"start":{"row":45,"column":0},"end":{"row":45,"column":1},"action":"insert","lines":["\t"]}]]},"timestamp":1540305160327}