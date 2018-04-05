package mystring

import (
	"testing"
	"fmt"
)

const TEST_STR  = `城市抢人大战|甘肃 腐败 公路|今日头条被罚|安邦被注资608亿|清明节|安徽首虎 减刑|“断崖式”降温|48岁王菲近照|孙俪身后的男人|
Youtube 枪击案|陈香梅去世|C罗首次倒钩破门|NBA勇士对雷霆|贸易战 商品清单|美团全资收购摩拜|欧冠皇马3-0尤文|特斯拉股价大反弹|中国人赴日游最多|
朴槿惠 电视直播|Spotify上市|研究生 导师压迫|广东老板怒推裁判|失控奔驰 检测|陈小武 答辩名单|1.9亿 康熙御用碗|铁路新运行图|全球市场巨震|
恒大战平泰国队|朝鲜民众听啥歌|快手反思低龄孕妇|WiFi万能钥匙被查|杨幂诈捐门后现身|最牛高空违建拆除|1522825544城市抢人大战|甘肃 腐败 公路|
今日头条被罚|安邦被注资608亿|清明节|安徽首虎 减刑|“断崖式”降温|48岁王菲近照|孙俪身后的男人|Youtube 枪击案|陈香梅去世|C罗首次倒钩破门|
NBA勇士对雷霆|贸易战 商品清单|美团全资收购摩拜|欧冠皇马3-0尤文|特斯拉股价大反弹|中国人赴日游最多|朴槿惠 电视直播|Spotify上市|研究生 导师压迫|
广东老板怒推裁判|失控奔驰 检测|陈小武 答辩名单|1.9亿 康熙御用碗|铁路新运行图|全球市场巨震|恒大战平泰国队|朝鲜民众听啥歌|快手反思低龄孕妇|
WiFi万能钥匙被查|杨幂诈捐门后现身|最牛高空违建拆除|1522825544城市抢人大战|甘肃 腐败 公路|今日头条被罚|安邦被注资608亿|清明节|安徽首虎 减刑|
“断崖式”降温|48岁王菲近照|孙俪身后的男人|Youtube 枪击案|陈香梅去世|C罗首次倒钩破门|NBA勇士对雷霆|贸易战 商品清单|美团全资收购摩拜|欧冠皇马3-0尤文|
特斯拉股价大反弹|中国人赴日游最多|朴槿惠 电视直播|Spotify上市|研究生 导师压迫|广东老板怒推裁判|失控奔驰 检测|陈小武 答辩名单|1.9亿 康熙御用碗|
铁路新运行图|全球市场巨震|恒大战平泰国队|朝鲜民众听啥歌|快手反思低龄孕妇|WiFi万能钥匙被查|杨幂诈捐门后现身|最牛高空违建拆除|1522825544`

func TestDBCtoSBC(t *testing.T) {
	DBC := "０１２３４５ａｂｃｄｅｆ"
	SBC := "012345abcdef"

	TOSBC := DBCtoSBC(DBC)
	if SBC != TOSBC {
		t.Error("dbc to sbc error dbc ", DBC, "sbc" , TOSBC)
	}
	fmt.Println(DBC)
	fmt.Println(TOSBC)
}

func TestDBCtoSBCNew(t *testing.T) {
	DBC := "０１２３４５ａｂｃｄｅｆ"
	SBC := "012345abcdef"

	TOSBC := DBCtoSBCNew(DBC)
	if SBC != TOSBC {
		t.Error("dbc to sbc error dbc ", DBC, "sbc" , TOSBC)
	}
	fmt.Println(DBC)
	fmt.Println(TOSBC)
}

func BenchmarkStrDBCtoSBC(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DBCtoSBC(TEST_STR)
	}
}

func BenchmarkStrDBCtoSBCNew(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DBCtoSBCNew(TEST_STR)
	}
}

// bench test result
// BenchmarkStrDBCtoSBC-4      	    3000	    410963 ns/op	 1007842 B/op	    1763 allocs/op
// BenchmarkStrDBCtoSBCNew-4   	  100000	     22075 ns/op	    6880 B/op	       7 allocs/op

func TestTrim(t *testing.T) {
	str := " te  st文本，换行符\n，制表符\t, lolol  "
	str_dst := "test文本，换行符，制表符,lolol"

	str_trim := Trim(str)
	if str_dst != str_trim {
		t.Error("trim error dbc ", str, " trim: " , str_trim)
	}
	fmt.Println(str)
	fmt.Println(str_trim)
}

func TestTrimNew(t *testing.T) {
	str := " te  st文本，换行符\n，制表符\t, lolol  "
	str_dst := "test文本，换行符，制表符,lolol"

	str_trim := TrimNew(str)
	if str_dst != str_trim {
		t.Error("trim error dbc ", str, " trim: " , str_trim)
	}
	fmt.Println(str)
	fmt.Println(str_trim)
}

func BenchmarkStrTrim(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Trim(TEST_STR)
	}
}

func BenchmarkStrTrimNew(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TrimNew(TEST_STR)
	}
}

func BenchmarkStrTrimNew2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TrimNew2(TEST_STR)
	}
}

// bench test result
// BenchmarkStrTrim-4       	  500000	      2886 ns/op	    4608 B/op	       2 allocs/op
// BenchmarkStrTrimNew-4    	   50000	     26124 ns/op	    6880 B/op	       7 allocs/op
// BenchmarkStrTrimNew2-4   	  100000	     17599 ns/op	    5120 B/op	       3 allocs/op

