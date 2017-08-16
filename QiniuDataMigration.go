package main

/*
*	七牛云跨账户数据迁移
 */
import (
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

func main() {

	// -------------------正式服务器密钥
	accessKey := "your new accessKey"
	scretKey := "your new scretKey"
	bucket := "bucket name"

	// 创建正式环境凭证
	mac := qbox.NewMac(accessKey, scretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}

	// 创建资源管理
	bucketManager := storage.NewBucketManager(mac, &cfg)

	// ------------------测试服务器密钥
	accessKeyTest := "test accessKey"
	scretKeyTest := "test scretKey"
	urlspaceTest := "test space url"
	bucketTest := "test space bucket"

	// 创建测试环境凭证
	macTest := qbox.NewMac(accessKeyTest, scretKeyTest)

	// 创建测试资源管理
	bucketManagerTest := storage.NewBucketManager(macTest, &cfg)

	// ----------------------------------------------遍历列表并移动
	// 每次最大列举数量为1000，规定最大为1000，超过1000用nextMarker循环列举
	limit := 1000
	// 指定文件前缀，为空代表获取所有数据
	prefix := ""
	// 指定文件目录
	delimiter := ""

	// 初始化列举marker
	marker := ""
	for {
		// 遍历测试环境文件列表
		entries, _, nextMarker, hashNext, err := bucketManagerTest.ListFiles(bucketTest, prefix, delimiter, marker, limit)
		if err != nil {
			fmt.Println("list err", err)
			break
		}

		for _, entry := range entries {
			// 这里需要将文件抓取到新服务器上面去
			url := urlspaceTest + entry.Key
			fmt.Println(url)
			// 上传到七牛云新服务器
			fetchRet, err := bucketManager.Fetch(url, bucket, entry.Key)
			if err != nil {
				fmt.Println("fatch error", err)
				return
			} else {
				fmt.Println(fetchRet.String())
			}
		}

		// 文件超过1000的部分，需要继续遍历
		if hashNext {
			marker = nextMarker
		} else {
			break
		}
	}
}
