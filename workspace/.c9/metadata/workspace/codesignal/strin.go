{"filter":false,"title":"strin.go","tooltip":"/codesignal/strin.go","undoManager":{"mark":0,"position":0,"stack":[[{"start":{"row":0,"column":0},"end":{"row":163,"column":0},"action":"insert","lines":["package main","","import (","\t\"fmt\"","\t\"strings\"",")","","/*","This one is a bit of a doozy, in C++ I recommend two functions:","","bool adjacentNodes(string1, string2) - compares two strings and only returns true if the two are different by one character. This function can terminate early if the number of different characters are greater than 2.","","-bool stringsRearrangement (vector<string>):","-- do","-- set test to bool true","--- for every adjacent pair of words, test if adjacent","---- if NOT adjacent test false","-- while (next_permutation of (inputarray) and !test)","-- return test","","I think this method is preferable to creating a graph because there are a lot of opportunities for early termination, so while there may be inputArray.size()! number of permutations, you'll run through them very quickly.","","","*/","","func main() {","","\tinputArray := []string{\"abc\", \"bef\", \"bcc\", \"bec\", \"bbc\", \"bdc\"}","","\tinputArray3 := []string{\"abc\", \"bef\", \"bcc\", \"bec\", \"bbc\", \"bdc\"}","","\tif len(inputArray) == 1 {","\t\tfmt.Println(\"false\")","\t\treturn","\t}","","\tcountTest := 0","","\tfor i := 0; i < len(inputArray); i++ {","\t\tif inputArray[i] == inputArray3[i] {","\t\t\tcountTest++","\t\t}","\t}","","\tif countTest == len(inputArray) {","\t\tfmt.Println(\"true\")","\t\treturn","\t}","","\tcountequal := 0","","\tfor i := 0; i < len(inputArray)-1; i++ {","\t\tif inputArray[i] == inputArray[i+1] {","\t\t\tcountequal++","\t\t}","\t}","","\tif countequal == len(inputArray)-1 {","\t\tfmt.Println(\"false\")","\t\treturn","\t}","","\tinputArray2 := make([]string, len(inputArray))","","\tfmt.Println(len(inputArray2))","","\tcopy(inputArray2, inputArray)","","\t//fmt.Println(inputArray2)","","\t//brute force permutations","","\tfor m := 0; m < len(inputArray); m++ {","","\t\tfor i := 0; i < len(inputArray); i++ {","\t\t\tfor j := 0; j < len(inputArray); j++ {","\t\t\t\t//\tif i == j {","\t\t\t\t//\t\tcontinue","\t\t\t\t//\t} else {","\t\t\t\tinputArray2[i], inputArray2[j] = inputArray2[j], inputArray2[i]","\t\t\t\t//\t}","","\t\t\t\tfmt.Println(inputArray2)","\t\t\t\ttestg := holderf(inputArray2)","","\t\t\t\tif testg == true {","\t\t\t\t\tfmt.Println(\"final true\")","\t\t\t\t\treturn","\t\t\t\t}","\t\t\t\t//\tcopy(inputArray2, inputArray)","\t\t\t}","\t\t\t//copy(inputArray2, inputArray)","\t\t}","\t}","","\tfmt.Println(\"false\")","","}","","func holderf(a []string) bool {","","\tcount := 0","\tfor i := 0; i < len(a)-1; i++ {","\t\ttest := strTest(a[i], a[i+1])","\t\t//fmt.Println(test)","","\t\tif test == true {","\t\t\tcount++","\t\t}","","\t}","\tif count == len(a)-1 {","\t\treturn true","\t}","","\treturn false","}","","func strTest(a string, b string) bool {","\tcount := 0","","\tfor _, x := range a {","\t\ttest := strings.ContainsAny(b, string(x))","\t\tif test == false {","\t\t\tcount++","\t\t}","\t}","","\tif count > 1 {","\t\treturn false","\t}","","\tcount = 0","","\tfor _, x := range b {","\t\ttest := strings.ContainsAny(a, string(x))","\t\tif test == false {","\t\t\tcount++","\t\t}","\t}","","\tif count > 1 {","\t\treturn false","\t}","","\t//return true","\tcount = 0","","\tfor i := 0; i < len(a); i++ {","\t\tif string(a[i]) != string(b[i]) {","\t\t\tcount++","\t\t}","\t}","","\t//\tfmt.Println(count)","","\tif count != 1 {","\t\treturn false","\t}","","\treturn true","","}",""],"id":1}]]},"ace":{"folds":[],"scrolltop":1973.25,"scrollleft":0,"selection":{"start":{"row":163,"column":0},"end":{"row":163,"column":0},"isBackwards":false},"options":{"guessTabSize":true,"useWrapMode":false,"wrapToView":true},"firstLineState":{"row":141,"state":"start","mode":"ace/mode/golang"}},"timestamp":1538085830254,"hash":"ad9ae9c0dc7855d3d77a29e75b6a5803468c0ad3"}