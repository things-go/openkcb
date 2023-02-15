package openkcb

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/things-go/openkcb/util"
)

type Response[T any] struct {
	Flag string `json:"flag,omitempty"`
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data *T     `json:"data,omitempty"`
}

type ResponseRaw struct {
	Flag string `json:"flag,omitempty"`
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data string `json:"data,omitempty"`
}

type OpenApiRequest struct {
	AppId         string `json:"appId,omitempty"`
	AppVersion    string `json:"appVersion,omitempty"`
	TransCode     string `json:"transCode,omitempty"`
	SequenceNo    string `json:"sequenceNo,omitempty"`
	Timestamp     string `json:"timestamp,omitempty"`
	Signature     string `json:"signature,omitempty"`
	EncryptedKey  string `json:"encryptedKey,omitempty"`
	EncryptedData string `json:"encryptedData,omitempty"`
}
type OpenApiResponse struct {
	Flag string       `json:"flag,omitempty"`
	Code string       `json:"code,omitempty"`
	Msg  string       `json:"msg,omitempty"`
	Data *BizResponse `json:"data,omitempty"`
}

type BizRequest struct {
	Head struct {
		PlatformCode string `json:"platformCode,omitempty"`
		BizNo        string `json:"bizNo,omitempty"`
	} `json:"head,omitempty"`
	Body any `json:"body,omitempty"`
}

type BizResponse struct {
	SequenceNo    string `json:"sequenceNo,omitempty"`
	Timestamp     string `json:"timestamp,omitempty"`
	Signature     string `json:"signature,omitempty"`
	EncryptedKey  string `json:"encryptedKey,omitempty"`
	EncryptedData string `json:"encryptedData,omitempty"`
}

type Validator interface {
	Validate() error
}

type Config struct {
	AppId         string `yaml:"appId" json:"appId"`
	AppVersion    string `yaml:"appVersion" json:"appVersion"`
	PlatformCode  string `yaml:"platformCode" json:"platformCode"`
	KcbUrl        string `yaml:"kcbUrl" json:"kcbUrl"`
	KcbPublicKey  string `yaml:"kcbPublicKey" json:"kcbPublicKey"`
	PrivateKey    string `yaml:"privateKey" json:"privateKey"`
	KcbPrivateKey string `yaml:"kcbPrivateKey" json:"kcbPrivateKey"` // 临时测试用，待删除
	PublicKey     string `yaml:"publicKey" json:"publicKey"`         // 临时测试用，待删除
}

type Client struct {
	config Config
	kcbPub *rsa.PublicKey
	priv   *rsa.PrivateKey

	kcbPriv *rsa.PrivateKey // 临时测试用，待删除
	pub     *rsa.PublicKey  // 临时测试用，待删除

	httpClient *http.Client
}

func NewClient(config Config) *Client {
	client := &Client{
		config:     config,
		httpClient: &http.Client{Timeout: time.Second * 60},
	}

	var err error

	client.kcbPub, err = util.ParseRSAPublicKey(client.config.KcbPublicKey)
	if err != nil {
		log.Fatal(err)
	}
	client.priv, err = util.ParseRSAPrivateKey(client.config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 临时测试用，待删除
	client.kcbPriv, err = util.ParseRSAPrivateKey(client.config.KcbPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 临时测试用，待删除
	client.pub, err = util.ParseRSAPublicKey(client.config.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (c *Client) Invoke(ctx context.Context, transCode string, req any, rsp *ResponseRaw, bizNo string) (err error) {
	// 固定长度16位
	randomKey := util.RandString(16)
	log.Println("随机key: ", randomKey)
	b, err := util.RsaEncrypt(c.kcbPub, []byte(randomKey))
	if err != nil {
		return
	}
	encryptedKey := base64.StdEncoding.EncodeToString(b)

	transData := BizRequest{}
	transData.Head.PlatformCode = c.config.PlatformCode

	if bizNo == "" {
		transData.Head.BizNo = util.RandString(16)
	} else {
		transData.Head.BizNo = bizNo
	}

	transData.Body = req

	b, err = json.Marshal(transData)
	if err != nil {
		return
	}
	log.Println("交易数据: ", string(b))
	b, _ = util.AesEncrypt(b, []byte(randomKey))
	encryptedData := base64.StdEncoding.EncodeToString(b)
	log.Println("加密后数据: ", encryptedData)

	oReq := OpenApiRequest{
		AppId:         c.config.AppId,
		AppVersion:    c.config.AppVersion,
		TransCode:     transCode,
		SequenceNo:    util.GenerateSequenceNo(),
		Timestamp:     strconv.Itoa(int(time.Now().UnixMilli())),
		Signature:     "",
		EncryptedKey:  encryptedKey,
		EncryptedData: encryptedData,
	}

	m := map[string]any{}
	m["appId"] = oReq.AppId
	m["appVersion"] = oReq.AppVersion
	m["transCode"] = oReq.TransCode
	m["sequenceNo"] = oReq.SequenceNo
	m["timestamp"] = oReq.Timestamp
	m["encryptedKey"] = oReq.EncryptedKey
	m["encryptedData"] = oReq.EncryptedData

	requestStr := util.ConcatMap(m, false)
	log.Println("待签名字符串: ", requestStr)
	sign, _ := util.SignSHA256WithRSA(c.priv, requestStr)
	log.Println("签名: ", sign)
	oReq.Signature = sign
	b, err = json.Marshal(oReq)
	if err != nil {
		return
	}
	log.Println("请求数据: ", string(b))

	httpRequest, err := http.NewRequest("POST", c.config.KcbUrl, bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return
	}
	httpRequest.Header.Add("User-Agent", "jjt-sdk-go")
	httpRequest.Header.Add("Content-Type", "application/json")

	httpResponse, err := c.httpClient.Do(httpRequest)

	if err != nil {
		if httpResponse != nil {
			httpResponse.Body.Close()
		}
		return
	}
	defer httpResponse.Body.Close()

	result, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(result))

	oRsp := OpenApiResponse{}

	if err = json.Unmarshal(result, &oRsp); err != nil {
		return
	}

	rsp.Flag = oRsp.Flag
	rsp.Code = oRsp.Code
	rsp.Msg = oRsp.Msg
	if oRsp.Data != nil {
		m := map[string]any{}
		m["sequenceNo"] = oRsp.Data.SequenceNo
		m["timestamp"] = oRsp.Data.Timestamp
		m["encryptedKey"] = oRsp.Data.EncryptedKey
		m["encryptedData"] = oRsp.Data.EncryptedData

		responseStr := util.ConcatMap(m, false)
		log.Println("应答数据:", responseStr)

		if err := util.VerifySHA256WithRSA(c.kcbPub, responseStr, oRsp.Data.Signature); err != nil {
			return errors.New("验签失败")
		}

		log.Println("验签成功")

		b, err := base64.StdEncoding.DecodeString(oRsp.Data.EncryptedKey)
		if err != nil {
			return err
		}
		b, err = util.RsaDecrypt(c.priv, b)
		if err != nil {
			return err
		}
		peerRandomKey := string(b)
		log.Println("对方随机key: ", peerRandomKey)
		b, err = base64.StdEncoding.DecodeString(oRsp.Data.EncryptedData)
		if err != nil {
			return err
		}

		b, err = util.AesDecrypt(b, []byte(peerRandomKey))
		if err != nil {
			return err
		}
		log.Println("解密后数据为:", string(b))

		//data := &ResponseRaw{}
		rsp.Flag = oRsp.Flag
		rsp.Code = oRsp.Code
		rsp.Msg = oRsp.Msg
		rsp.Data = string(b)
	}

	return nil
}
