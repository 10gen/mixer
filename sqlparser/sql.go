//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	tuple       Tuple
	valExprs    ValExprs
	values      Values
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const ASC = 57368
const DESC = 57369
const VALUES = 57370
const INTO = 57371
const DUPLICATE = 57372
const KEY = 57373
const DEFAULT = 57374
const SET = 57375
const LOCK = 57376
const ID = 57377
const STRING = 57378
const NUMBER = 57379
const VALUE_ARG = 57380
const COMMENT = 57381
const LE = 57382
const GE = 57383
const NE = 57384
const NULL_SAFE_EQUAL = 57385
const UNION = 57386
const MINUS = 57387
const EXCEPT = 57388
const INTERSECT = 57389
const JOIN = 57390
const STRAIGHT_JOIN = 57391
const LEFT = 57392
const RIGHT = 57393
const INNER = 57394
const OUTER = 57395
const CROSS = 57396
const NATURAL = 57397
const USE = 57398
const FORCE = 57399
const ON = 57400
const AND = 57401
const OR = 57402
const NOT = 57403
const UNARY = 57404
const CASE = 57405
const WHEN = 57406
const THEN = 57407
const ELSE = 57408
const END = 57409
const BEGIN = 57410
const COMMIT = 57411
const ROLLBACK = 57412
const NAMES = 57413
const REPLACE = 57414
const ADMIN = 57415
const SHOW = 57416
const DATABASES = 57417
const TABLES = 57418
const PROXY = 57419
const VARIABLES = 57420
const CREATE = 57421
const ALTER = 57422
const DROP = 57423
const RENAME = 57424
const TABLE = 57425
const INDEX = 57426
const VIEW = 57427
const TO = 57428
const IGNORE = 57429
const IF = 57430
const UNIQUE = 57431
const USING = 57432

var yyToknames = []string{
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"AND",
	"OR",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"BEGIN",
	"COMMIT",
	"ROLLBACK",
	"NAMES",
	"REPLACE",
	"ADMIN",
	"SHOW",
	"DATABASES",
	"TABLES",
	"PROXY",
	"VARIABLES",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 216
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 592

var yyAct = []int{

	111, 317, 148, 384, 71, 102, 352, 187, 108, 138,
	271, 227, 309, 262, 264, 119, 197, 73, 210, 109,
	393, 202, 97, 98, 393, 393, 157, 188, 3, 89,
	162, 163, 85, 60, 285, 286, 287, 288, 289, 78,
	290, 291, 278, 141, 75, 315, 49, 80, 50, 157,
	82, 363, 74, 157, 86, 225, 44, 225, 46, 362,
	95, 62, 47, 34, 35, 36, 37, 103, 52, 53,
	54, 254, 118, 395, 361, 124, 79, 394, 392, 342,
	137, 81, 76, 115, 116, 117, 338, 92, 145, 51,
	93, 150, 140, 343, 149, 122, 147, 152, 314, 155,
	76, 159, 304, 263, 154, 216, 302, 296, 255, 189,
	224, 185, 186, 190, 334, 336, 161, 134, 120, 121,
	256, 129, 136, 151, 214, 125, 358, 217, 196, 75,
	193, 310, 75, 200, 274, 206, 205, 74, 162, 163,
	74, 56, 58, 59, 57, 207, 61, 144, 68, 204,
	123, 72, 328, 155, 335, 220, 263, 329, 307, 244,
	103, 233, 206, 221, 175, 176, 177, 237, 235, 236,
	242, 243, 231, 246, 247, 248, 249, 250, 251, 252,
	253, 232, 131, 223, 234, 238, 153, 213, 215, 212,
	162, 163, 326, 310, 103, 103, 84, 327, 360, 75,
	75, 245, 359, 267, 332, 345, 331, 74, 269, 273,
	330, 275, 131, 203, 258, 260, 127, 225, 203, 130,
	369, 276, 270, 75, 347, 266, 133, 280, 281, 230,
	156, 74, 34, 35, 36, 37, 279, 146, 229, 222,
	379, 295, 378, 298, 299, 231, 198, 18, 282, 266,
	199, 87, 61, 259, 377, 114, 283, 199, 150, 297,
	118, 131, 103, 124, 173, 174, 175, 176, 177, 308,
	101, 115, 116, 117, 157, 306, 194, 230, 192, 106,
	313, 316, 303, 122, 340, 312, 229, 170, 171, 172,
	173, 174, 175, 176, 177, 324, 325, 191, 96, 231,
	231, 126, 105, 341, 294, 390, 120, 121, 99, 76,
	344, 339, 337, 125, 321, 320, 75, 160, 349, 219,
	293, 350, 353, 391, 348, 285, 286, 287, 288, 289,
	354, 290, 291, 61, 368, 218, 201, 69, 123, 142,
	139, 257, 135, 364, 132, 83, 128, 366, 365, 170,
	171, 172, 173, 174, 175, 176, 177, 346, 88, 67,
	155, 18, 374, 301, 376, 375, 373, 367, 397, 208,
	143, 381, 353, 65, 63, 383, 382, 318, 385, 385,
	385, 75, 386, 387, 265, 388, 18, 114, 357, 74,
	319, 272, 118, 323, 398, 124, 356, 203, 399, 90,
	400, 114, 101, 115, 116, 117, 118, 94, 396, 124,
	38, 106, 91, 70, 380, 122, 76, 115, 116, 117,
	239, 18, 240, 241, 39, 106, 17, 16, 15, 122,
	40, 41, 42, 43, 105, 14, 13, 12, 120, 121,
	99, 55, 209, 45, 277, 125, 211, 48, 105, 114,
	77, 268, 120, 121, 118, 18, 389, 124, 370, 125,
	351, 355, 322, 305, 76, 115, 116, 117, 195, 261,
	123, 113, 110, 106, 112, 118, 311, 122, 124, 107,
	18, 19, 20, 21, 123, 76, 115, 116, 117, 164,
	165, 169, 167, 168, 150, 104, 105, 333, 122, 228,
	120, 121, 284, 371, 372, 226, 100, 125, 22, 181,
	182, 183, 184, 292, 178, 179, 180, 158, 64, 33,
	66, 120, 121, 11, 10, 9, 8, 7, 125, 6,
	5, 4, 123, 2, 1, 0, 166, 170, 171, 172,
	173, 174, 175, 176, 177, 170, 171, 172, 173, 174,
	175, 176, 177, 123, 0, 0, 0, 0, 27, 28,
	29, 0, 30, 32, 31, 0, 0, 0, 0, 23,
	24, 26, 25, 300, 0, 0, 170, 171, 172, 173,
	174, 175, 176, 177, 170, 171, 172, 173, 174, 175,
	176, 177,
}
var yyPact = []int{

	475, -1000, -1000, 183, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -42, -54, -9, -30, -1000, -1000, -1000,
	-1000, 51, 217, 416, 357, -1000, -1000, -1000, 355, -1000,
	330, 302, 404, 65, -64, -23, 217, -1000, -17, 217,
	-1000, 310, -71, 217, -71, 329, 389, 389, 398, 217,
	254, -1000, -1000, -1000, 367, -1000, 262, 302, 313, 45,
	302, 159, 309, -1000, 181, -1000, 41, 307, 55, 217,
	-1000, 305, -1000, -58, 304, 350, 83, 217, 302, -1000,
	429, 47, -1000, 389, 47, 398, 47, 221, -1000, -1000,
	298, 40, 73, 469, -1000, 429, 381, -1000, -1000, -1000,
	47, 253, 234, -1000, 232, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 47, -1000, 213, 274, 301,
	387, 274, -1000, 47, 217, -1000, 349, -87, -1000, 92,
	-1000, 300, -1000, -1000, 284, -1000, 206, 73, 469, 516,
	450, -1000, 516, 389, 4, 516, 194, 367, -1000, -1000,
	217, 111, 429, 429, 47, 214, 399, 47, 47, 134,
	47, 47, 47, 47, 47, 47, 47, 47, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -35, 2, 14, 469,
	-1000, 235, 367, -1000, 416, 24, 516, 356, 274, 274,
	208, -1000, 378, 429, -1000, 516, -1000, -1000, -1000, 70,
	217, -1000, -59, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 356, 274, -1000, -1000, 47, 203, 271, 285, 242,
	31, -1000, -1000, -1000, -1000, -1000, -1000, 516, -1000, 214,
	47, 47, 516, 508, -1000, 338, 193, 193, 193, 91,
	91, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 0, 367,
	-4, 77, -1000, 429, 67, 214, 183, 129, -8, -1000,
	378, 362, 376, 73, 280, -1000, -1000, 279, -1000, -1000,
	159, 516, 382, 194, 194, -1000, -1000, 138, 98, 156,
	152, 150, 52, -1000, 277, -20, 276, -1000, 516, 219,
	47, -1000, -1000, -27, -1000, 11, -1000, 47, 125, -1000,
	327, 171, -1000, -1000, -1000, 274, 362, -1000, 47, 47,
	-1000, -1000, 384, 374, 271, 62, -1000, 148, -1000, 144,
	-1000, -1000, -1000, -1000, -25, -40, -48, -1000, -1000, -1000,
	47, 516, -1000, -1000, 516, 47, 316, 214, -1000, -1000,
	281, 167, -1000, 477, -1000, 378, 429, 47, 429, -1000,
	-1000, 210, 198, 196, 516, 516, 407, -1000, 47, 47,
	-1000, -1000, -1000, 362, 73, 164, 73, 217, 217, 217,
	274, 516, -1000, 289, -28, -1000, -29, -33, 159, -1000,
	401, 347, -1000, 217, -1000, -1000, -1000, 217, -1000, 217,
	-1000,
}
var yyPgo = []int{

	0, 534, 533, 27, 531, 530, 529, 527, 526, 525,
	524, 523, 410, 520, 519, 518, 22, 23, 517, 513,
	506, 505, 11, 502, 499, 148, 497, 3, 21, 5,
	495, 489, 14, 479, 2, 19, 7, 476, 474, 15,
	472, 8, 471, 469, 13, 468, 463, 462, 461, 10,
	460, 6, 458, 1, 456, 16, 451, 12, 4, 17,
	196, 450, 447, 446, 444, 443, 442, 0, 9, 437,
	436, 435, 428, 427, 426, 90, 29, 424,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 69, 70,
	71, 74, 73, 73, 73, 73, 8, 8, 8, 9,
	9, 9, 10, 11, 11, 11, 77, 12, 13, 13,
	14, 14, 14, 14, 14, 15, 15, 16, 16, 17,
	17, 17, 20, 20, 18, 18, 18, 21, 21, 22,
	22, 22, 22, 19, 19, 19, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 24, 24, 24, 25, 25,
	26, 26, 26, 26, 27, 27, 28, 28, 76, 76,
	76, 75, 75, 29, 29, 29, 29, 29, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 31, 31,
	31, 31, 31, 31, 31, 32, 32, 37, 37, 35,
	35, 39, 36, 36, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 38, 38, 40, 40, 40, 42, 45, 45, 43,
	43, 44, 46, 46, 41, 41, 33, 33, 33, 33,
	47, 47, 48, 48, 49, 49, 50, 50, 51, 52,
	52, 52, 53, 53, 53, 54, 54, 54, 55, 55,
	56, 56, 57, 57, 58, 58, 59, 60, 60, 61,
	61, 62, 62, 63, 63, 63, 63, 63, 64, 64,
	65, 65, 66, 66, 67, 68,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 1,
	1, 5, 3, 3, 4, 5, 5, 8, 4, 6,
	7, 4, 5, 4, 5, 5, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 0, 2,
	2, 0, 2, 1, 3, 3, 2, 3, 3, 3,
	4, 3, 4, 5, 6, 3, 4, 2, 1, 1,
	1, 1, 1, 1, 1, 2, 1, 1, 3, 3,
	1, 3, 1, 3, 1, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 3, 4, 5, 4,
	1, 1, 1, 1, 1, 1, 5, 0, 1, 1,
	2, 4, 0, 2, 1, 3, 1, 1, 1, 1,
	0, 3, 0, 2, 0, 3, 1, 3, 2, 0,
	1, 1, 0, 2, 4, 0, 2, 4, 0, 3,
	1, 3, 0, 5, 1, 3, 3, 0, 2, 0,
	3, 0, 1, 1, 1, 1, 1, 1, 0, 1,
	0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 94, 95, 97, 96, 83, 84, 85,
	87, 89, 88, -14, 49, 50, 51, 52, -12, -77,
	-12, -12, -12, -12, 98, -65, 100, 104, -62, 100,
	102, 98, 98, 99, 100, -12, 90, 93, 91, 92,
	-67, 35, -3, 17, -15, 18, -13, 29, -25, 35,
	9, -58, 86, -59, -41, -67, 35, -61, 103, 99,
	-67, 98, -67, 35, -60, 103, -67, -60, 29, -76,
	10, 23, -76, -75, 9, -67, 44, -16, -17, 73,
	-20, 35, -29, -34, -30, 67, 44, -33, -41, -35,
	-40, -67, -38, -42, 20, 36, 37, 38, 25, -39,
	71, 72, 48, 103, 28, 78, 39, -25, 33, 76,
	-25, 53, 35, 45, 76, 35, 67, -67, -68, 35,
	-68, 101, 35, 20, 64, -67, -25, -29, -34, -34,
	44, -76, -34, -75, -36, -34, 9, 53, -18, -67,
	19, 76, 65, 66, -31, 21, 67, 23, 24, 22,
	68, 69, 70, 71, 72, 73, 74, 75, 45, 46,
	47, 40, 41, 42, 43, -29, -29, -36, -3, -34,
	-34, 44, 44, -39, 44, -45, -34, -55, 33, 44,
	-58, 35, -28, 10, -59, -34, -67, -68, 20, -66,
	105, -63, 97, 95, 32, 96, 13, 35, 35, 35,
	-68, -55, 33, -76, 106, 53, -21, -22, -24, 44,
	35, -39, -17, -67, 73, -29, -29, -34, -35, 21,
	23, 24, -34, -34, 25, 67, -34, -34, -34, -34,
	-34, -34, -34, -34, 106, 106, 106, 106, -16, 18,
	-16, -43, -44, 79, -32, 28, -3, -58, -56, -41,
	-28, -49, 13, -29, 64, -67, -68, -64, 101, -32,
	-58, -34, -28, 53, -23, 54, 55, 56, 57, 58,
	60, 61, -19, 35, 19, -22, 76, -35, -34, -34,
	65, 25, 106, -16, 106, -46, -44, 81, -29, -57,
	64, -37, -35, -57, 106, 53, -49, -53, 15, 14,
	35, 35, -47, 11, -22, -22, 54, 59, 54, 59,
	54, 54, 54, -26, 62, 102, 63, 35, 106, 35,
	65, -34, 106, 82, -34, 80, 30, 53, -41, -53,
	-34, -50, -51, -34, -68, -48, 12, 14, 64, 54,
	54, 99, 99, 99, -34, -34, 31, -35, 53, 53,
	-52, 26, 27, -49, -29, -36, -29, 44, 44, 44,
	7, -34, -51, -53, -27, -67, -27, -27, -58, -54,
	16, 34, 106, 53, 106, 106, 7, 21, -67, -67,
	-67,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 46, 46,
	46, 46, 46, 210, 201, 0, 0, 28, 29, 30,
	46, 0, 0, 0, 50, 52, 53, 54, 55, 48,
	0, 0, 0, 0, 199, 0, 0, 211, 0, 0,
	202, 0, 197, 0, 197, 0, 98, 98, 101, 0,
	0, 214, 19, 51, 0, 56, 47, 0, 0, 88,
	0, 26, 0, 194, 0, 164, 214, 0, 0, 0,
	215, 0, 215, 0, 0, 0, 0, 0, 0, 32,
	0, 0, 33, 98, 0, 101, 0, 17, 57, 59,
	64, 214, 62, 63, 103, 0, 0, 134, 135, 136,
	0, 164, 0, 150, 0, 166, 167, 168, 169, 130,
	153, 154, 155, 151, 152, 157, 49, 188, 0, 0,
	96, 0, 27, 0, 0, 215, 0, 212, 38, 0,
	41, 0, 43, 198, 0, 215, 188, 99, 0, 100,
	0, 34, 102, 98, 0, 132, 0, 0, 60, 65,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 118, 119,
	120, 121, 122, 123, 124, 106, 0, 0, 0, 132,
	145, 0, 0, 117, 0, 0, 158, 0, 0, 0,
	96, 89, 174, 0, 195, 196, 165, 36, 200, 0,
	0, 215, 208, 203, 204, 205, 206, 207, 42, 44,
	45, 0, 0, 35, 31, 0, 96, 67, 73, 0,
	85, 87, 58, 66, 61, 104, 105, 108, 109, 0,
	0, 0, 111, 0, 115, 0, 137, 138, 139, 140,
	141, 142, 143, 144, 107, 129, 131, 146, 0, 0,
	0, 162, 159, 0, 192, 0, 126, 192, 0, 190,
	174, 182, 0, 97, 0, 213, 39, 0, 209, 22,
	23, 133, 170, 0, 0, 76, 77, 0, 0, 0,
	0, 0, 90, 74, 0, 0, 0, 110, 112, 0,
	0, 116, 147, 0, 149, 0, 160, 0, 0, 20,
	0, 125, 127, 21, 189, 0, 182, 25, 0, 0,
	215, 40, 172, 0, 68, 71, 78, 0, 80, 0,
	82, 83, 84, 69, 0, 0, 0, 75, 70, 86,
	0, 113, 148, 156, 163, 0, 0, 0, 191, 24,
	183, 175, 176, 179, 37, 174, 0, 0, 0, 79,
	81, 0, 0, 0, 114, 161, 0, 128, 0, 0,
	178, 180, 181, 182, 173, 171, 72, 0, 0, 0,
	0, 184, 177, 185, 0, 94, 0, 0, 193, 18,
	0, 0, 91, 0, 92, 93, 186, 0, 95, 0,
	187,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 106, 73, 71, 53, 72, 76, 74, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	46, 45, 47, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 70, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 69, 3, 48,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 49, 50, 51, 52, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	77, 78, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line sql.y:170
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:176
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		yyVAL.statement = yyS[yypt-0].statement
	case 4:
		yyVAL.statement = yyS[yypt-0].statement
	case 5:
		yyVAL.statement = yyS[yypt-0].statement
	case 6:
		yyVAL.statement = yyS[yypt-0].statement
	case 7:
		yyVAL.statement = yyS[yypt-0].statement
	case 8:
		yyVAL.statement = yyS[yypt-0].statement
	case 9:
		yyVAL.statement = yyS[yypt-0].statement
	case 10:
		yyVAL.statement = yyS[yypt-0].statement
	case 11:
		yyVAL.statement = yyS[yypt-0].statement
	case 12:
		yyVAL.statement = yyS[yypt-0].statement
	case 13:
		yyVAL.statement = yyS[yypt-0].statement
	case 14:
		yyVAL.statement = yyS[yypt-0].statement
	case 15:
		yyVAL.statement = yyS[yypt-0].statement
	case 16:
		yyVAL.statement = yyS[yypt-0].statement
	case 17:
		//line sql.y:196
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 18:
		//line sql.y:200
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 19:
		//line sql.y:204
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 20:
		//line sql.y:211
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:215
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 22:
		//line sql.y:227
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 23:
		//line sql.y:231
		{
			cols := make(Columns, 0, len(yyS[yypt-0].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-0].updateExprs))
			for _, col := range yyS[yypt-0].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 24:
		//line sql.y:244
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 25:
		//line sql.y:250
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 26:
		//line sql.y:256
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 27:
		//line sql.y:260
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 28:
		//line sql.y:266
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		//line sql.y:272
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		//line sql.y:278
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		//line sql.y:284
		{
			yyVAL.statement = &Admin{Name: yyS[yypt-3].bytes, Values: yyS[yypt-1].valExprs}
		}
	case 32:
		//line sql.y:290
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyS[yypt-0].expr}
		}
	case 33:
		//line sql.y:294
		{
			yyVAL.statement = &Show{Section: "variables", LikeOrWhere: yyS[yypt-0].expr}
		}
	case 34:
		//line sql.y:298
		{
			yyVAL.statement = &Show{Section: "tables", From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 35:
		//line sql.y:302
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyS[yypt-2].bytes), From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 36:
		//line sql.y:308
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 37:
		//line sql.y:312
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 38:
		//line sql.y:317
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 39:
		//line sql.y:323
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 40:
		//line sql.y:327
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 41:
		//line sql.y:332
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 42:
		//line sql.y:338
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 43:
		//line sql.y:344
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 44:
		//line sql.y:348
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 45:
		//line sql.y:353
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 46:
		//line sql.y:358
		{
			SetAllowComments(yylex, true)
		}
	case 47:
		//line sql.y:362
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 48:
		//line sql.y:368
		{
			yyVAL.bytes2 = nil
		}
	case 49:
		//line sql.y:372
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 50:
		//line sql.y:378
		{
			yyVAL.str = AST_UNION
		}
	case 51:
		//line sql.y:382
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 52:
		//line sql.y:386
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 53:
		//line sql.y:390
		{
			yyVAL.str = AST_EXCEPT
		}
	case 54:
		//line sql.y:394
		{
			yyVAL.str = AST_INTERSECT
		}
	case 55:
		//line sql.y:399
		{
			yyVAL.str = ""
		}
	case 56:
		//line sql.y:403
		{
			yyVAL.str = AST_DISTINCT
		}
	case 57:
		//line sql.y:409
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 58:
		//line sql.y:413
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 59:
		//line sql.y:419
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 60:
		//line sql.y:423
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 61:
		//line sql.y:427
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 62:
		//line sql.y:433
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 63:
		//line sql.y:437
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 64:
		//line sql.y:442
		{
			yyVAL.bytes = nil
		}
	case 65:
		//line sql.y:446
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 66:
		//line sql.y:450
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 67:
		//line sql.y:456
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 68:
		//line sql.y:460
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 69:
		//line sql.y:466
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 70:
		//line sql.y:470
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 71:
		//line sql.y:474
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 72:
		//line sql.y:478
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 73:
		//line sql.y:483
		{
			yyVAL.bytes = nil
		}
	case 74:
		//line sql.y:487
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 75:
		//line sql.y:491
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 76:
		//line sql.y:497
		{
			yyVAL.str = AST_JOIN
		}
	case 77:
		//line sql.y:501
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 78:
		//line sql.y:505
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 79:
		//line sql.y:509
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 80:
		//line sql.y:513
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 81:
		//line sql.y:517
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 82:
		//line sql.y:521
		{
			yyVAL.str = AST_JOIN
		}
	case 83:
		//line sql.y:525
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 84:
		//line sql.y:529
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 85:
		//line sql.y:535
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 86:
		//line sql.y:539
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 87:
		//line sql.y:543
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 88:
		//line sql.y:549
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 89:
		//line sql.y:553
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 90:
		//line sql.y:558
		{
			yyVAL.indexHints = nil
		}
	case 91:
		//line sql.y:562
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 92:
		//line sql.y:566
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 93:
		//line sql.y:570
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 94:
		//line sql.y:576
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 95:
		//line sql.y:580
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 96:
		//line sql.y:585
		{
			yyVAL.boolExpr = nil
		}
	case 97:
		//line sql.y:589
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 98:
		//line sql.y:594
		{
			yyVAL.expr = nil
		}
	case 99:
		//line sql.y:598
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 100:
		//line sql.y:602
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 101:
		//line sql.y:607
		{
			yyVAL.valExpr = nil
		}
	case 102:
		//line sql.y:611
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 103:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 104:
		//line sql.y:618
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 105:
		//line sql.y:622
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 106:
		//line sql.y:626
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 107:
		//line sql.y:630
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 108:
		//line sql.y:636
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 109:
		//line sql.y:640
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 110:
		//line sql.y:644
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 111:
		//line sql.y:648
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 112:
		//line sql.y:652
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 113:
		//line sql.y:656
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 114:
		//line sql.y:660
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 115:
		//line sql.y:664
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 116:
		//line sql.y:668
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 117:
		//line sql.y:672
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 118:
		//line sql.y:678
		{
			yyVAL.str = AST_EQ
		}
	case 119:
		//line sql.y:682
		{
			yyVAL.str = AST_LT
		}
	case 120:
		//line sql.y:686
		{
			yyVAL.str = AST_GT
		}
	case 121:
		//line sql.y:690
		{
			yyVAL.str = AST_LE
		}
	case 122:
		//line sql.y:694
		{
			yyVAL.str = AST_GE
		}
	case 123:
		//line sql.y:698
		{
			yyVAL.str = AST_NE
		}
	case 124:
		//line sql.y:702
		{
			yyVAL.str = AST_NSE
		}
	case 125:
		//line sql.y:708
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 126:
		//line sql.y:712
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 127:
		//line sql.y:718
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 128:
		//line sql.y:722
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 129:
		//line sql.y:728
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 130:
		//line sql.y:732
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 131:
		//line sql.y:738
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 132:
		//line sql.y:744
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 133:
		//line sql.y:748
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 134:
		//line sql.y:754
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 135:
		//line sql.y:758
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 136:
		//line sql.y:762
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 137:
		//line sql.y:766
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 138:
		//line sql.y:770
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 139:
		//line sql.y:774
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 140:
		//line sql.y:778
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 141:
		//line sql.y:782
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 142:
		//line sql.y:786
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 143:
		//line sql.y:790
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 144:
		//line sql.y:794
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 145:
		//line sql.y:798
		{
			if num, ok := yyS[yypt-0].valExpr.(NumVal); ok {
				switch yyS[yypt-1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
			}
		}
	case 146:
		//line sql.y:813
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 147:
		//line sql.y:817
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 148:
		//line sql.y:821
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 149:
		//line sql.y:825
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 150:
		//line sql.y:829
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 151:
		//line sql.y:835
		{
			yyVAL.bytes = IF_BYTES
		}
	case 152:
		//line sql.y:839
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 153:
		//line sql.y:845
		{
			yyVAL.byt = AST_UPLUS
		}
	case 154:
		//line sql.y:849
		{
			yyVAL.byt = AST_UMINUS
		}
	case 155:
		//line sql.y:853
		{
			yyVAL.byt = AST_TILDA
		}
	case 156:
		//line sql.y:859
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 157:
		//line sql.y:864
		{
			yyVAL.valExpr = nil
		}
	case 158:
		//line sql.y:868
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 159:
		//line sql.y:874
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 160:
		//line sql.y:878
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 161:
		//line sql.y:884
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 162:
		//line sql.y:889
		{
			yyVAL.valExpr = nil
		}
	case 163:
		//line sql.y:893
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 164:
		//line sql.y:899
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 165:
		//line sql.y:903
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 166:
		//line sql.y:909
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 167:
		//line sql.y:913
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 168:
		//line sql.y:917
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 169:
		//line sql.y:921
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 170:
		//line sql.y:926
		{
			yyVAL.valExprs = nil
		}
	case 171:
		//line sql.y:930
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 172:
		//line sql.y:935
		{
			yyVAL.boolExpr = nil
		}
	case 173:
		//line sql.y:939
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 174:
		//line sql.y:944
		{
			yyVAL.orderBy = nil
		}
	case 175:
		//line sql.y:948
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 176:
		//line sql.y:954
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 177:
		//line sql.y:958
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 178:
		//line sql.y:964
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 179:
		//line sql.y:969
		{
			yyVAL.str = AST_ASC
		}
	case 180:
		//line sql.y:973
		{
			yyVAL.str = AST_ASC
		}
	case 181:
		//line sql.y:977
		{
			yyVAL.str = AST_DESC
		}
	case 182:
		//line sql.y:982
		{
			yyVAL.limit = nil
		}
	case 183:
		//line sql.y:986
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 184:
		//line sql.y:990
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 185:
		//line sql.y:995
		{
			yyVAL.str = ""
		}
	case 186:
		//line sql.y:999
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 187:
		//line sql.y:1003
		{
			if !bytes.Equal(yyS[yypt-1].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyS[yypt-0].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 188:
		//line sql.y:1016
		{
			yyVAL.columns = nil
		}
	case 189:
		//line sql.y:1020
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 190:
		//line sql.y:1026
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 191:
		//line sql.y:1030
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 192:
		//line sql.y:1035
		{
			yyVAL.updateExprs = nil
		}
	case 193:
		//line sql.y:1039
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 194:
		//line sql.y:1045
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 195:
		//line sql.y:1049
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 196:
		//line sql.y:1055
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 197:
		//line sql.y:1060
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1062
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1065
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1067
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1070
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1072
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1076
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1078
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1080
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1082
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1084
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		//line sql.y:1087
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		//line sql.y:1089
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		//line sql.y:1092
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		//line sql.y:1094
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		//line sql.y:1097
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		//line sql.y:1099
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		//line sql.y:1103
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 215:
		//line sql.y:1108
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
