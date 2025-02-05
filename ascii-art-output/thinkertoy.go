package main

func thinkertoy(char byte) []string {
	font := make(map[byte][]string, 100)
	font[32] = []string{
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	font[33] = []string{
		"  o   ",
		"  |   ",
		"  o   ",
		"     ",
		"  0  ",
		"      ",
	}
	font[34] = []string{
		" o o  ",
		" | |  ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	font[35] = []string{
		" | |  ",
		"-o-o- ",
		" | |  ",
		"-o-o- ",
		" | |  ",
		"      ",
	}
	font[36] = []string{
		"  | |  ",
		" -o o- ",
		"  | | ",
		" -o o- ",
		"  | |  ",
		"      ",
	}
	font[37] = []string{
		"      ",
		"     0 ",
		"o   /  ",
		"   /   ",
		"  /   o",
		" 0     ",
	}
	font[38] = []string{
		"    o ",
		"   / |",
		"      ",
		"  o-0-",
		"    | ",
		"      ",
	}
	font[39] = []string{
		" o    ",
		" |    ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	font[40] = []string{
		"    ",
		" /   ",
		"o   ",
		"|  ",
		"o   ",
		" __   ",
	}
	font[41] = []string{
		"   ",
		"__ ",
		"  o",
		"  |",
		"  o",
		" / ",
	}
	font[42] = []string{
		"o | o ",
		" L|/  ",
		"--O-- ",
		" /|L  ",
		"o | o ",
		"      ",
	}
	font[43] = []string{
		"  |  ",
		"  o  ",
		" -o- ",
		"  |  ",
		"     ",
		"     ",
	}
	font[44] = []string{
		"      ",
		"      ",
		"      ",
		"      ",
		"  o   ",
		"  |   ",
	}
	font[45] = []string{
		"          ",
		"          ",
		"          ",
		"   o-o    ",
		"          ",
		"          ",
	}
	font[46] = []string{
		"  ",
		"  ",
		"  ",
		"  ",
		"o ",
		"  ",
	}
	font[47] = []string{
		"        ",
		"      o ",
		"    o   ",
		"  o     ",
		" o      ",
		"        ",
	}
	font[48] = []string{
		"       ",
		"  o-o  ",
		" o   o ",
		" o   o ",
		" o   o ",
		"  o-o  ",
	}
	font[49] = []string{
		"     ",
		"  o  ",
		" /|  ",
		"o |  ",
		"  |  ",
		"o-o-o",
	}
	font[50] = []string{
		"        ",
		" ooooo  ",
		"o    o ",
		"   o   ",
		" o     ",
		"ooooo  ",
	}
	font[51] = []string{
		"       ",
		"ooooo  ",
		"    |  ",
		"oooo   ",
		"    |   ",
		"ooooo  ",
	}
	font[52] = []string{
		"       ",
		"o   o  ",
		"|   |  ",
		"oooooo ",
		"    |   ",
		"    o   ",
	}
	font[53] = []string{
		"       ",
		"oooooo ",
		"|      ",
		"ooooo  ",
		"    |   ",
		"ooooo  ",
	}
	font[54] = []string{
		"      ",
		"ooooo ",
		"|     ",
		"ooooo ",
		"|   | ",
		"ooooo ",
	}
	font[55] = []string{
		"       ",
		"oooooo ",
		"    /  ",
		"   o   ",
		"   |   ",
		"   o   ",
	}
	font[56] = []string{
		"       ",
		" o--o ",
		" |  |  ",
		" o--o ",
		" |  |  ",
		" o--o ",
	}
	font[57] = []string{
		"       ",
		" ooooo ",
		"|     |",
		" oooooo",
		"     / ",
		"ooooo  ",
	}
	font[58] = []string{
		"o ",
		"  ",
		"  ",
		"o ",
		"   ",
		"   ",
	}
	font[59] = []string{
		"   ",
		"o ",
		"  ",
		"  ",
		"o ",
		"| ",
	}
	font[60] = []string{
		"    ",
		"   o",
		" o  ",
		"o   ",
		" o  ",
		"  o ",
	}
	font[61] = []string{
		"    ",
		"    ",
		"    ",
		"o--o",
		"o--o",
		"    ",
		"    ",
	}
	font[62] = []string{
		"  ",
		"o ",
		" o ",
		"  o ",
		" o ",
		"o  ",
	}
	font[63] = []string{
		" oooo ",
		"o    o",
		"   o  ",
		"  o   ",
		"      ",
		"  o   ",
	}
	font[64] = []string{
		" oooo ",
		"o   o ",
		"o ooo ",
		"o o o ",
		" oooo ",
		"      ",
	}
	font[65] = []string{
		"  o   ",
		" o o  ",
		"ooooo ",
		"o   o ",
		"o   o ",
		"      ",
	}
	font[66] = []string{
		"ooooo  ",
		"o   o  ",
		"ooooo  ",
		"o   o  ",
		"ooooo  ",
		"       ",
	}
	font[67] = []string{
		" oooo ",
		"o    ",
		"o     ",
		"o    ",
		" oooo ",
		"      ",
	}
	font[68] = []string{
		"oooo   ",
		"o   o  ",
		"o   o  ",
		"o   o  ",
		"oooo   ",
		"       ",
	}
	font[69] = []string{
		"oooooo",
		"o     ",
		"oooo  ",
		"o     ",
		"oooooo",
		"      ",
	}
	font[70] = []string{
		"oooooo",
		"o     ",
		"oooo  ",
		"o     ",
		"o     ",
		"      ",
	}
	font[71] = []string{
		" oooo ",
		"o    o",
		"o     ",
		"o   oo",
		" oooo ",
		"      ",
	}
	font[72] = []string{
		"o   o ",
		"o   o ",
		"ooooo ",
		"o   o ",
		"o   o ",
		"      ",
	}
	font[73] = []string{
		" ooo ",
		"  o  ",
		"  o  ",
		"  o  ",
		" ooo ",
		"     ",
	}
	font[74] = []string{
		"    o ",
		"    o ",
		"    o ",
		"o   o ",
		" ooo  ",
		"      ",
	}
	font[75] = []string{
		"o   o ",
		"o  o  ",
		"ooo   ",
		"o  o  ",
		"o   o ",
		"      ",
	}
	font[76] = []string{
		"o     ",
		"o     ",
		"o     ",
		"o     ",
		"ooooo ",
		"      ",
	}
	font[77] = []string{
		"o   o ",
		"ooooo ",
		"o o o ",
		"o   o ",
		"o   o ",
		"      ",
	}
	font[78] = []string{
		"o   o ",
		"oo   o ",
		"o o  o ",
		"o  o o ",
		"o   oo ",
		"      ",
	}
	font[79] = []string{
		" ooo  ",
		"o   o ",
		"o   o ",
		"o   o ",
		" ooo  ",
		"      ",
	}
	font[80] = []string{
		"oooo  ",
		"o   o ",
		"oooo  ",
		"o     ",
		"o     ",
		"      ",
	}
	font[81] = []string{
		" ooo  ",
		"o   o ",
		"o   o ",
		"o  oo ",
		" ooo  ",
		"    o ",
	}
	font[82] = []string{
		"oooo  ",
		"o   o ",
		"oooo  ",
		"o  o  ",
		"o   o ",
		"      ",
	}
	font[83] = []string{
		" ooo  ",
		"o     ",
		" ooo  ",
		"    o ",
		" ooo  ",
		"      ",
	}
	font[84] = []string{
		"oooooo",
		"  o   ",
		"  o   ",
		"  o   ",
		"  o   ",
		"      ",
	}
	font[85] = []string{
		"o   o ",
		"o   o ",
		"o   o ",
		"o   o ",
		" ooo  ",
		"      ",
	}
	font[86] = []string{
		"o   o ",
		"o   o ",
		"o   o ",
		" o o  ",
		"  o   ",
		"      ",
	}
	font[87] = []string{
		"o   o ",
		"o   o ",
		"o o o ",
		"ooooo ",
		"o o o ",
		"      ",
	}
	font[88] = []string{
		"o   o ",
		" o o  ",
		"  o   ",
		" o o  ",
		"o   o ",
		"      ",
	}
	font[89] = []string{
		"o   o ",
		" o o  ",
		"  o   ",
		"  o   ",
		"  o   ",
		"      ",
	}
	font[90] = []string{
		"ooooo ",
		"   o  ",
		"  o   ",
		" o    ",
		"ooooo ",
		"      ",
	}
	font[91] = []string{
		"  ooo",
		"  o  ",
		"  o  ",
		"  o  ",
		"  ooo",
		"     ",
	}
	font[92] = []string{
		"o      ",
		" o     ",
		"  o    ",
		"   o   ",
		"    o  ",
		"     o ",
	}
	font[93] = []string{
		"ooo  ",
		"  o  ",
		"  o  ",
		"  o  ",
		"ooo  ",
		"     ",
	}
	font[94] = []string{
		"   o   ",
		"  o o  ",
		" o   o ",
		"o     o",
		"       ",
		"       ",
	}
	font[95] = []string{
		"        ",
		"        ",
		"        ",
		"        ",
		"oooooooo",
		"        ",
	}
	font[96] = []string{
		"  0   ",
		"  |   ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	font[97] = []string{
		"      ",
		"     ",
		"oooo  ",
		"o  o ",
		"oooo_  ",
		"      ",
	}
	font[98] = []string{
		"o     ",
		"o     ",
		"oooo  ",
		"o   o ",
		"oooo  ",
		"      ",
	}
	font[99] = []string{
		"      ",
		"      ",
		" ooo  ",
		"o     ",
		" ooo  ",
		"      ",
	}
	font[100] = []string{
		"    o ",
		"    o ",
		" oooo ",
		"o   o ",
		" oooo ",
		"      ",
	}
	font[101] = []string{
		"      ",
		"  ooo ",
		"ooooo ",
		"o     ",
		" ooo  ",
		"      ",
	}
	font[102] = []string{
		"  ooo ",
		" o    ",
		"oooo  ",
		"o     ",
		"o     ",
		"      ",
	}
	font[103] = []string{
		"      ",
		"oooo  ",
		"|  | ",
		"oooo  ",
		"   | ",
		"oooo  ",
	}
	font[104] = []string{
		"o     ",
		"o     ",
		"oooo  ",
		"o   o ",
		"o   o ",
		"      ",
	}
	font[105] = []string{
		"  o   ",
		"      ",
		" oo   ",
		"  o   ",
		" ooo  ",
		"      ",
	}
	font[106] = []string{
		"    o ",
		"      ",
		"   oo ",
		"    o ",
		"o   o ",
		" oo   ",
	}
	font[107] = []string{
		"o     ",
		"o     ",
		"o  o  ",
		"ooo   ",
		"o  o  ",
		"      ",
	}
	font[108] = []string{
		" oo   ",
		"  o   ",
		"  o   ",
		"  o   ",
		" ooo  ",
		"      ",
	}
	font[109] = []string{
		"      ",
		"      ",
		"ooooo ",
		"o o o ",
		"o o o ",
		"      ",
	}
	font[110] = []string{
		"      ",
		"      ",
		"oooo  ",
		"o   o ",
		"o   o ",
		"      ",
	}
	font[111] = []string{
		"      ",
		"      ",
		" ooo  ",
		"o   o ",
		" ooo  ",
		"      ",
	}
	font[112] = []string{
		"      ",
		"      ",
		"oooo  ",
		"o   o ",
		"oooo  ",
		"o     ",
	}
	font[113] = []string{
		"      ",
		"      ",
		" oooo ",
		"o   o ",
		" oooo ",
		"    o ",
	}
	font[114] = []string{
		"      ",
		"      ",
		"oooo  ",
		"o     ",
		"o     ",
		"      ",
	}
	font[115] = []string{
		"      ",
		"      ",
		" ooo  ",
		"   oo ",
		" ooo  ",
		"      ",
	}
	font[116] = []string{
		"  o   ",
		"  o   ",
		" ooo  ",
		"  o   ",
		"  oo  ",
		"      ",
	}
	font[117] = []string{
		"      ",
		"      ",
		"o   o ",
		"o   o ",
		" ooo  ",
		"      ",
	}
	font[118] = []string{
		"      ",
		"      ",
		"o   o ",
		"o   o ",
		" o o  ",
		"  o   ",
	}
	font[119] = []string{
		"      ",
		"      ",
		"o o o ",
		"o o o ",
		"ooooo ",
		"      ",
	}
	font[120] = []string{
		"      ",
		"      ",
		"o   o ",
		" o o  ",
		"o   o ",
		"      ",
	}
	font[121] = []string{
		"      ",
		"      ",
		"o   o ",
		"o   o ",
		" ooo  ",
		"    o ",
	}
	font[122] = []string{
		"      ",
		"      ",
		"ooooo ",
		"  /   ",
		"ooooo ",
		"      ",
	}
	font[123] = []string{
		"   o  ",
		"  o   ",
		" oo   ",
		"  o   ",
		"   o  ",
		"    o ",
	}
	font[124] = []string{
		"  o   ",
		"  o   ",
		"  o   ",
		"  o   ",
		"  o   ",
		"  o   ",
	}
	font[125] = []string{
		"o     ",
		" o    ",
		"  oo  ",
		"   o  ",
		"    o ",
		"o     ",
	}
	font[126] = []string{
		"     ",
		"o o o",
		" o o ",
		"     ",
		"     ",
		"     ",
	}
	if v, ok := font[char]; ok {
		return v
	}
	return font[32]
}
