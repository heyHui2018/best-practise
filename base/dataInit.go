package base

// import (
// 	"bufio"
// 	"github.com/heyHui2018/best-practise/models"
// 	"io"
// 	"os"
// 	"strings"
// )
//
// func DataInit() {
// 	models.QuestionMap = make(map[string]*models.Question)
// 	question := new(models.Question)
//
// 	file, err := os.Open("/conf/leetCode.txt")
// 	if err != nil {
// 		panic(err)
// 		return
// 	}
// 	defer file.Close()
//
// 	br := bufio.NewReader(file)
// 	for {
// 		data, _, err := br.ReadLine()
// 		if err == io.EOF {
// 			break
// 		}
// 		list := strings.Split(string(data), " ")
// 		if len(list) == 4 {
// 			question.Id = list[0]
// 			question.Target = list[1]
// 			question.KeyWord = list[2]
// 			question.Solution = list[3]
// 			// 入库
// 		}
// 	}
// }
