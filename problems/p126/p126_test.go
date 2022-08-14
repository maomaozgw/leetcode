package p126

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func Test_findLadders(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Exampe 1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: [][]string{
				{"hit", "hot", "dot", "dog", "cog"},
				{"hit", "hot", "lot", "log", "cog"},
			},
		},
		{
			name: "Example 2",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList: []string{
					"hot", "dot", "dog", "lot", "log",
				},
			},
			want: [][]string{},
		},
		{
			name: "WA a",
			args: args{
				beginWord: "a",
				endWord:   "c",
				wordList: []string{
					"a", "b", "c",
				},
			},
			want: [][]string{
				{
					"a", "c",
				},
			},
		},
		{
			name: "WA 3",
			args: args{
				beginWord: "hot",
				endWord:   "dog",
				wordList: []string{
					"hot", "dog", "dot",
				},
			},
			want: [][]string{
				{"hot", "dot", "dog"},
			},
		},
		{
			name: "WA 4",
			args: args{
				beginWord: "teach",
				endWord:   "place",
				wordList: []string{
					"peale", "wilts", "place", "fetch", "purer", "pooch", "peace", "poach", "berra", "teach", "rheum", "peach",
				},
			},
			want: [][]string{
				{"teach", "peach", "peace", "place"},
			},
		},
		{
			name: "WA 5",
			args: args{
				beginWord: "qa",
				endWord:   "sq",
				wordList: []string{
					"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar",
					"ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma",
					"re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge",
					"th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh",
					"wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be",
					"fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye",
				},
			},
			want: [][]string{
				{"qa", "ca", "ci", "si", "sq"}, {"qa", "ta", "ti", "si", "sq"}, {"qa", "la", "li", "si", "sq"},
				{"qa", "ba", "bi", "si", "sq"}, {"qa", "ha", "hi", "si", "sq"}, {"qa", "pa", "pi", "si", "sq"},
				{"qa", "ma", "mi", "si", "sq"}, {"qa", "na", "ni", "si", "sq"}, {"qa", "la", "le", "se", "sq"},
				{"qa", "fa", "fe", "se", "sq"}, {"qa", "ra", "re", "se", "sq"}, {"qa", "ga", "ge", "se", "sq"},
				{"qa", "ma", "me", "se", "sq"}, {"qa", "na", "ne", "se", "sq"}, {"qa", "ba", "be", "se", "sq"},
				{"qa", "ha", "he", "se", "sq"}, {"qa", "ya", "ye", "se", "sq"}, {"qa", "ga", "go", "so", "sq"},
				{"qa", "ta", "to", "so", "sq"}, {"qa", "ya", "yo", "so", "sq"}, {"qa", "pa", "po", "so", "sq"},
				{"qa", "ha", "ho", "so", "sq"}, {"qa", "la", "lo", "so", "sq"}, {"qa", "ca", "co", "so", "sq"},
				{"qa", "na", "no", "so", "sq"}, {"qa", "ma", "mo", "so", "sq"}, {"qa", "ma", "mb", "sb", "sq"},
				{"qa", "na", "nb", "sb", "sq"}, {"qa", "ya", "yb", "sb", "sq"}, {"qa", "pa", "pb", "sb", "sq"},
				{"qa", "ra", "rb", "sb", "sq"}, {"qa", "ta", "tb", "sb", "sq"}, {"qa", "ca", "cm", "sm", "sq"},
				{"qa", "ta", "tm", "sm", "sq"}, {"qa", "pa", "pm", "sm", "sq"}, {"qa", "fa", "fm", "sm", "sq"},
				{"qa", "la", "ln", "sn", "sq"}, {"qa", "ra", "rn", "sn", "sq"}, {"qa", "ma", "mn", "sn", "sq"},
				{"qa", "ba", "br", "sr", "sq"}, {"qa", "ca", "cr", "sr", "sq"}, {"qa", "fa", "fr", "sr", "sq"},
				{"qa", "ma", "mr", "sr", "sq"}, {"qa", "la", "lr", "sr", "sq"}, {"qa", "pa", "ph", "sh", "sq"},
				{"qa", "ra", "rh", "sh", "sq"}, {"qa", "ta", "th", "sh", "sq"}, {"qa", "ma", "mt", "st", "sq"},
				{"qa", "la", "lt", "st", "sq"}, {"qa", "pa", "pt", "st", "sq"}, {"qa", "ta", "tc", "sc", "sq"},
			},
		},
		{
			name: "TE 1",
			args: args{
				beginWord: "aaaaa",
				endWord:   "ggggg",
				wordList: []string{
					"aaaaa", "caaaa", "cbaaa", "daaaa", "dbaaa", "eaaaa", "ebaaa", "faaaa", "fbaaa", "gaaaa", "gbaaa",
					"haaaa", "hbaaa", "iaaaa", "ibaaa", "jaaaa", "jbaaa", "kaaaa", "kbaaa", "laaaa", "lbaaa", "maaaa",
					"mbaaa", "naaaa", "nbaaa", "oaaaa", "obaaa", "paaaa", "pbaaa", "bbaaa", "bbcaa", "bbcba", "bbdaa",
					"bbdba", "bbeaa", "bbeba", "bbfaa", "bbfba", "bbgaa", "bbgba", "bbhaa", "bbhba", "bbiaa", "bbiba",
					"bbjaa", "bbjba", "bbkaa", "bbkba", "bblaa", "bblba", "bbmaa", "bbmba", "bbnaa", "bbnba", "bboaa",
					"bboba", "bbpaa", "bbpba", "bbbba", "abbba", "acbba", "dbbba", "dcbba", "ebbba", "ecbba", "fbbba",
					"fcbba", "gbbba", "gcbba", "hbbba", "hcbba", "ibbba", "icbba", "jbbba", "jcbba", "kbbba", "kcbba",
					"lbbba", "lcbba", "mbbba", "mcbba", "nbbba", "ncbba", "obbba", "ocbba", "pbbba", "pcbba", "ccbba",
					"ccaba", "ccaca", "ccdba", "ccdca", "cceba", "cceca", "ccfba", "ccfca", "ccgba", "ccgca", "cchba",
					"cchca", "cciba", "ccica", "ccjba", "ccjca", "cckba", "cckca", "cclba", "cclca", "ccmba", "ccmca",
					"ccnba", "ccnca", "ccoba", "ccoca", "ccpba", "ccpca", "cccca", "accca", "adcca", "bccca", "bdcca",
					"eccca", "edcca", "fccca", "fdcca", "gccca", "gdcca", "hccca", "hdcca", "iccca", "idcca", "jccca",
					"jdcca", "kccca", "kdcca", "lccca", "ldcca", "mccca", "mdcca", "nccca", "ndcca", "occca", "odcca",
					"pccca", "pdcca", "ddcca", "ddaca", "ddada", "ddbca", "ddbda", "ddeca", "ddeda", "ddfca", "ddfda",
					"ddgca", "ddgda", "ddhca", "ddhda", "ddica", "ddida", "ddjca", "ddjda", "ddkca", "ddkda", "ddlca",
					"ddlda", "ddmca", "ddmda", "ddnca", "ddnda", "ddoca", "ddoda", "ddpca", "ddpda", "dddda", "addda",
					"aedda", "bddda", "bedda", "cddda", "cedda", "fddda", "fedda", "gddda", "gedda", "hddda", "hedda",
					"iddda", "iedda", "jddda", "jedda", "kddda", "kedda", "lddda", "ledda", "mddda", "medda", "nddda",
					"nedda", "oddda", "oedda", "pddda", "pedda", "eedda", "eeada", "eeaea", "eebda", "eebea", "eecda",
					"eecea", "eefda", "eefea", "eegda", "eegea", "eehda", "eehea", "eeida", "eeiea", "eejda", "eejea",
					"eekda", "eekea", "eelda", "eelea", "eemda", "eemea", "eenda", "eenea", "eeoda", "eeoea", "eepda",
					"eepea", "eeeea", "ggggg", "agggg", "ahggg", "bgggg", "bhggg", "cgggg", "chggg", "dgggg", "dhggg",
					"egggg", "ehggg", "fgggg", "fhggg", "igggg", "ihggg", "jgggg", "jhggg", "kgggg", "khggg", "lgggg",
					"lhggg", "mgggg", "mhggg", "ngggg", "nhggg", "ogggg", "ohggg", "pgggg", "phggg", "hhggg", "hhagg",
					"hhahg", "hhbgg", "hhbhg", "hhcgg", "hhchg", "hhdgg", "hhdhg", "hhegg", "hhehg", "hhfgg", "hhfhg",
					"hhigg", "hhihg", "hhjgg", "hhjhg", "hhkgg", "hhkhg", "hhlgg", "hhlhg", "hhmgg", "hhmhg", "hhngg",
					"hhnhg", "hhogg", "hhohg", "hhpgg", "hhphg", "hhhhg", "ahhhg", "aihhg", "bhhhg", "bihhg", "chhhg",
					"cihhg", "dhhhg", "dihhg", "ehhhg", "eihhg", "fhhhg", "fihhg", "ghhhg", "gihhg", "jhhhg", "jihhg",
					"khhhg", "kihhg", "lhhhg", "lihhg", "mhhhg", "mihhg", "nhhhg", "nihhg", "ohhhg", "oihhg", "phhhg",
					"pihhg", "iihhg", "iiahg", "iiaig", "iibhg", "iibig", "iichg", "iicig", "iidhg", "iidig", "iiehg",
					"iieig", "iifhg", "iifig", "iighg", "iigig", "iijhg", "iijig", "iikhg", "iikig", "iilhg", "iilig",
					"iimhg", "iimig", "iinhg", "iinig", "iiohg", "iioig", "iiphg", "iipig", "iiiig", "aiiig", "ajiig",
					"biiig", "bjiig", "ciiig", "cjiig", "diiig", "djiig", "eiiig", "ejiig", "fiiig", "fjiig", "giiig",
					"gjiig", "hiiig", "hjiig", "kiiig", "kjiig", "liiig", "ljiig", "miiig", "mjiig", "niiig", "njiig",
					"oiiig", "ojiig", "piiig", "pjiig", "jjiig", "jjaig", "jjajg", "jjbig", "jjbjg", "jjcig", "jjcjg",
					"jjdig", "jjdjg", "jjeig", "jjejg", "jjfig", "jjfjg", "jjgig", "jjgjg", "jjhig", "jjhjg", "jjkig",
					"jjkjg", "jjlig", "jjljg", "jjmig", "jjmjg", "jjnig", "jjnjg", "jjoig", "jjojg", "jjpig", "jjpjg",
					"jjjjg", "ajjjg", "akjjg", "bjjjg", "bkjjg", "cjjjg", "ckjjg", "djjjg", "dkjjg", "ejjjg", "ekjjg",
					"fjjjg", "fkjjg", "gjjjg", "gkjjg", "hjjjg", "hkjjg", "ijjjg", "ikjjg", "ljjjg", "lkjjg", "mjjjg",
					"mkjjg", "njjjg", "nkjjg", "ojjjg", "okjjg", "pjjjg", "pkjjg", "kkjjg", "kkajg", "kkakg", "kkbjg",
					"kkbkg", "kkcjg", "kkckg", "kkdjg", "kkdkg", "kkejg", "kkekg", "kkfjg", "kkfkg", "kkgjg", "kkgkg",
					"kkhjg", "kkhkg", "kkijg", "kkikg", "kkljg", "kklkg", "kkmjg", "kkmkg", "kknjg", "kknkg", "kkojg",
					"kkokg", "kkpjg", "kkpkg", "kkkkg", "ggggx", "gggxx", "ggxxx", "gxxxx", "xxxxx", "xxxxy", "xxxyy",
					"xxyyy", "xyyyy", "yyyyy", "yyyyw", "yyyww", "yywww", "ywwww", "wwwww", "wwvww", "wvvww", "vvvww",
					"vvvwz", "avvwz", "aavwz", "aaawz", "aaaaz",
				},
			},
			want: [][]string{
				{
					"aaaaa", "aaaaz", "aaawz", "aavwz", "avvwz", "vvvwz", "vvvww", "wvvww", "wwvww", "wwwww", "ywwww",
					"yywww", "yyyww", "yyyyw", "yyyyy", "xyyyy", "xxyyy", "xxxyy", "xxxxy", "xxxxx", "gxxxx", "ggxxx",
					"gggxx", "ggggx", "ggggg",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLadders(tt.args.beginWord, tt.args.endWord, tt.args.wordList); !cmp.Equal(got, tt.want, cmpopts.EquateEmpty()) {
				content, _ := json.Marshal(got)
				fmt.Println(string(content))
				t.Errorf("findLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
