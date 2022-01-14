package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	start := 10601
	curPath := getCurrPath()
	fmt.Println("curPath:", curPath)
	parentPath := getParentPath(curPath)
	fmt.Println("parentPath",parentPath)
	//1 遍历文件
	dir,err := ioutil.ReadDir(curPath)
	if err != nil{
		panic(err)
	}

	endTokenId := start
	for _, fi := range dir {
		if fi.IsDir() || !strings.HasSuffix(fi.Name(),".gif"){
			continue
		}
		name := path.Base(fi.Name())
		suffix := path.Ext(fi.Name())
		token := name[0:len(name)-len(suffix)]
		fmt.Println("token",token)
		tokenId,_ := strconv.Atoi(token)
		if tokenId > endTokenId{
			endTokenId = tokenId
		}
	}
	fmt.Println("endTokenId",endTokenId)
	endDirName  := start + 1000
	if endDirName>endTokenId{
		endDirName = endTokenId
	}
	outputDir := checkParentDir(parentPath,start,endDirName)
	notExistMsg := ""
	//2 每1000，新建上一级，并移动到上一级
	for i := 0;i<=endTokenId-start;i++ {
		//1000次，检测路径
		if (i+1) % 1000 == 0{
			end := start + i + 999
			if end > endTokenId{
				end = endTokenId
			}
			outputDir = checkParentDir(parentPath,start,end)
		}
		curTokenId := start+i
		//gif路径
		gifPath := fmt.Sprintf("%s/%d.gif",curPath,curTokenId)
		fmt.Println("gifPath",gifPath)
		//3 打印缺失的图片
		outputPath := fmt.Sprintf("%s/%d.gif",outputDir,curTokenId)
		if _,err := os.Stat(gifPath);os.IsNotExist(err){
			notExistMsg = fmt.Sprintf("%s\n%s",notExistMsg,gifPath)
			continue
		}
		//4 git move to out
		os.Rename(gifPath,outputPath)
	}

	//5写入err.log
	errLogPath := fmt.Sprintf("%s/err.log",parentPath)
	checkFileCreate(errLogPath)
	errLogFile,_ := os.OpenFile(errLogPath,2,0666)
	if _,err := io.WriteString(errLogFile,notExistMsg);err != nil{
		panic(err)
	}
}

func checkParentDir(parentPath string,start int,end int) string {
	fullDir := fmt.Sprintf("%s/%d-%d",parentPath,start,end)
	if _,err := os.Stat(fullDir);os.IsNotExist(err){
		os.Mkdir(fullDir,0777)
		os.Chmod(fullDir,0777)
		fmt.Println("新建文件夹",fullDir)
	}
	return fullDir
}

func checkFileCreate(filePath string)  {
	if _,err := os.Stat(filePath);os.IsNotExist(err){
		os.Create(filePath)
	}
}

func getCurrPath() string {
	str, _ := os.Getwd()
	return str
}

func getParentPath(cur string) string {
	return filepath.Dir(cur)
}
