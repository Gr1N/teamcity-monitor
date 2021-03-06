package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"

	"github.com/parnurzeal/gorequest"
)

type HttpResponse struct {
	Resp *http.Response
	Body string
}

type BuildStatus struct {
	BuildTypeId string `json:"buildTypeId"`
	Status      string `json:"status"`
	StatusText  string `json:"statusText"`
	BuildType   struct {
		Name string `json:"name"`
	} `json:"buildType"`
	LastChanges struct {
		Change []struct {
			Username string `json:"username"`
		} `json:"change"`
	} `json:"lastChanges"`
}

const (
	tcLocationBuildTypes = "%s/httpAuth/app/rest/buildTypes/id:%s/builds/count:1"
)

var (
	tcUrl              string
	tcBasicAuthEncoded string
	tcRefreshTimeout   int
	tcBuilds           []string
	tcBuildsLaout      [][]map[string]interface{}
)

func init() {
	configPath := filepath.Join(beego.AppPath, "conf", "teamcity.json")
	tcConfig, err := config.NewConfig("json", configPath)
	if err != nil {
		panic(err)
	}

	tcUrl = tcConfig.String("url")

	tcBasicAuth := fmt.Sprintf(
		"%s:%s",
		tcConfig.String("login"),
		tcConfig.String("password"),
	)
	tcBasicAuthEncoded = base64.StdEncoding.EncodeToString([]byte(tcBasicAuth))

	tcRefreshTimeout, _ = tcConfig.Int("refreshTimeout")

	tcRawBuilds, _ := tcConfig.DIY("builds")
	tcRawBuildsCasted := tcRawBuilds.([]interface{})
	for _, build := range tcRawBuildsCasted {
		tcBuilds = append(tcBuilds, build.(string))
	}

	tcRawBuildsLaout, _ := tcConfig.DIY("builds_layout")
	tcRawBuildsLaoutCasted := tcRawBuildsLaout.([]interface{})
	for _, buildLayout := range tcRawBuildsLaoutCasted {
		buildLayoutCasted := buildLayout.([]interface{})
		layout := make([]map[string]interface{}, len(buildLayoutCasted))

		for i := range buildLayoutCasted {
			layout[i] = buildLayoutCasted[i].(map[string]interface{})
		}

		tcBuildsLaout = append(tcBuildsLaout, layout)
	}
}

func asyncHttpGets(urls []string) <-chan *HttpResponse {
	ch := make(chan *HttpResponse, len(urls))

	for _, url := range urls {
		go func(url string) {
			beego.Debug(fmt.Sprintf("Fetching url: %s", url))
			resp, body, errs := gorequest.
				New().
				Get(url).
				Set("Accept", "application/json").
				Set("Authorization", "Basic "+tcBasicAuthEncoded).
				End()

			if errs != nil {
				beego.Error(fmt.Sprintf(
					"Got errors (%s) while fetching url (%s)",
					errs, url,
				))
			}

			ch <- &HttpResponse{resp, body}
		}(url)
	}

	return ch
}

func Builds() map[string]interface{} {
	builds := map[string]interface{}{
		"refreshTimeout": tcRefreshTimeout,
		"builds":         tcBuilds,
		"buildsLayout":   tcBuildsLaout,
	}

	return builds
}

func BuildsStatus() []map[string]interface{} {
	urls := make([]string, len(tcBuilds))
	for i, build := range tcBuilds {
		urls[i] = fmt.Sprintf(tcLocationBuildTypes, tcUrl, build)
	}

	results := asyncHttpGets(urls)
	buildsStatus := []map[string]interface{}{}
	for _ = range urls {
		result := <-results

		resp := result.Resp
		if resp == nil || resp.StatusCode != http.StatusOK {
			continue
		}

		buildStatus := &BuildStatus{}
		json.Unmarshal([]byte(result.Body), &buildStatus)

		buildStatusMap := map[string]interface{}{
			"id":         buildStatus.BuildTypeId,
			"name":       buildStatus.BuildType.Name,
			"status":     buildStatus.Status,
			"statusText": buildStatus.StatusText,
		}
		buildsStatusChange := buildStatus.LastChanges.Change
		if len(buildsStatusChange) > 0 {
			buildStatusMap["lastCommiter"] = buildsStatusChange[0].Username
		} else {
			buildStatusMap["lastCommiter"] = nil
		}

		buildsStatus = append(buildsStatus, buildStatusMap)
	}

	return buildsStatus
}
