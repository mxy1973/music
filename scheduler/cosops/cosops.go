package cosops

import (
	"context"
	"github.com/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
	"avenssi/config"
)

var SID string
var SKEY string
var EP string

var AcceURL string

func init()  {
	SID = "AKIDXszqRmxykGZcQykSQ9ez5vUn34pR20uY"
	SKEY = "X0GI5cuxkHm49zaw2uX7TaR2OULwUpyr"
	EP = config.GetCosAddr() //"https://lb-videos-1258876329.cos.ap-chengdu.myqcloud.com"

}

func UploadToCos(cosname string) bool {
	u, _ := url.Parse(EP) // https:...tencentyun.com
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: SID,
			SecretKey:SKEY,
		},
	})

	_, err := c.Object.PutFromFile(context.Background(), cosname, "../videos", nil)
	if err != nil {
		log.Printf("Object PutFromFile error: %s", err)
		return false
	}

	return true
}

func DeleteObject(cosname string) bool { //cosname需为 EP/ 下
	u, _ := url.Parse(EP) // https:...tencentyun.com
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: SID,
			SecretKey:SKEY,
		},
	})

	_, err := c.Object.Delete(context.Background(), cosname)
	if err != nil {
		log.Printf("DeleteObject error: %s", err)
		return false
	}

	return true
}
