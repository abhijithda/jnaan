package minimumheighttrees

import (
	"reflect"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				n:     4,
				edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			},
			want: []int{1},
		},
		{
			name: "Example 2",
			args: args{
				n:     6,
				edges: [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}},
			},
			want: []int{3, 4},
		},
		{
			name: "Example 3 Bug",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			want: []int{0},
		},
		{
			name: "Example 4",
			args: args{
				n:     1,
				edges: [][]int{{}},
			},
			want: []int{0},
		},
		{
			name: "Example 5",
			args: args{
				n:     2,
				edges: [][]int{},
			},
			want: []int{0, 1},
		},
		{
			name: "Example 6 TLE",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {3, 4}, {4, 5}},
			},
			want: []int{3},
		},
		{
			name: "Example 7 TLE",
			args: args{
				n:     1616,
				edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}, {2, 5}, {1, 6}, {2, 7}, {7, 8}, {5, 9}, {9, 10}, {0, 11}, {8, 12}, {12, 13}, {3, 14}, {11, 15}, {8, 16}, {6, 17}, {10, 18}, {12, 19}, {5, 20}, {1, 21}, {21, 22}, {11, 23}, {2, 24}, {15, 25}, {6, 26}, {14, 27}, {13, 28}, {7, 29}, {16, 30}, {10, 31}, {8, 32}, {13, 33}, {22, 34}, {16, 35}, {34, 36}, {21, 37}, {19, 38}, {4, 39}, {34, 40}, {25, 41}, {8, 42}, {39, 43}, {21, 44}, {20, 45}, {1, 46}, {25, 47}, {33, 48}, {41, 49}, {26, 50}, {18, 51}, {1, 52}, {15, 53}, {11, 54}, {43, 55}, {9, 56}, {45, 57}, {21, 58}, {41, 59}, {40, 60}, {55, 61}, {58, 62}, {37, 63}, {34, 64}, {42, 65}, {47, 66}, {22, 67}, {18, 68}, {2, 69}, {20, 70}, {23, 71}, {36, 72}, {66, 73}, {62, 74}, {49, 75}, {67, 76}, {54, 77}, {58, 78}, {27, 79}, {39, 80}, {33, 81}, {71, 82}, {80, 83}, {21, 84}, {42, 85}, {1, 86}, {66, 87}, {71, 88}, {8, 89}, {23, 90}, {55, 91}, {47, 92}, {52, 93}, {43, 94}, {18, 95}, {11, 96}, {32, 97}, {11, 98}, {10, 99}, {10, 100}, {36, 101}, {67, 102}, {4, 103}, {72, 104}, {89, 105}, {82, 106}, {39, 107}, {59, 108}, {90, 109}, {89, 110}, {71, 111}, {38, 112}, {14, 113}, {65, 114}, {30, 115}, {99, 116}, {76, 117}, {110, 118}, {30, 119}, {9, 120}, {112, 121}, {87, 122}, {40, 123}, {94, 124}, {124, 125}, {72, 126}, {35, 127}, {45, 128}, {24, 129}, {73, 130}, {60, 131}, {60, 132}, {126, 133}, {100, 134}, {113, 135}, {65, 136}, {30, 137}, {78, 138}, {18, 139}, {34, 140}, {77, 141}, {41, 142}, {34, 143}, {53, 144}, {10, 145}, {94, 146}, {66, 147}, {101, 148}, {96, 149}, {11, 150}, {70, 151}, {51, 152}, {110, 153}, {37, 154}, {69, 155}, {109, 156}, {72, 157}, {112, 158}, {84, 159}, {32, 160}, {82, 161}, {36, 162}, {107, 163}, {99, 164}, {129, 165}, {115, 166}, {95, 167}, {76, 168}, {2, 169}, {140, 170}, {99, 171}, {146, 172}, {47, 173}, {20, 174}, {135, 175}, {69, 176}, {53, 177}, {97, 178}, {43, 179}, {102, 180}, {83, 181}, {8, 182}, {138, 183}, {120, 184}, {8, 185}, {89, 186}, {120, 187}, {24, 188}, {2, 189}, {87, 190}, {99, 191}, {176, 192}, {60, 193}, {58, 194}, {20, 195}, {10, 196}, {81, 197}, {159, 198}, {37, 199}, {180, 200}, {12, 201}, {145, 202}, {132, 203}, {5, 204}, {95, 205}, {80, 206}, {4, 207}, {125, 208}, {104, 209}, {32, 210}, {145, 211}, {41, 212}, {106, 213}, {113, 214}, {143, 215}, {76, 216}, {115, 217}, {66, 218}, {218, 219}, {39, 220}, {98, 221}, {32, 222}, {41, 223}, {136, 224}, {38, 225}, {49, 226}, {32, 227}, {107, 228}, {109, 229}, {77, 230}, {130, 231}, {94, 232}, {122, 233}, {201, 234}, {111, 235}, {85, 236}, {178, 237}, {220, 238}, {1, 239}, {127, 240}, {29, 241}, {193, 242}, {94, 243}, {120, 244}, {81, 245}, {9, 246}, {28, 247}, {93, 248}, {228, 249}, {133, 250}, {243, 251}, {26, 252}, {90, 253}, {194, 254}, {63, 255}, {170, 256}, {243, 257}, {26, 258}, {17, 259}, {41, 260}, {205, 261}, {181, 262}, {33, 263}, {213, 264}, {150, 265}, {152, 266}, {212, 267}, {56, 268}, {150, 269}, {153, 270}, {258, 271}, {47, 272}, {66, 273}, {23, 274}, {12, 275}, {82, 276}, {120, 277}, {35, 278}, {238, 279}, {44, 280}, {178, 281}, {158, 282}, {120, 283}, {181, 284}, {180, 285}, {283, 286}, {121, 287}, {34, 288}, {244, 289}, {143, 290}, {49, 291}, {267, 292}, {193, 293}, {162, 294}, {158, 295}, {166, 296}, {202, 297}, {57, 298}, {30, 299}, {225, 300}, {248, 301}, {199, 302}, {185, 303}, {194, 304}, {241, 305}, {147, 306}, {219, 307}, {150, 308}, {118, 309}, {125, 310}, {124, 311}, {47, 312}, {65, 313}, {255, 314}, {204, 315}, {196, 316}, {121, 317}, {277, 318}, {250, 319}, {299, 320}, {320, 321}, {49, 322}, {27, 323}, {115, 324}, {151, 325}, {17, 326}, {5, 327}, {77, 328}, {228, 329}, {297, 330}, {169, 331}, {147, 332}, {163, 333}, {158, 334}, {96, 335}, {155, 336}, {333, 337}, {137, 338}, {336, 339}, {30, 340}, {189, 341}, {169, 342}, {247, 343}, {274, 344}, {132, 345}, {15, 346}, {232, 347}, {51, 348}, {79, 349}, {141, 350}, {125, 351}, {174, 352}, {221, 353}, {62, 354}, {100, 355}, {336, 356}, {196, 357}, {332, 358}, {294, 359}, {324, 360}, {64, 361}, {196, 362}, {226, 363}, {12, 364}, {55, 365}, {19, 366}, {8, 367}, {143, 368}, {35, 369}, {63, 370}, {101, 371}, {129, 372}, {70, 373}, {82, 374}, {188, 375}, {328, 376}, {237, 377}, {304, 378}, {167, 379}, {6, 380}, {205, 381}, {316, 382}, {46, 383}, {76, 384}, {331, 385}, {382, 386}, {348, 387}, {250, 388}, {93, 389}, {35, 390}, {66, 391}, {82, 392}, {166, 393}, {264, 394}, {132, 395}, {315, 396}, {81, 397}, {239, 398}, {160, 399}, {306, 400}, {341, 401}, {265, 402}, {123, 403}, {243, 404}, {232, 405}, {139, 406}, {203, 407}, {14, 408}, {134, 409}, {402, 410}, {38, 411}, {146, 412}, {316, 413}, {276, 414}, {174, 415}, {204, 416}, {332, 417}, {182, 418}, {77, 419}, {165, 420}, {58, 421}, {349, 422}, {312, 423}, {121, 424}, {128, 425}, {276, 426}, {373, 427}, {383, 428}, {249, 429}, {39, 430}, {255, 431}, {231, 432}, {167, 433}, {418, 434}, {232, 435}, {316, 436}, {82, 437}, {216, 438}, {83, 439}, {31, 440}, {312, 441}, {76, 442}, {73, 443}, {169, 444}, {298, 445}, {224, 446}, {80, 447}, {111, 448}, {270, 449}, {350, 450}, {349, 451}, {256, 452}, {258, 453}, {75, 454}, {373, 455}, {430, 456}, {14, 457}, {209, 458}, {280, 459}, {110, 460}, {162, 461}, {431, 462}, {372, 463}, {146, 464}, {240, 465}, {27, 466}, {123, 467}, {324, 468}, {69, 469}, {28, 470}, {327, 471}, {384, 472}, {426, 473}, {25, 474}, {351, 475}, {252, 476}, {432, 477}, {86, 478}, {442, 479}, {92, 480}, {99, 481}, {213, 482}, {64, 483}, {97, 484}, {410, 485}, {273, 486}, {456, 487}, {156, 488}, {32, 489}, {154, 490}, {330, 491}, {115, 492}, {311, 493}, {87, 494}, {368, 495}, {264, 496}, {322, 497}, {152, 498}, {253, 499}, {106, 500}, {454, 501}, {356, 502}, {77, 503}, {452, 504}, {25, 505}, {61, 506}, {260, 507}, {395, 508}, {251, 509}, {45, 510}, {180, 511}, {78, 512}, {316, 513}, {294, 514}, {491, 515}, {191, 516}, {130, 517}, {9, 518}, {283, 519}, {95, 520}, {203, 521}, {514, 522}, {312, 523}, {18, 524}, {306, 525}, {343, 526}, {368, 527}, {45, 528}, {192, 529}, {217, 530}, {241, 531}, {224, 532}, {203, 533}, {118, 534}, {445, 535}, {264, 536}, {415, 537}, {0, 538}, {20, 539}, {427, 540}, {468, 541}, {346, 542}, {53, 543}, {318, 544}, {241, 545}, {367, 546}, {461, 547}, {22, 548}, {47, 549}, {526, 550}, {438, 551}, {260, 552}, {352, 553}, {328, 554}, {197, 555}, {359, 556}, {482, 557}, {457, 558}, {51, 559}, {244, 560}, {401, 561}, {89, 562}, {341, 563}, {305, 564}, {499, 565}, {194, 566}, {223, 567}, {248, 568}, {197, 569}, {285, 570}, {495, 571}, {298, 572}, {535, 573}, {79, 574}, {6, 575}, {188, 576}, {48, 577}, {394, 578}, {136, 579}, {150, 580}, {328, 581}, {207, 582}, {527, 583}, {415, 584}, {429, 585}, {44, 586}, {154, 587}, {166, 588}, {6, 589}, {529, 590}, {103, 591}, {434, 592}, {104, 593}, {239, 594}, {226, 595}, {349, 596}, {394, 597}, {457, 598}, {491, 599}, {351, 600}, {102, 601}, {152, 602}, {183, 603}, {30, 604}, {177, 605}, {288, 606}, {140, 607}, {146, 608}, {373, 609}, {553, 610}, {595, 611}, {359, 612}, {566, 613}, {123, 614}, {270, 615}, {395, 616}, {2, 617}, {414, 618}, {418, 619}, {447, 620}, {289, 621}, {42, 622}, {614, 623}, {106, 624}, {421, 625}, {197, 626}, {247, 627}, {437, 628}, {371, 629}, {82, 630}, {573, 631}, {393, 632}, {345, 633}, {283, 634}, {599, 635}, {469, 636}, {574, 637}, {248, 638}, {292, 639}, {165, 640}, {125, 641}, {563, 642}, {336, 643}, {473, 644}, {390, 645}, {356, 646}, {166, 647}, {229, 648}, {209, 649}, {510, 650}, {452, 651}, {469, 652}, {433, 653}, {48, 654}, {47, 655}, {625, 656}, {395, 657}, {117, 658}, {418, 659}, {367, 660}, {73, 661}, {368, 662}, {232, 663}, {651, 664}, {292, 665}, {603, 666}, {436, 667}, {552, 668}, {123, 669}, {528, 670}, {661, 671}, {121, 672}, {307, 673}, {315, 674}, {453, 675}, {29, 676}, {134, 677}, {406, 678}, {655, 679}, {160, 680}, {462, 681}, {318, 682}, {27, 683}, {326, 684}, {391, 685}, {106, 686}, {406, 687}, {650, 688}, {510, 689}, {347, 690}, {497, 691}, {292, 692}, {185, 693}, {259, 694}, {472, 695}, {177, 696}, {666, 697}, {391, 698}, {195, 699}, {424, 700}, {57, 701}, {286, 702}, {93, 703}, {17, 704}, {1, 705}, {427, 706}, {467, 707}, {545, 708}, {130, 709}, {254, 710}, {559, 711}, {465, 712}, {155, 713}, {540, 714}, {308, 715}, {646, 716}, {96, 717}, {117, 718}, {485, 719}, {618, 720}, {31, 721}, {26, 722}, {495, 723}, {28, 724}, {443, 725}, {621, 726}, {33, 727}, {118, 728}, {83, 729}, {255, 730}, {567, 731}, {232, 732}, {39, 733}, {501, 734}, {220, 735}, {81, 736}, {424, 737}, {383, 738}, {630, 739}, {246, 740}, {118, 741}, {145, 742}, {683, 743}, {611, 744}, {653, 745}, {294, 746}, {99, 747}, {747, 748}, {416, 749}, {719, 750}, {43, 751}, {388, 752}, {212, 753}, {165, 754}, {565, 755}, {191, 756}, {70, 757}, {151, 758}, {554, 759}, {414, 760}, {350, 761}, {415, 762}, {33, 763}, {39, 764}, {429, 765}, {723, 766}, {647, 767}, {543, 768}, {661, 769}, {61, 770}, {629, 771}, {48, 772}, {677, 773}, {58, 774}, {314, 775}, {211, 776}, {597, 777}, {761, 778}, {485, 779}, {2, 780}, {656, 781}, {360, 782}, {64, 783}, {381, 784}, {226, 785}, {340, 786}, {545, 787}, {697, 788}, {607, 789}, {301, 790}, {756, 791}, {582, 792}, {111, 793}, {724, 794}, {631, 795}, {423, 796}, {272, 797}, {45, 798}, {519, 799}, {444, 800}, {650, 801}, {236, 802}, {462, 803}, {531, 804}, {640, 805}, {433, 806}, {325, 807}, {293, 808}, {32, 809}, {19, 810}, {104, 811}, {655, 812}, {206, 813}, {615, 814}, {738, 815}, {520, 816}, {749, 817}, {811, 818}, {229, 819}, {175, 820}, {172, 821}, {810, 822}, {502, 823}, {458, 824}, {235, 825}, {660, 826}, {708, 827}, {422, 828}, {329, 829}, {570, 830}, {157, 831}, {430, 832}, {529, 833}, {19, 834}, {88, 835}, {168, 836}, {703, 837}, {474, 838}, {598, 839}, {124, 840}, {356, 841}, {608, 842}, {499, 843}, {678, 844}, {674, 845}, {791, 846}, {183, 847}, {762, 848}, {373, 849}, {613, 850}, {505, 851}, {379, 852}, {541, 853}, {737, 854}, {792, 855}, {397, 856}, {383, 857}, {450, 858}, {224, 859}, {310, 860}, {815, 861}, {420, 862}, {244, 863}, {409, 864}, {505, 865}, {520, 866}, {245, 867}, {437, 868}, {763, 869}, {686, 870}, {849, 871}, {740, 872}, {252, 873}, {501, 874}, {2, 875}, {171, 876}, {79, 877}, {233, 878}, {849, 879}, {823, 880}, {323, 881}, {661, 882}, {516, 883}, {6, 884}, {4, 885}, {176, 886}, {741, 887}, {284, 888}, {62, 889}, {667, 890}, {652, 891}, {842, 892}, {363, 893}, {243, 894}, {359, 895}, {288, 896}, {640, 897}, {540, 898}, {40, 899}, {27, 900}, {591, 901}, {368, 902}, {42, 903}, {738, 904}, {185, 905}, {705, 906}, {120, 907}, {310, 908}, {464, 909}, {681, 910}, {236, 911}, {530, 912}, {300, 913}, {361, 914}, {95, 915}, {53, 916}, {683, 917}, {720, 918}, {796, 919}, {506, 920}, {425, 921}, {772, 922}, {510, 923}, {149, 924}, {406, 925}, {510, 926}, {175, 927}, {456, 928}, {750, 929}, {633, 930}, {848, 931}, {459, 932}, {504, 933}, {404, 934}, {702, 935}, {818, 936}, {138, 937}, {572, 938}, {778, 939}, {576, 940}, {330, 941}, {547, 942}, {524, 943}, {31, 944}, {136, 945}, {633, 946}, {905, 947}, {738, 948}, {792, 949}, {772, 950}, {626, 951}, {28, 952}, {528, 953}, {656, 954}, {148, 955}, {705, 956}, {100, 957}, {6, 958}, {631, 959}, {340, 960}, {883, 961}, {887, 962}, {288, 963}, {263, 964}, {952, 965}, {306, 966}, {130, 967}, {615, 968}, {919, 969}, {896, 970}, {330, 971}, {259, 972}, {111, 973}, {866, 974}, {562, 975}, {659, 976}, {620, 977}, {381, 978}, {607, 979}, {322, 980}, {367, 981}, {580, 982}, {814, 983}, {117, 984}, {502, 985}, {826, 986}, {917, 987}, {689, 988}, {206, 989}, {807, 990}, {388, 991}, {425, 992}, {572, 993}, {143, 994}, {573, 995}, {492, 996}, {910, 997}, {713, 998}, {69, 999}, {797, 1000}, {738, 1001}, {259, 1002}, {951, 1003}, {748, 1004}, {979, 1005}, {371, 1006}, {270, 1007}, {354, 1008}, {290, 1009}, {184, 1010}, {560, 1011}, {746, 1012}, {955, 1013}, {832, 1014}, {633, 1015}, {330, 1016}, {840, 1017}, {649, 1018}, {614, 1019}, {885, 1020}, {896, 1021}, {684, 1022}, {285, 1023}, {671, 1024}, {995, 1025}, {613, 1026}, {549, 1027}, {519, 1028}, {853, 1029}, {280, 1030}, {620, 1031}, {265, 1032}, {546, 1033}, {914, 1034}, {797, 1035}, {647, 1036}, {274, 1037}, {328, 1038}, {681, 1039}, {175, 1040}, {0, 1041}, {1019, 1042}, {829, 1043}, {420, 1044}, {513, 1045}, {835, 1046}, {845, 1047}, {351, 1048}, {960, 1049}, {414, 1050}, {839, 1051}, {805, 1052}, {959, 1053}, {984, 1054}, {819, 1055}, {188, 1056}, {937, 1057}, {1000, 1058}, {443, 1059}, {946, 1060}, {122, 1061}, {871, 1062}, {91, 1063}, {207, 1064}, {473, 1065}, {12, 1066}, {314, 1067}, {205, 1068}, {423, 1069}, {155, 1070}, {557, 1071}, {207, 1072}, {40, 1073}, {338, 1074}, {714, 1075}, {802, 1076}, {346, 1077}, {1026, 1078}, {32, 1079}, {231, 1080}, {347, 1081}, {1064, 1082}, {750, 1083}, {1072, 1084}, {542, 1085}, {405, 1086}, {908, 1087}, {835, 1088}, {534, 1089}, {656, 1090}, {789, 1091}, {1011, 1092}, {335, 1093}, {787, 1094}, {82, 1095}, {714, 1096}, {471, 1097}, {759, 1098}, {270, 1099}, {136, 1100}, {792, 1101}, {982, 1102}, {782, 1103}, {965, 1104}, {1024, 1105}, {1065, 1106}, {395, 1107}, {134, 1108}, {479, 1109}, {565, 1110}, {152, 1111}, {1014, 1112}, {504, 1113}, {533, 1114}, {813, 1115}, {196, 1116}, {1084, 1117}, {597, 1118}, {477, 1119}, {546, 1120}, {732, 1121}, {337, 1122}, {395, 1123}, {320, 1124}, {428, 1125}, {915, 1126}, {868, 1127}, {55, 1128}, {160, 1129}, {647, 1130}, {305, 1131}, {926, 1132}, {1007, 1133}, {747, 1134}, {13, 1135}, {1094, 1136}, {249, 1137}, {720, 1138}, {752, 1139}, {1114, 1140}, {173, 1141}, {553, 1142}, {1051, 1143}, {473, 1144}, {765, 1145}, {646, 1146}, {25, 1147}, {480, 1148}, {129, 1149}, {397, 1150}, {700, 1151}, {829, 1152}, {611, 1153}, {384, 1154}, {709, 1155}, {539, 1156}, {896, 1157}, {199, 1158}, {721, 1159}, {977, 1160}, {941, 1161}, {685, 1162}, {1145, 1163}, {949, 1164}, {176, 1165}, {226, 1166}, {821, 1167}, {661, 1168}, {487, 1169}, {1152, 1170}, {557, 1171}, {588, 1172}, {766, 1173}, {682, 1174}, {737, 1175}, {291, 1176}, {334, 1177}, {925, 1178}, {337, 1179}, {1004, 1180}, {743, 1181}, {746, 1182}, {80, 1183}, {153, 1184}, {631, 1185}, {546, 1186}, {419, 1187}, {435, 1188}, {507, 1189}, {977, 1190}, {1113, 1191}, {221, 1192}, {510, 1193}, {916, 1194}, {234, 1195}, {91, 1196}, {307, 1197}, {474, 1198}, {1176, 1199}, {7, 1200}, {846, 1201}, {754, 1202}, {1061, 1203}, {1086, 1204}, {746, 1205}, {950, 1206}, {23, 1207}, {705, 1208}, {821, 1209}, {48, 1210}, {603, 1211}, {848, 1212}, {85, 1213}, {1122, 1214}, {717, 1215}, {954, 1216}, {240, 1217}, {875, 1218}, {372, 1219}, {87, 1220}, {588, 1221}, {736, 1222}, {1169, 1223}, {422, 1224}, {237, 1225}, {1065, 1226}, {1203, 1227}, {558, 1228}, {51, 1229}, {1100, 1230}, {1179, 1231}, {671, 1232}, {806, 1233}, {708, 1234}, {15, 1235}, {347, 1236}, {794, 1237}, {101, 1238}, {336, 1239}, {1231, 1240}, {4, 1241}, {430, 1242}, {1077, 1243}, {1224, 1244}, {419, 1245}, {54, 1246}, {1187, 1247}, {899, 1248}, {1110, 1249}, {1165, 1250}, {277, 1251}, {41, 1252}, {723, 1253}, {1143, 1254}, {440, 1255}, {353, 1256}, {118, 1257}, {105, 1258}, {1250, 1259}, {820, 1260}, {373, 1261}, {414, 1262}, {556, 1263}, {610, 1264}, {721, 1265}, {497, 1266}, {596, 1267}, {66, 1268}, {302, 1269}, {865, 1270}, {0, 1271}, {1216, 1272}, {398, 1273}, {830, 1274}, {1258, 1275}, {1092, 1276}, {843, 1277}, {269, 1278}, {935, 1279}, {600, 1280}, {882, 1281}, {835, 1282}, {544, 1283}, {508, 1284}, {33, 1285}, {578, 1286}, {1052, 1287}, {281, 1288}, {492, 1289}, {293, 1290}, {1071, 1291}, {779, 1292}, {5, 1293}, {982, 1294}, {1195, 1295}, {1155, 1296}, {965, 1297}, {949, 1298}, {455, 1299}, {37, 1300}, {433, 1301}, {908, 1302}, {640, 1303}, {593, 1304}, {438, 1305}, {366, 1306}, {849, 1307}, {289, 1308}, {1028, 1309}, {1059, 1310}, {1022, 1311}, {351, 1312}, {274, 1313}, {1031, 1314}, {836, 1315}, {291, 1316}, {879, 1317}, {466, 1318}, {786, 1319}, {755, 1320}, {582, 1321}, {1137, 1322}, {168, 1323}, {217, 1324}, {739, 1325}, {1, 1326}, {47, 1327}, {1239, 1328}, {1199, 1329}, {640, 1330}, {917, 1331}, {167, 1332}, {894, 1333}, {1240, 1334}, {130, 1335}, {1333, 1336}, {164, 1337}, {302, 1338}, {1115, 1339}, {1237, 1340}, {1222, 1341}, {710, 1342}, {446, 1343}, {825, 1344}, {152, 1345}, {794, 1346}, {1242, 1347}, {80, 1348}, {596, 1349}, {674, 1350}, {204, 1351}, {859, 1352}, {105, 1353}, {918, 1354}, {795, 1355}, {528, 1356}, {960, 1357}, {108, 1358}, {591, 1359}, {1138, 1360}, {881, 1361}, {167, 1362}, {214, 1363}, {288, 1364}, {97, 1365}, {1, 1366}, {1091, 1367}, {1185, 1368}, {386, 1369}, {11, 1370}, {319, 1371}, {1353, 1372}, {1237, 1373}, {367, 1374}, {954, 1375}, {172, 1376}, {133, 1377}, {312, 1378}, {647, 1379}, {614, 1380}, {82, 1381}, {968, 1382}, {232, 1383}, {1236, 1384}, {1025, 1385}, {262, 1386}, {247, 1387}, {398, 1388}, {407, 1389}, {552, 1390}, {1269, 1391}, {1334, 1392}, {1164, 1393}, {864, 1394}, {531, 1395}, {643, 1396}, {124, 1397}, {469, 1398}, {20, 1399}, {152, 1400}, {806, 1401}, {671, 1402}, {727, 1403}, {257, 1404}, {666, 1405}, {747, 1406}, {1171, 1407}, {892, 1408}, {129, 1409}, {655, 1410}, {652, 1411}, {781, 1412}, {877, 1413}, {452, 1414}, {287, 1415}, {114, 1416}, {1101, 1417}, {901, 1418}, {6, 1419}, {1154, 1420}, {660, 1421}, {357, 1422}, {750, 1423}, {602, 1424}, {460, 1425}, {584, 1426}, {697, 1427}, {1082, 1428}, {643, 1429}, {127, 1430}, {1356, 1431}, {823, 1432}, {492, 1433}, {507, 1434}, {1175, 1435}, {1, 1436}, {333, 1437}, {598, 1438}, {1382, 1439}, {556, 1440}, {1278, 1441}, {172, 1442}, {438, 1443}, {917, 1444}, {650, 1445}, {501, 1446}, {733, 1447}, {1066, 1448}, {317, 1449}, {185, 1450}, {1217, 1451}, {384, 1452}, {105, 1453}, {661, 1454}, {1270, 1455}, {1312, 1456}, {363, 1457}, {1363, 1458}, {319, 1459}, {132, 1460}, {1406, 1461}, {643, 1462}, {1436, 1463}, {1169, 1464}, {1175, 1465}, {897, 1466}, {963, 1467}, {1050, 1468}, {817, 1469}, {563, 1470}, {646, 1471}, {1281, 1472}, {1342, 1473}, {759, 1474}, {1468, 1475}, {557, 1476}, {1400, 1477}, {358, 1478}, {1272, 1479}, {1360, 1480}, {226, 1481}, {615, 1482}, {1239, 1483}, {231, 1484}, {780, 1485}, {257, 1486}, {1462, 1487}, {1309, 1488}, {59, 1489}, {1065, 1490}, {508, 1491}, {1345, 1492}, {1200, 1493}, {1114, 1494}, {12, 1495}, {603, 1496}, {544, 1497}, {167, 1498}, {471, 1499}, {110, 1500}, {1089, 1501}, {608, 1502}, {189, 1503}, {543, 1504}, {304, 1505}, {1169, 1506}, {669, 1507}, {215, 1508}, {397, 1509}, {679, 1510}, {363, 1511}, {1007, 1512}, {961, 1513}, {652, 1514}, {267, 1515}, {425, 1516}, {74, 1517}, {544, 1518}, {806, 1519}, {1508, 1520}, {800, 1521}, {649, 1522}, {223, 1523}, {629, 1524}, {585, 1525}, {227, 1526}, {298, 1527}, {448, 1528}, {237, 1529}, {502, 1530}, {227, 1531}, {1302, 1532}, {426, 1533}, {1188, 1534}, {310, 1535}, {1448, 1536}, {1122, 1537}, {1391, 1538}, {393, 1539}, {194, 1540}, {1381, 1541}, {551, 1542}, {637, 1543}, {227, 1544}, {407, 1545}, {88, 1546}, {636, 1547}, {1065, 1548}, {847, 1549}, {1270, 1550}, {232, 1551}, {169, 1552}, {465, 1553}, {332, 1554}, {1146, 1555}, {1516, 1556}, {1108, 1557}, {1082, 1558}, {382, 1559}, {877, 1560}, {1198, 1561}, {1480, 1562}, {376, 1563}, {1182, 1564}, {827, 1565}, {376, 1566}, {979, 1567}, {803, 1568}, {790, 1569}, {1339, 1570}, {202, 1571}, {574, 1572}, {1472, 1573}, {376, 1574}, {1380, 1575}, {305, 1576}, {1431, 1577}, {1504, 1578}, {640, 1579}, {1515, 1580}, {154, 1581}, {584, 1582}, {971, 1583}, {657, 1584}, {1439, 1585}, {162, 1586}, {949, 1587}, {194, 1588}, {498, 1589}, {831, 1590}, {194, 1591}, {573, 1592}, {638, 1593}, {288, 1594}, {356, 1595}, {1495, 1596}, {718, 1597}, {88, 1598}, {512, 1599}, {530, 1600}, {1354, 1601}, {887, 1602}, {908, 1603}, {491, 1604}, {733, 1605}, {245, 1606}, {373, 1607}, {413, 1608}, {614, 1609}, {112, 1610}, {350, 1611}, {411, 1612}, {216, 1613}, {1290, 1614}, {462, 1615}},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinHeightTrees(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinHeightTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
