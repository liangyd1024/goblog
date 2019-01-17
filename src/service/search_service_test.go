//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2019/1/2

package service

import (
	"fmt"
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
	"runtime"
	"testing"
)

var engine = riot.Engine{}

func init() {
	engine.Init(types.EngineOpts{
		Using:       3,
		GseDict:     "zh",
		UseStore:    true,
		StoreFolder: "./indexer",
		StoreShards: 8,
		StoreEngine: "",
	})
	engine.Flush()
}

func TestIndex(t *testing.T) {
	engine.Index("1", types.DocData{
		Content: "与框架无关的库，用于解析和验证支持多部分表单和文件的表单/ JSON数",
		Labels:  []string{"JSON", "VUE"},
	}, true)
	engine.Index("2", types.DocData{
		Content: "引擎从 StoreFolder 指定的目录中读取 文档索引数据..........",
		Labels:  []string{"HTML", "JS"},
	}, true)
	engine.Index("3", types.DocData{
		Content: "Golang实施的Raft共识协议，由HashiCorp提供。",
	}, true)
	engine.Flush()
}

func TestGetAllDocId(t *testing.T) {
	fmt.Printf("call TestGetAllDocId GetAllDocIds:%v", engine.GetAllDocIds())
}

func TestSearch(t *testing.T) {
	searchResp := engine.Search(types.SearchReq{
		Text: "grpc",
	})
	fmt.Printf("call TestSearch searchResp:%v", searchResp)
}

func TestIndexRefresh(t *testing.T) {
	engine.IndexDoc("2", types.DocData{
		Content: "grpc电子商务微服务",
	}, true)
	engine.Flush()

	fmt.Println("call TestIndexRefresh GetAllDocIds:", engine.GetAllDocIds())

	searchResp := engine.Search(types.SearchReq{
		Text: "grpc",
	})
	fmt.Println("call TestIndexRefresh searchResp:", searchResp)
	searchResp1 := engine.Search(types.SearchReq{
		Text: "StoreFolder",
	})
	fmt.Println("call TestIndexRefresh searchResp1:", searchResp1)
}

func TestSearch2(t *testing.T) {
	runtime.Gosched()
	searchResp := engine.Search(types.SearchReq{
		//Text: "VUE",
		Labels: []string{"JS"},
	})
	fmt.Printf("call TestSearch2 searchResp:%v", searchResp)
}

func TestRemove(t *testing.T) {
	engine.RemoveDoc("1",true)
	engine.Flush()
	searchResp := engine.Search(types.SearchReq{
		Text: "协议",
	})
	fmt.Printf("call TestRemove searchResp:%v", searchResp)
	fmt.Printf("call TestGetAllDocId GetAllDocIds:%v", engine.GetAllDocIds())
}
