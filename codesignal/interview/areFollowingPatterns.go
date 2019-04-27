package main

import "fmt"

//import "sort"

func main() {

	strings := []string{"syf",
		"syf",
		"oxerkx",
		"oxerkx",
		"syf",
		"xgwatff",
		"pmnfaw",
		"t",
		"ajyvgwd",
		"xmhb",
		"ajg",
		"syf",
		"syf",
		"wjddgkopae",
		"fgrpstxd",
		"t",
		"i",
		"psw",
		"wjddgkopae",
		"wjddgkopae",
		"oxerkx",
		"zf",
		"jvdtdxbefr",
		"rbmphtrmo",
		"syf",
		"yssdddhyn",
		"syf",
		"jvdtdxbefr",
		"funnd",
		"syf",
		"syf",
		"wd",
		"syf",
		"vnntavj",
		"wjddgkopae",
		"yssdddhyn",
		"wcvk",
		"wjddgkopae",
		"fh",
		"zf",
		"gpkdcwf",
		"qkbw",
		"zf",
		"teppnr",
		"jvdtdxbefr",
		"fmn",
		"i",
		"hzmihfrmq",
		"wjddgkopae",
		"syf",
		"vnntavj",
		"dung",
		"kn",
		"qkxo",
		"ajyvgwd",
		"fs",
		"kanixyaepl",
		"syf",
		"tl",
		"yzhaa",
		"dung",
		"wa",
		"syf",
		"jtucivim",
		"tl",
		"kanixyaepl",
		"oxerkx",
		"wjddgkopae",
		"ey",
		"ai",
		"zf",
		"di",
		"oxerkx",
		"dung",
		"i",
		"oxerkx",
		"wmtqpwzgh",
		"t",
		"beascd",
		"me",
		"akklwhtpi",
		"nxl",
		"cnq",
		"bighexy",
		"ddhditvzdu",
		"funnd",
		"wmt",
		"dgx",
		"fs",
		"xmhb",
		"qtcxvdcl",
		"thbmn",
		"pkrisgmr",
		"mkcfscyb",
		"x",
		"oxerkx",
		"funnd",
		"iesr",
		"funnd",
		"t",
	}
	patterns := []string{"enrylabgky",
		"enrylabgky",
		"dqlqaihd",
		"dqlqaihd",
		"enrylabgky",
		"ramsnzhyr",
		"tkibsntkbr",
		"l",
		"bgtws",
		"xwuaep",
		"o",
		"enrylabgky",
		"enrylabgky",
		"e",
		"auljuhtj",
		"l",
		"d",
		"jfzokgt",
		"e",
		"e",
		"dqlqaihd",
		"fgglhiedk",
		"nj",
		"quhv",
		"enrylabgky",
		"oadats",
		"enrylabgky",
		"nj",
		"zwupro",
		"enrylabgky",
		"enrylabgky",
		"pyw",
		"enrylabgky",
		"bedpuycdp",
		"e",
		"oadats",
		"i",
		"e",
		"fobyfznrxm",
		"fgglhiedk",
		"irxtd",
		"oyvf",
		"fgglhiedk",
		"ebpp",
		"nj",
		"p",
		"d",
		"cufxylz",
		"e",
		"enrylabgky",
		"bedpuycdp",
		"mitzb",
		"shsnw",
		"papmvh",
		"bgtws",
		"chtp",
		"pze",
		"enrylabgky",
		"klp",
		"wpx",
		"mitzb",
		"fo",
		"enrylabgky",
		"bvcigrirhe",
		"klp",
		"pze",
		"dqlqaihd",
		"e",
		"iufunacwjo",
		"bubgww",
		"fgglhiedk",
		"og",
		"dqlqaihd",
		"mitzb",
		"d",
		"dqlqaihd",
		"mysidv",
		"l",
		"naj",
		"clftmrwl",
		"fjb",
		"zjjnrffb",
		"sh",
		"gcn",
		"ouispza",
		"zwupro",
		"c",
		"rdank",
		"chtp",
		"xwuaep",
		"jufhm",
		"iyntbgm",
		"sufs",
		"mkivpe",
		"bxdd",
		"dqlqaihd",
		"zwupro",
		"vzxbbculgv",
		"zwupro",
		"l",
	}

	stringsMap := make(map[string]int)
	patternsMap := make(map[string]int)

	for _, x := range strings {
		stringsMap[x]++
	}

	for _, x := range patterns {
		patternsMap[x]++
	}

	if len(patternsMap) != len(stringsMap) {
		fmt.Println("false")
		return false
	}

	var stringArry []int
	var patternsArry []int

	for _, i := range stringsMap {
		stringArry = append(stringArry, i)
	}

	for _, i := range patternsMap {
		patternsArry = append(patternsArry, i)
	}

	sum1 := 0
	sum2 := 0

	for _, x := range stringArry {
		sum1 += x
	}

	for _, x := range patternsArry {
		sum2 += x
	}

	if sum1 == sum2 {
		return true
	}

	for i := 0; i < len(strings); i++ {
		if patternsMap[strings[i]] != stringsMap[patterns[i]] {
			fmt.Println("false", "arrays are not equal")
			return false
		}
	}

	fmt.Println("true")

	return true
}
