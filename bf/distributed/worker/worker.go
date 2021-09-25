package worker

import (
	"errors"
	"github.com/Winszheng/crawler/distributed/config"
	"github.com/Winszheng/crawler/single/engine"
	"github.com/Winszheng/crawler/single/zhenai/parser"
	"log"
)

type SerializedParser struct {
	Name string // 函数名
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	// 自己动手，丰衣足食
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Iterms,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Iterms: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

// deserializeParser反序列化出真正的parser
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParserCityList:
		return engine.NewFuncParser(parser.ParserCityList, config.ParserCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		return engine.NewFuncParser(parser.ParseProfile, config.ParseProfile), nil
	default:
		return nil, errors.New("unknown parser name")
	}
}
