package euler

import "testing"

// PE0001 {{{
var pe0001Cases = []struct {
	Input    int
	Expected int
}{
	{10, 23},
	{1000, 233168},
}

func TestPE0001(t *testing.T) {
	for _, tc := range pe0001Cases {
		if actual := PE0001(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// }}}

// PE0002 {{{
var pe0002Cases = []struct {
	Input    int
	Expected int
}{
	{3, 2},
	{10, 10},
	{4000000, 4613732},
}

func TestPE0002(t *testing.T) {
	for _, tc := range pe0002Cases {
		if actual, _ := PE0002(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// }}}

// PE0003 {{{
var pe0003Cases = []struct {
	Input    int
	Expected int
}{
	{13195, 29},
	{600851475143, 6857},
}

func TestPE0003SortSort(t *testing.T) {
	for _, tc := range pe0003Cases {
		if actual, _ := PE0003SortSort(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestPE0003SortSlice(t *testing.T) {
	for _, tc := range pe0003Cases {
		if actual, _ := PE0003SortSlice(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestPE0003b(t *testing.T) {
	for _, tc := range pe0003Cases {
		if actual := PE0003b(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

// Benchmarks {{{

func BenchmarkPE0003SortSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0003SortSort(600851475143)
	}
}

func BenchmarkPE0003SortSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0003SortSlice(600851475143)
	}
}

func BenchmarkPE0003b(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0003b(600851475143)
	}
}

// }}}
// }}}

// PE0004 {{{
var pe0004Cases = []struct {
	Input    int
	Expected int
}{
	{2, 9009},
	{3, 906609},
}

func TestPE0004(t *testing.T) {
	for _, tc := range pe0004Cases {
		actual := PE0004(tc.Input)
		if actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// }}}

// PE0005 {{{
var pe0005Cases = []struct {
	Input    int
	Expected int
}{
	{10, 2520},
	{20, 232792560},
}

func TestPE0005(t *testing.T) {
	for _, tc := range pe0005Cases {
		if actual := PE0005(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// }}}

// PE0006 {{{
var pe0006Cases = []struct {
	Input    int
	Expected int
}{
	{10, 2640},
	{100, 25164150},
}

func TestPE0006(t *testing.T) {
	for _, tc := range pe0006Cases {
		if actual := PE0006(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

// }}}

// PE0007 {{{
var pe0007Cases = []struct {
	Input    int
	Expected int
}{
	{1, 2},
	{2, 3},
	{3, 5},
	{4, 7},
	{10001, 104743},
}

func TestPE0007(t *testing.T) {
	for _, tc := range pe0007Cases {
		if actual := PE0007(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// }}}

// PE0008 {{{
var pe0008Cases = []struct {
	InputAdjacentNumber int
	InputGrid           string
	Expected            int
}{
	{4, "73167176531330624919225119674426574742355349194934" +
		"96983520312774506326239578318016984801869478851843" +
		"85861560789112949495459501737958331952853208805511" +
		"12540698747158523863050715693290963295227443043557" +
		"66896648950445244523161731856403098711121722383113" +
		"62229893423380308135336276614282806444486645238749" +
		"30358907296290491560440772390713810515859307960866" +
		"70172427121883998797908792274921901699720888093776" +
		"65727333001053367881220235421809751254540594752243" +
		"52584907711670556013604839586446706324415722155397" +
		"53697817977846174064955149290862569321978468622482" +
		"83972241375657056057490261407972968652414535100474" +
		"82166370484403199890008895243450658541227588666881" +
		"16427171479924442928230863465674813919123162824586" +
		"17866458359124566529476545682848912883142607690042" +
		"24219022671055626321111109370544217506941658960408" +
		"07198403850962455444362981230987879927244284909188" +
		"84580156166097919133875499200524063689912560717606" +
		"05886116467109405077541002256983155200055935729725" +
		"71636269561882670428252483600823257530420752963450", 5832},
	{13, "73167176531330624919225119674426574742355349194934" +
		"96983520312774506326239578318016984801869478851843" +
		"85861560789112949495459501737958331952853208805511" +
		"12540698747158523863050715693290963295227443043557" +
		"66896648950445244523161731856403098711121722383113" +
		"62229893423380308135336276614282806444486645238749" +
		"30358907296290491560440772390713810515859307960866" +
		"70172427121883998797908792274921901699720888093776" +
		"65727333001053367881220235421809751254540594752243" +
		"52584907711670556013604839586446706324415722155397" +
		"53697817977846174064955149290862569321978468622482" +
		"83972241375657056057490261407972968652414535100474" +
		"82166370484403199890008895243450658541227588666881" +
		"16427171479924442928230863465674813919123162824586" +
		"17866458359124566529476545682848912883142607690042" +
		"24219022671055626321111109370544217506941658960408" +
		"07198403850962455444362981230987879927244284909188" +
		"84580156166097919133875499200524063689912560717606" +
		"05886116467109405077541002256983155200055935729725" +
		"71636269561882670428252483600823257530420752963450", 23514624000},
}

func TestPE0008(t *testing.T) {
	for _, tc := range pe0008Cases {
		if actual, _ := PE0008(tc.InputAdjacentNumber, tc.InputGrid); actual != tc.Expected {
			t.Errorf(
				"InputAdjacentNumber:%v\nInputGrid:%v\nExpected:%v\nActual:%v",
				tc.InputAdjacentNumber,
				tc.InputGrid,
				tc.Expected,
				actual,
			)
		}
	}
}

// }}}

// PE0009 {{{
func TestPE0009(t *testing.T) {
	expected := 31875000
	if actual := PE0009(); actual != expected {
		t.Errorf(
			"Expected:%v\nActual:%v",
			expected,
			actual,
		)
	}
}

// }}}

// PE0010 {{{
var pe0010Cases = []struct {
	Input    int
	Expected int
}{
	{10, 17},
	{2000000, 142913828922},
}

func TestPE0010(t *testing.T) {
	for _, tc := range pe0010Cases {
		if actual := PE0010(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// }}}

// PE0011 {{{
var pe0011Cases = []struct {
	Expected int
}{
	{70600674},
}

func TestPE0011(t *testing.T) {
	for _, tc := range pe0011Cases {
		if actual := PE0011(); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// }}}

// PE0012 {{{
var pe0012Cases = []struct {
	Input    int
	Expected int
}{
	{5, 28},
	{500, 76576500},
}

func TestPE0012a(t *testing.T) {
	for _, tc := range pe0012Cases {
		if actual := PE0012a(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestPE0012b(t *testing.T) {
	for _, tc := range pe0012Cases {
		if actual := PE0012b(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// Benchmarks {{{
func BenchmarkPE0012a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0012a(500)
	}
}

func BenchmarkPE0012b(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0012b(500)
	}
}

// }}}
// }}}

// PE0013 {{{
var pe0013Cases = []struct {
	Input    []string
	Expected int
}{
	{[]string{
		"37107287533902102798797998220837590246510135740250",
		"46376937677490009712648124896970078050417018260538",
		"74324986199524741059474233309513058123726617309629",
		"91942213363574161572522430563301811072406154908250",
		"23067588207539346171171980310421047513778063246676",
		"89261670696623633820136378418383684178734361726757",
		"28112879812849979408065481931592621691275889832738",
		"44274228917432520321923589422876796487670272189318",
		"47451445736001306439091167216856844588711603153276",
		"70386486105843025439939619828917593665686757934951",
		"62176457141856560629502157223196586755079324193331",
		"64906352462741904929101432445813822663347944758178",
		"92575867718337217661963751590579239728245598838407",
		"58203565325359399008402633568948830189458628227828",
		"80181199384826282014278194139940567587151170094390",
		"35398664372827112653829987240784473053190104293586",
		"86515506006295864861532075273371959191420517255829",
		"71693888707715466499115593487603532921714970056938",
		"54370070576826684624621495650076471787294438377604",
		"53282654108756828443191190634694037855217779295145",
		"36123272525000296071075082563815656710885258350721",
		"45876576172410976447339110607218265236877223636045",
		"17423706905851860660448207621209813287860733969412",
		"81142660418086830619328460811191061556940512689692",
		"51934325451728388641918047049293215058642563049483",
		"62467221648435076201727918039944693004732956340691",
		"15732444386908125794514089057706229429197107928209",
		"55037687525678773091862540744969844508330393682126",
		"18336384825330154686196124348767681297534375946515",
		"80386287592878490201521685554828717201219257766954",
		"78182833757993103614740356856449095527097864797581",
		"16726320100436897842553539920931837441497806860984",
		"48403098129077791799088218795327364475675590848030",
		"87086987551392711854517078544161852424320693150332",
		"59959406895756536782107074926966537676326235447210",
		"69793950679652694742597709739166693763042633987085",
		"41052684708299085211399427365734116182760315001271",
		"65378607361501080857009149939512557028198746004375",
		"35829035317434717326932123578154982629742552737307",
		"94953759765105305946966067683156574377167401875275",
		"88902802571733229619176668713819931811048770190271",
		"25267680276078003013678680992525463401061632866526",
		"36270218540497705585629946580636237993140746255962",
		"24074486908231174977792365466257246923322810917141",
		"91430288197103288597806669760892938638285025333403",
		"34413065578016127815921815005561868836468420090470",
		"23053081172816430487623791969842487255036638784583",
		"11487696932154902810424020138335124462181441773470",
		"63783299490636259666498587618221225225512486764533",
		"67720186971698544312419572409913959008952310058822",
		"95548255300263520781532296796249481641953868218774",
		"76085327132285723110424803456124867697064507995236",
		"37774242535411291684276865538926205024910326572967",
		"23701913275725675285653248258265463092207058596522",
		"29798860272258331913126375147341994889534765745501",
		"18495701454879288984856827726077713721403798879715",
		"38298203783031473527721580348144513491373226651381",
		"34829543829199918180278916522431027392251122869539",
		"40957953066405232632538044100059654939159879593635",
		"29746152185502371307642255121183693803580388584903",
		"41698116222072977186158236678424689157993532961922",
		"62467957194401269043877107275048102390895523597457",
		"23189706772547915061505504953922979530901129967519",
		"86188088225875314529584099251203829009407770775672",
		"11306739708304724483816533873502340845647058077308",
		"82959174767140363198008187129011875491310547126581",
		"97623331044818386269515456334926366572897563400500",
		"42846280183517070527831839425882145521227251250327",
		"55121603546981200581762165212827652751691296897789",
		"32238195734329339946437501907836945765883352399886",
		"75506164965184775180738168837861091527357929701337",
		"62177842752192623401942399639168044983993173312731",
		"32924185707147349566916674687634660915035914677504",
		"99518671430235219628894890102423325116913619626622",
		"73267460800591547471830798392868535206946944540724",
		"76841822524674417161514036427982273348055556214818",
		"97142617910342598647204516893989422179826088076852",
		"87783646182799346313767754307809363333018982642090",
		"10848802521674670883215120185883543223812876952786",
		"71329612474782464538636993009049310363619763878039",
		"62184073572399794223406235393808339651327408011116",
		"66627891981488087797941876876144230030984490851411",
		"60661826293682836764744779239180335110989069790714",
		"85786944089552990653640447425576083659976645795096",
		"66024396409905389607120198219976047599490197230297",
		"64913982680032973156037120041377903785566085089252",
		"16730939319872750275468906903707539413042652315011",
		"94809377245048795150954100921645863754710598436791",
		"78639167021187492431995700641917969777599028300699",
		"15368713711936614952811305876380278410754449733078",
		"40789923115535562561142322423255033685442488917353",
		"44889911501440648020369068063960672322193204149535",
		"41503128880339536053299340368006977710650566631954",
		"81234880673210146739058568557934581403627822703280",
		"82616570773948327592232845941706525094512325230608",
		"22918802058777319719839450180888072429661980811197",
		"77158542502016545090413245809786882778948721859617",
		"72107838435069186155435662884062257473692284509516",
		"20849603980134001723930671666823555245252804609722",
		"53503534226472524250874054075591789781264330331690",
	}, 5537376230},
}

func TestPE0013(t *testing.T) {
	for _, tc := range pe0013Cases {
		if actual := PE0013(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual=%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

// }}}

// PE0014 {{{
var pe0014Cases = []struct {
	Input    int
	Expected int
}{
	{1000000, 837799},
}

func TestPE0014a(t *testing.T) {
	for _, tc := range pe0014Cases {
		if actual := PE0014a(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestPE0014b(t *testing.T) {
	for _, tc := range pe0014Cases {
		if actual := PE0014b(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// Benchmarks {{{
func BenchmarkPE0014a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0014a(b.N)
	}
}

func BenchmarkPE0014b(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PE0014b(b.N)
	}
}

// }}}
// }}}

// PE0015 {{{
var pe0015Cases = []struct {
	Input    int
	Expected int64
}{
	{2, 6},
	{3, 20},
	{20, 137846528820},
}

func TestPE0015(t *testing.T) {
	for _, tc := range pe0015Cases {
		if actual := PE0015(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestPE0015Dp(t *testing.T) {
	for _, tc := range pe0015Cases {
		if actual := PE0015Dp(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestPE0015Memoization(t *testing.T) {
	for _, tc := range pe0015Cases {
		if actual := PE0015Memoization(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// Benchmarks {{{
func BenchmarkPE0015(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0015(20)
	}
}

func BenchmarkPE0015Dp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0015Dp(20)
	}
}

func BenchmarkPE0015Memoization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0015Memoization(20)
	}
}

// }}}
// }}}

// PE0016 {{{
var pe0016Cases = []struct {
	Input    int
	Expected int
}{
	{15, 26},
	{1000, 1366},
}

func TestPE0016(t *testing.T) {
	for _, tc := range pe0016Cases {
		if actual := PE0016(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// }}}

// PE0017 {{{
var pe0017Cases = []struct {
	Input    int
	Expected int
}{
	{1000, 21124},
}

func TestPE0017SortSort(t *testing.T) {
	for _, tc := range pe0017Cases {
		actual, err := PE0017SortSort(tc.Input)
		if err != nil {
			t.Fatal(err)
		} else if actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestPE0017SortSlice(t *testing.T) {
	for _, tc := range pe0017Cases {
		actual, err := PE0017SortSort(tc.Input)
		if err != nil {
			t.Fatal(err)
		} else if actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// Benchmarks {{{
func BenchmarkPE0017SortSort(b *testing.B) {
	b.ResetTimer()
	PE0017SortSort(b.N)
}

func BenchmarkPE0017SortSlice(b *testing.B) {
	b.ResetTimer()
	PE0017SortSlice(b.N)
}

// }}}
// }}}

// PE0018 {{{
var pe0018Cases = []struct {
	Input0   int
	Input1   []string
	Expected int
}{
	{4, []string{
		"3",
		"7", "4",
		"2", "4", "6",
		"8", "5", "9", "3",
	}, 23},
	{15, []string{
		"75",
		"95", "64",
		"17", "47", "82",
		"18", "35", "87", "10",
		"20", "04", "82", "47", "65",
		"19", "01", "23", "75", "03", "34",
		"88", "02", "77", "73", "07", "63", "67",
		"99", "65", "04", "28", "06", "16", "70", "92",
		"41", "41", "26", "56", "83", "40", "80", "70", "33",
		"41", "48", "72", "33", "47", "32", "37", "16", "94", "29",
		"53", "71", "44", "65", "25", "43", "91", "52", "97", "51", "14",
		"70", "11", "33", "28", "77", "73", "17", "78", "39", "68", "17", "57",
		"91", "71", "52", "38", "17", "14", "91", "43", "58", "50", "27", "29", "48",
		"63", "66", "04", "68", "89", "53", "67", "30", "73", "16", "69", "87", "40", "31",
		"04", "62", "98", "27", "23", "09", "70", "98", "73", "93", "38", "53", "60", "04", "23",
	}, 1074},
}

func TestPE0018(t *testing.T) {
	for _, tc := range pe0018Cases {
		if actual := PE0018Memoization(tc.Input0, tc.Input1); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// }}}

// PE0019 {{{
var pe0019Cases = []struct {
	Input0    int
	Input1    int
	Expeceted int
}{
	{1904, 1904, 1},
	{1901, 2000, 171},
}

func TestPE0019(t *testing.T) {
	for _, tc := range pe0019Cases {
		if actual := PE0019(tc.Input0, tc.Input1); actual != tc.Expeceted {
			t.Errorf("expected=%v, actual=%v", tc.Expeceted, actual)
		}
	}
}

// }}}

// PE0020 {{{
var pe0020Cases = []struct {
	Input    int
	Expected int
}{
	{10, 27},
	{100, 648},
}

func TestPE0020(t *testing.T) {
	for _, tc := range pe0020Cases {
		if actual := PE0020(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// }}}

// PE0021 {{{
var pe0021Cases = []struct {
	Input    int
	Expected int
}{
	{-1, 0},
	{0, 0},
	{1, 0},
	{2, 0},
	{300, 504},
	{10000, 31626},
}

func TestPE0021a(t *testing.T) {
	for _, tc := range pe0021Cases {
		if actual := PE0021a(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

func TestPE0021Memoization(t *testing.T) {
	for _, tc := range pe0021Cases {
		if actual := PE0021Memoization(tc.Input); actual != tc.Expected {
			t.Errorf("expected=%v, actual=%v", tc.Expected, actual)
		}
	}
}

// Benchmarks {{{
func BenchmarkPE0021a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0021a(1000)
	}
}

func BenchmarkPE0021Memoization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0021Memoization(1000)
	}
}

// }}}
// }}}

// PE0022 {{{
var pe0022Cases = []struct {
	Input    string
	Expected int
}{
	{"p022_test.txt", 6*1 + 2*9 + 12*3},
	{"p022_names.txt", 871198282},
}

func TestPE0022(t *testing.T) {
	t.Skip("skip... test file is not included in repository")
	for _, tc := range pe0022Cases {
		if actual, _ := PE0022(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

// }}}

func Test_wordWorth(t *testing.T) {
	cases := []struct {
		Input    string
		Expected int
	}{
		{"COLIN", 53},
	}
	for _, tc := range cases {
		if actual := wordWorth(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input: %v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

// PE0023 {{{
func TestPE0023(t *testing.T) {
	if actual := PE0023(); actual != 4179871 {
		t.Errorf("Expected:%v\nActual:%v", 4179871, actual)
	}
}

// }}}

// PE0024 {{{
var pe0024Cases = []struct {
	Input    int
	Expected string
}{
	{1000000, "2783915460"},
}

func TestPE0024a(t *testing.T) {
	for _, tc := range pe0024Cases {
		if actual := PE0024a(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestPE0024b(t *testing.T) {
	for _, tc := range pe0024Cases {
		if actual := PE0024b(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

// }}}

// PE0025 {{{
/*
メモリの確保は矢張り重い
% go test -v ./project/euler -run PE0025 -bench PE0025 -benchmem
=== RUN   TestPE0025a
--- PASS: TestPE0025a (0.07s)
=== RUN   TestPE0025b
--- PASS: TestPE0025b (0.09s)
=== RUN   TestPE0025c
--- PASS: TestPE0025c (10.16s)
BenchmarkPE0025a-4            20          89606950 ns/op        11923335 B/op      46874 allocs/op
BenchmarkPE0025b-4            20          89584965 ns/op        11923821 B/op      46877 allocs/op
BenchmarkPE0025c-4             1        10095828900 ns/op       3396665360 B/op 23628161 allocs/op
PASS
ok      github.com/py0n/exercise/project/euler  25.186s
*/

var pe0025Cases = []struct {
	Input    int
	Expected int
}{
	{1000, 4782},
}

func TestPE0025a(t *testing.T) {
	for _, tc := range pe0025Cases {
		if actual := PE0025a(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestPE0025b(t *testing.T) {
	for _, tc := range pe0025Cases {
		if actual := PE0025b(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

func TestPE0025c(t *testing.T) {
	t.Skip()
	for _, tc := range pe0025Cases {
		if actual := PE0025c(tc.Input); actual != tc.Expected {
			t.Errorf(
				"Input:%v\nExpected:%v\nActual:%v",
				tc.Input,
				tc.Expected,
				actual,
			)
		}
	}
}

// Benchmarks {{{

func BenchmarkPE0025a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0025a(1000)
	}
}

func BenchmarkPE0025b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE0025b(1000)
	}
}

func BenchmarkPE0025c(b *testing.B) {
	b.Skip()
	for i := 0; i < b.N; i++ {
		PE0025c(1000)
	}
}

// }}}
// }}}

// vim:set foldmethod=marker:
